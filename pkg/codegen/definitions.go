package codegen

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

const defPrefix = "#/definitions/"

var globalDefRegistry = map[string]*schemas.JSONSchema{}

func GenerateDefinitions(f *jen.File, defs map[string]*schemas.JSONSchema) {
	for name, schema := range defs {
		if !registerDefinition(name, schema) {
			continue
		}

		generateDefinition(f, name, schema)
	}
}

// registerDefinition registers a definition to the global definition registry
// and returns true if the definition is successfully registered.
// If the definition is already registered, it returns false.
func registerDefinition(ref string, schema *schemas.JSONSchema) bool {
	// check if the ref is already registered
	if regSchema, ok := globalDefRegistry[ref]; ok {
		if regSchema.Title != schema.Title || regSchema.Description != schema.Description {
			panic(fmt.Sprintf("definition %s is already registered with a different schema", ref))
		}

		return false
	}

	globalDefRegistry[ref] = schema
	return true
}

func generateDefinition(f *jen.File, name string, schema *schemas.JSONSchema) {
	if err := validateAsDefinition(name, schema); err != nil {
		panic(err)
	}

	switch {
	case len(schema.Type) == 1:
		if err := generateDefinitionType(f, name, schema); err != nil {
			panic(err)
		}
	case schema.OneOf != nil:
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
	default:
		panic(fmt.Sprintf("unsupported definition %s", name))
	}
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
			GenerateFieldsFromProperties(schema.Properties)...,
		)
	default:
		panic(fmt.Sprintf("unsupported type %s for definition %s", schema.Type[0], name))
	}

	return nil
}

//nolint:unparam
func generateDefinitionOneOf(f *jen.File, name string, schema *schemas.JSONSchema) error {
	// if all enum values are strings, then generate a single string enum type
	sEnum := []jen.Code{}
	for i, oneOf := range schema.OneOf {
		if len(oneOf.Type) != 1 || oneOf.Type[0] != schemas.TypeNameString || len(oneOf.Enum) != 1 {
			break
		}

		sEnum = append(sEnum, jen.Comment(oneOf.Description))
		eName := name + "_" + strcase.ToCamel(oneOf.Enum[0])
		sEnum = append(sEnum, jen.Id(eName).Op(name).Op("=").Lit(oneOf.Enum[0]))

		if i == len(schema.OneOf)-1 {
			f.Comment(schema.Description)
			f.Type().Id(name).String()
			f.Const().Defs(sEnum...)
			return nil
		}
	}
	// TODO: implement

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

func generateDefinitionRef(f *jen.File, name string, schema *schemas.JSONSchema) error {
	if !strings.HasPrefix(*schema.Ref, defPrefix) {
		return fmt.Errorf("ref %s is not a definition", *schema.Ref)
	}

	defTypeName := strings.TrimPrefix(*schema.Ref, defPrefix)
	f.Type().Id(name).Op(defTypeName)

	return nil
}

// validateAsDefinition validates if the schema is a valid definition.
func validateAsDefinition(name string, schema *schemas.JSONSchema) error {
	if len(schema.Type) != 1 && schema.OneOf == nil && len(schema.AllOf) != 1 && schema.Ref == nil {
		return fmt.Errorf("definition %s is unsupported", name)
	}

	return nil
}
