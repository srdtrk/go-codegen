package codegen

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"

	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
)

const defPrefix = "#/definitions/"

var (
	globalDefRegistry = map[string]*schemas.JSONSchema{}
	generatingDefs    = false

	changesMap = map[string]*schemas.JSONSchema{}

	// nullableTypes is used to store the names of the nullable types generated
	// so that they are not generated multiple times.
	nullableTypes = []string{}
)

func generateDefinitions(f *jen.File) {
	generatingDefs = true
	for name, schema := range globalDefRegistry {
		generateDefinition(f, name, schema)
	}
	generatingDefs = false

	changesOccured := len(changesMap) != 0
	for changesOccured {
		tempMap := make(map[string]*schemas.JSONSchema)
		for k, v := range changesMap {
			tempMap[k] = v
			RegisterDefinition(k, v)
		}
		// clear changesMap
		changesMap = map[string]*schemas.JSONSchema{}

		generatingDefs = true
		for k, v := range tempMap {
			generateDefinition(f, k, v)
		}
		generatingDefs = false

		changesOccured = len(changesMap) != 0
	}
}

// RegisterDefinitions registers definitions to the global definition registry.
func RegisterDefinitions(definitions map[string]*schemas.JSONSchema) bool {
	allSuccessful := true
	for name, schema := range definitions {
		res := RegisterDefinition(name, schema)
		if !res {
			allSuccessful = false
		}
	}

	return allSuccessful
}

// RegisterDefinition registers a definition to the global definition registry
// and returns true if the definition is successfully registered.
// If the definition is already registered, it returns false.
func RegisterDefinition(ref string, schema *schemas.JSONSchema) bool {
	// check if the ref is already registered
	if regSchema, ok := globalDefRegistry[ref]; ok {
		if err := areDefinitionsEqual(regSchema, schema); err != nil {
			panic(fmt.Sprintf("duplicate definition `%s` with differing implementations: %s", ref, err.Error()))
		}

		return false
	}

	if generatingDefs {
		changesMap[ref] = schema
		return true
	}

	globalDefRegistry[ref] = schema
	return true
}

// GetDefinition returns a definition from the global definition registry.
// If the definition is not found, it returns false.
func GetDefinition(ref string) (*schemas.JSONSchema, bool) {
	schema, ok := globalDefRegistry[ref]
	if !ok {
		return nil, false
	}

	return schema, true
}

// TODO: this function may require a refactor, possibly allowing for recursive definitions.
func generateDefinition(f *jen.File, name string, schema *schemas.JSONSchema) {
	if err := validateAsDefinition(name, schema); err != nil {
		panic(err)
	}

	switch {
	case len(schema.Type) == 1:
		if err := generateDefinitionType(f, name, schema); err != nil {
			panic(err)
		}
	case len(schema.OneOf) != 0:
		if err := generateDefinitionOneOf(f, name, schema); err != nil {
			panic(err)
		}
	case len(schema.AllOf) == 1:
		if err := generateDefinitionAllOf(f, name, schema); err != nil {
			panic(err)
		}
	case schema.Ref != nil:
		if err := generateDefinitionRef(f, name, schema); err != nil {
			panic(err)
		}
	case len(schema.AnyOf) != 0:
		if _, err := generateDefinitionAnyOf(f, name, schema); err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("unsupported definition %s", name))
	}
}

func areDefinitionsEqual(a, b *schemas.JSONSchema) error {
	if a == nil || b == nil {
		return fmt.Errorf("nil schema")
	}

	if !slices.Equal(a.Type, b.Type) {
		return fmt.Errorf("different types")
	}

	if len(a.Type) == 1 {
		switch a.Type[0] {
		case schemas.TypeNameString:
			return nil
		case schemas.TypeNameInteger:
			return nil
		case schemas.TypeNameNumber:
			return nil
		case schemas.TypeNameBoolean:
			return nil
		}
	}

	if a.Description != b.Description {
		return fmt.Errorf("different descriptions")
	}

	return nil
}

func generateDefinitionType(f *jen.File, name string, schema *schemas.JSONSchema) error {
	if err := validateAsDefinition(name, schema); err != nil {
		return err
	}

	f.Comment(schema.Description)

	switch schema.Type[0] {
	case schemas.TypeNameString:
		f.Type().Id(name).String()
		if schema.Enum != nil {
			err := generateEnumString(f, name, schema.Enum)
			if err != nil {
				return err
			}
		}
	case schemas.TypeNameInteger:
		f.Type().Id(name).Int()
	case schemas.TypeNameNumber:
		f.Type().Id(name).Float64()
	case schemas.TypeNameBoolean:
		f.Type().Id(name).Bool()
	case schemas.TypeNameObject:
		f.Type().Id(name).Struct(
			generateFieldsFromProperties(schema.Properties, true)...,
		)
	case schemas.TypeNameArray:
		switch {
		case len(schema.Items) > 1 && !slices.ContainsFunc(schema.Items, func(s schemas.JSONSchema) bool { return len(s.Type) != 1 || s.Type[0] != schema.Items[0].Type[0] }):
			// This case means that all items have the same type. This is similar to having an array of a single item.
			fallthrough
		case len(schema.Items) == 1 && schema.MaxItems == nil && schema.MinItems == nil:
			item := &schema.Items[0]
			itemName, err := getType(name, item, nil, "")
			if err != nil {
				return err
			}

			f.Type().Id(name).Index().Id(itemName)

			RegisterDefinition(itemName, item)
		case schema.MaxItems != nil && schema.MinItems != nil && *schema.MaxItems == *schema.MinItems && len(schema.Items) == *schema.MaxItems:
			err := generateDefinitionTuple(f, name, schema)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported array definition %s", name)
		}
	default:
		return fmt.Errorf("unsupported type %s for definition %s", schema.Type[0], name)
	}

	return nil
}

func generateDefinitionOneOf(f *jen.File, name string, schema *schemas.JSONSchema) error {
	// Check if all oneOf types are the same
	if len(schema.OneOf[0].Type) != 1 {
		return fmt.Errorf("type of the enum variant %s is not supported", name)
	}

	areSame := !slices.ContainsFunc(schema.OneOf, func(s *schemas.JSONSchema) bool {
		return len(s.Type) != 1 || s.Type[0] != schema.OneOf[0].Type[0]
	})

	if areSame {
		return generateDefinitionOneOfAllSame(f, name, schema)
	}

	funcName := "Implements_" + name
	f.Comment(schema.Description)
	f.Type().Id(name).Interface(jen.Id(funcName).Params())

	for _, oneOf := range schema.OneOf {
		if len(oneOf.Type) != 1 {
			return fmt.Errorf("type of the enum variant %s is not supported", name)
		}
		switch oneOf.Type[0] {
		case schemas.TypeNameObject:
			if len(oneOf.Properties) != 1 {
				return fmt.Errorf("properties of the enum variant %s is not supported", name)
			}

			for k, prop := range oneOf.Properties {
				camelKey := strcase.ToCamel(k)
				propName := name + "_" + camelKey

				// Assert Enum implementation
				f.Var().Id("_").Op(name).Op("=").Params(
					jen.Op("*").Op(propName),
				).Params(jen.Nil())

				if prop.Ref != nil {
					if !strings.HasPrefix(*prop.Ref, defPrefix) {
						panic(fmt.Errorf("unknown reference %s", *prop.Ref))
					}
					refType := strings.TrimPrefix(*prop.Ref, defPrefix)

					f.Comment(prop.Description)
					f.Type().Id(propName).Struct(
						jen.Id(camelKey).Op(refType).Tag(map[string]string{"json": k}),
					)
				} else {
					err := generateDefinitionType(f, propName, prop)
					if err != nil {
						return err
					}
				}

				f.Func().Params(
					jen.Op("*").Id(propName),
				).Id(funcName).Params().Block()
			}

		case schemas.TypeNameString:
			if len(oneOf.Enum) != 1 {
				panic(fmt.Errorf("unsupported enum variant %v", oneOf))
			}

			variantName := name + "_" + strcase.ToCamel(oneOf.Enum[0])

			f.Var().Id("_").Op(name).Op("=").Params(
				jen.Op("*").Op(variantName),
			).Params(jen.Nil())

			f.Type().Id(variantName).String()

			f.Comment(oneOf.Description)
			constName := variantName + "_Value"
			f.Const().Id(constName).Op(variantName).Op("=").Lit(oneOf.Enum[0])

			f.Func().Params(
				jen.Op("*").Id(variantName),
			).Id(funcName).Params().Block()
		}
	}

	// TODO: implement more cases and combine some cases such as all strings and all objects.

	return nil
}

func generateDefinitionOneOfAllSame(f *jen.File, name string, schema *schemas.JSONSchema) error {
	switch schema.OneOf[0].Type[0] {
	case schemas.TypeNameObject:
		f.Comment(schema.Description)
		f.Type().Id(name).Struct(
			generateFieldsFromOneOf(schema.OneOf, name+"_")...,
		)
	case schemas.TypeNameString:
		f.Comment(schema.Description)
		f.Type().Id(name).String()

		constants := []jen.Code{}
		for _, oneOf := range schema.OneOf {
			if len(oneOf.Enum) != 1 {
				panic(fmt.Errorf("unsupported string enum %s", name))
			}

			constName := name + "_" + strcase.ToCamel(oneOf.Enum[0])
			constants = append(constants, jen.Comment(oneOf.Description))
			constants = append(constants, jen.Id(constName).Op(name).Op("=").Lit(oneOf.Enum[0]))
		}

		f.Const().Defs(constants...)
	default:
		types.DefaultLogger().Error().Msgf("unsupported oneOf type %s for definition %s, please create an issue in 'https://github.com/srdtrk/go-codegen'", schema.OneOf[0].Type[0], name)
	}
	return nil
}

// generateEnumString generates a string enum type.
func generateEnumString(f *jen.File, name string, enum []string) error {
	constants := make([]jen.Code, len(enum))
	for i, e := range enum {
		eName := name + "_" + strcase.ToCamel(e)
		constants[i] = jen.Id(eName).Op(name).Op("=").Lit(e)
	}

	f.Const().Defs(constants...)

	return nil
}

func generateDefinitionAllOf(f *jen.File, name string, schema *schemas.JSONSchema) error {
	if len(schema.AllOf) != 1 || schema.AllOf[0].Ref == nil {
		return fmt.Errorf("wrapper %s is not supported", name)
	}

	if !strings.HasPrefix(*schema.AllOf[0].Ref, defPrefix) {
		return fmt.Errorf("ref %s is not a definition", *schema.AllOf[0].Ref)
	}

	defTypeName := strings.TrimPrefix(*schema.AllOf[0].Ref, defPrefix)

	f.Comment(schema.Description)
	f.Type().Id(name).Op(defTypeName)

	return nil
}

func generateDefinitionTuple(f *jen.File, name string, schema *schemas.JSONSchema) error {
	if schema.MaxItems == nil || schema.MinItems == nil || *schema.MaxItems != *schema.MinItems {
		return fmt.Errorf("unsupported tuple from %s", name)
	}

	isRequired := true
	typeName, err := getType(name, schema, &isRequired, "")
	if err != nil {
		return err
	}

	pseudoProps := make(map[string]*schemas.JSONSchema)
	marshalCode := []jen.Code{}
	unmarshalCode := []jen.Code{
		jen.Var().Id("arr").Index().Qual("encoding/json", "RawMessage"),
		jen.If(jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Op("&").Id("arr")), jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Err())),
		jen.If(jen.Len(jen.Id("arr")).Op("!=").Lit(*schema.MaxItems)).Block(jen.Return(jen.Qual("errors", "New").Call(jen.Lit(fmt.Sprintf("expected %d elements in the tuple", *schema.MaxItems))))),
		jen.Line(),
	}
	var marshalReturnCode *jen.Statement
	for i := 0; i < *schema.MaxItems; i++ {
		fieldName := fmt.Sprintf("F%d", i)
		pseudoProps[fieldName] = &schema.Items[i]

		varName := fmt.Sprintf("f%d", i)
		marshalCode = append(marshalCode, jen.Id(varName).Op(",").Err().Op(":=").Qual("encoding/json", "Marshal").Call(jen.Id("t").Dot(fieldName)))
		marshalCode = append(marshalCode, jen.If(jen.Err().Op("!=").Nil()).Block(jen.Return(jen.Nil(), jen.Err())))
		marshalCode = append(marshalCode, jen.Line())

		unmarshalCode = append(unmarshalCode, jen.If(
			jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("arr").Index(jen.Lit(i)), jen.Op("&").Id("t").Dot(fieldName)),
			jen.Err().Op("!=").Nil(),
		).Block(jen.Return(jen.Err())))
		unmarshalCode = append(unmarshalCode, jen.Line())

		if i == 0 {
			marshalReturnCode = jen.Lit("[").Op("+").String().Parens(jen.Id(varName))
		} else {
			marshalReturnCode = marshalReturnCode.Op("+").Lit(",").Op("+").String().Parens(jen.Id(varName))
		}
	}

	marshalReturnCode = marshalReturnCode.Op("+").Lit("]")
	marshalCode = append(marshalCode, jen.Return(jen.Index().Byte().Parens(marshalReturnCode), jen.Nil()))
	unmarshalCode = append(unmarshalCode, jen.Return(jen.Nil()))

	f.Var().Defs(
		jen.Id("_").Qual("encoding/json", "Marshaler").Op("=").Parens(jen.Op("*").Id(typeName)).Parens(jen.Nil()),
		jen.Id("_").Qual("encoding/json", "Unmarshaler").Op("=").Parens(jen.Op("*").Id(typeName)).Parens(jen.Nil()),
	)

	f.Comment(fmt.Sprintf("%s is a tuple with custom marshal and unmarshal methods", typeName))
	f.Comment(schema.Description)
	f.Type().Id(typeName).Struct(
		generateFieldsFromProperties(pseudoProps, false)...,
	)

	f.Comment(fmt.Sprintf("MarshalJSON implements the json.Marshaler interface for %s", typeName))
	f.Func().Params(
		jen.Id("t").Id(typeName),
	).Id("MarshalJSON").Params().Params(jen.Index().Byte(), jen.Error()).Block(
		marshalCode...,
	)

	f.Comment(fmt.Sprintf("UnmarshalJSON implements the json.Unmarshaler interface for %s", typeName))
	f.Func().Params(
		jen.Id("t").Op("*").Id(typeName),
	).Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Error().Block(
		unmarshalCode...,
	)

	return nil
}

func generateDefinitionRef(f *jen.File, name string, schema *schemas.JSONSchema) error {
	if !strings.HasPrefix(*schema.Ref, defPrefix) {
		return fmt.Errorf("ref %s is not a definition", *schema.Ref)
	}

	defTypeName := strings.TrimPrefix(*schema.Ref, defPrefix)
	f.Type().Id(name).Op(defTypeName)

	return nil
}

func generateDefinitionAnyOf(f *jen.File, name string, schema *schemas.JSONSchema) (string, error) {
	if schema.Ref != nil {
		if !strings.HasPrefix(*schema.Ref, defPrefix) {
			return "", fmt.Errorf("unknown reference %s", *schema.Ref)
		}

		typeStr := strings.TrimPrefix(*schema.Ref, defPrefix)
		return typeStr, nil
	}

	if len(schema.AnyOf) != 2 {
		return "", fmt.Errorf("anyOf %s is not supported", name)
	}

	if len(schema.AnyOf[1].Type) != 1 || schema.AnyOf[1].Type[0] != schemas.TypeNameNull {
		return "", fmt.Errorf("anyOf %s is not an option type", name)
	}

	subName, err := generateDefinitionAnyOf(f, name, schema.AnyOf[0])
	if err != nil {
		return "", err
	}

	name = "Nullable_" + subName
	if !slices.Contains(nullableTypes, name) {
		f.Comment(fmt.Sprintf("%s is a nullable type of %s", name, subName))
		f.Comment(schema.Description)
		f.Type().Id(name).Op("=").Op("*").Id(subName)

		nullableTypes = append(nullableTypes, name)
	}

	return name, nil
}

// validateAsDefinition validates if the schema is a valid definition.
func validateAsDefinition(name string, schema *schemas.JSONSchema) error {
	if len(schema.Type) != 1 && len(schema.OneOf) == 0 && len(schema.AllOf) != 1 && schema.Ref == nil && len(schema.AnyOf) == 0 {
		return fmt.Errorf("definition %s is unsupported", name)
	}

	return nil
}
