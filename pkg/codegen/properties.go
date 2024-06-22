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

func generateFieldsFromProperties(props map[string]*schemas.JSONSchema, useTags bool) []jen.Code {
	fields := []jen.Code{}
	for name, schema := range props {
		// add comment
		fields = append(fields, jen.Comment(schema.Description))
		// add field
		fields = append(fields, generateFieldFromSchema(name, name, schema, nil, "", useTags))
	}
	return fields
}

func generateFieldFromSchema(name, jsonKey string, schema *schemas.JSONSchema, required *bool, typePrefix string, useTags bool) jen.Code {
	if name == "" || jsonKey == "" {
		panic(fmt.Errorf("cannot determine the name of the field for schema %v", schema))
	}
	pascalName := strcase.ToCamel(name)

	typeStr, err := getType(pascalName, schema, required, typePrefix, true)
	if err != nil {
		panic(err)
	}

	if useTags {
		tags := map[string]string{}
		if strings.HasPrefix(typeStr, "*") {
			tags["json"] = jsonKey + ",omitempty"
		} else {
			tags["json"] = jsonKey
		}

		return jen.Id(pascalName).Op(typeStr).Tag(tags)
	}

	return jen.Id(pascalName).Op(typeStr)
}

func getType(name string, schema *schemas.JSONSchema, required *bool, typePrefix string, isField bool) (string, error) {
	if len(schema.Type) == 0 {
		var underlyingSchemas []*schemas.JSONSchema
		switch {
		case schema.AllOf != nil:
			if len(schema.AllOf) > 1 {
				return "", fmt.Errorf("length of allOf is greater than 1 in %s", name)
			}
			underlyingSchemas = schema.AllOf
		case schema.AnyOf != nil:
			switch len(schema.AnyOf) {
			case 0:
				return "", fmt.Errorf("length of anyOf is 0 in %s", name)
			case 1:
				underlyingSchemas = schema.AnyOf
			case 2:
				if slices.ContainsFunc(schema.AnyOf, func(s *schemas.JSONSchema) bool {
					return slices.Contains(s.Type, schemas.TypeNameNull)
				}) {
					if isField {
						notRequired := false
						return getType(name, schema.AnyOf[0], &notRequired, typePrefix, true)
					}

					required := true
					return getType(name, schema.AnyOf[0], &required, "Nullable_"+typePrefix, false)
				}
			default:
				return "", fmt.Errorf("length of anyOf is %d in %s", len(schema.AllOf), name)
			}
		case schema.Ref != nil:
			if !strings.HasPrefix(*schema.Ref, defPrefix) {
				return "", fmt.Errorf("cannot determine the type of %s: ref is not prefixed with %s", name, defPrefix)
			}

			typeStr := typePrefix + strings.TrimPrefix(*schema.Ref, defPrefix)
			if isField && required != nil && !*required {
				typeStr = "*" + typeStr
			}

			return typeStr, nil
		case schema.OneOf != nil:
			if schema.Title != "" {
				return schema.Title, nil
			}

			return "", fmt.Errorf("cannot determine the type of %s: title is empty", name)
		default:
			return "", fmt.Errorf("cannot determine the type of %s: type is not matched; %v", name, schema)
		}

		if len(underlyingSchemas) == 0 || len(underlyingSchemas) > 2 {
			return "", fmt.Errorf("cannot determine the type of %s: underlying schemas length is %d", name, len(underlyingSchemas))
		}

		if underlyingSchemas[0].Ref == nil || len(*underlyingSchemas[0].Ref) == 0 {
			return "", fmt.Errorf("cannot determine the type of %s: ref is nil or empty", name)
		}

		var isOptional bool
		if required != nil {
			isOptional = !*required
		} else {
			isOptional = slices.ContainsFunc(underlyingSchemas, func(s *schemas.JSONSchema) bool {
				return slices.Contains(s.Type, schemas.TypeNameNull)
			})
		}

		if !strings.HasPrefix(*underlyingSchemas[0].Ref, defPrefix) {
			return "", fmt.Errorf("cannot determine the type of %s: ref is not prefixed with %s", name, defPrefix)
		}

		typeStr := strings.TrimPrefix(*underlyingSchemas[0].Ref, defPrefix)
		typeStr = typePrefix + typeStr
		if isOptional {
			typeStr = "*" + typeStr
		}

		return typeStr, nil
	}

	var isOptional bool
	if required != nil {
		isOptional = !*required
	} else {
		isOptional = slices.Contains(schema.Type, schemas.TypeNameNull)
	}

	var typeStr string
	switch schema.Type[0] {
	case schemas.TypeNameString:
		typeStr = "string"
	case schemas.TypeNameInteger:
		typeStr = "int"
	case schemas.TypeNameNumber:
		typeStr = "float64"
	case schemas.TypeNameBoolean:
		typeStr = "bool"
	case schemas.TypeNameNull:
		typeStr = "any"
		types.DefaultLogger().Warn().Msgf("null type is used in %s, any is used instead", name)
	case schemas.TypeNameArray:
		if typePrefix != "" {
			return "", fmt.Errorf("cannot determine the type of array %s; type prefix is not supported", name)
		}

		switch {
		case len(schema.Items) > 1 && !slices.ContainsFunc(schema.Items, func(s schemas.JSONSchema) bool { return len(s.Type) != 1 || s.Type[0] != schema.Items[0].Type[0] }):
			// This case means that all items have the same type. This is similar to having an array of a single item.
			fallthrough
		case len(schema.Items) == 1 && schema.MaxItems == nil && schema.MinItems == nil:
			baseType, err := getType(schema.Title, &schema.Items[0], nil, "", false)
			if err != nil {
				return "", err
			}

			typeStr = "[]" + baseType

			isOptional = false // arrays are always nullable
		case schema.MaxItems != nil && schema.MinItems != nil && *schema.MaxItems == *schema.MinItems && len(schema.Items) == *schema.MaxItems:
			typeStr = "Tuple_of_"
			for i := 0; i < *schema.MaxItems; i++ {
				itemType, err := getType(name, &schema.Items[i], nil, "", false)
				if err != nil {
					return "", err
				}

				typeStr += itemType

				if i < *schema.MaxItems-1 {
					typeStr += "_and_"
				}
			}

			RegisterDefinition(typeStr, schema)
		default:
			return "", fmt.Errorf("unsupported array definition %s", name)
		}
	case schemas.TypeNameObject:
		switch {
		case schema.Title != "":
			typeStr = schema.Title
		case len(schema.Properties) == 1:
			for k := range schema.Properties {
				typeStr = strcase.ToCamel(k)
			}
		default:
			return "", fmt.Errorf("cannot determine the type of object %s", name)
		}
	default:
		return "", fmt.Errorf("cannot determine the type of %s: type is not matched", name)
	}

	typeStr = typePrefix + typeStr

	if isOptional {
		typeStr = "*" + typeStr
	}

	return typeStr, nil
}
