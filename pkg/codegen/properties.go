package codegen

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateFieldsFromProperties(props map[string]*schemas.JSONSchema) []jen.Code {
	fields := []jen.Code{}
	for name, schema := range props {
		// add comment
		fields = append(fields, jen.Comment(schema.Description))
		// add field
		fields = append(fields, GenerateFieldFromProperty(name, schema))
	}
	return fields
}

func GenerateFieldFromProperty(name string, schema *schemas.JSONSchema) jen.Code {
	pascalName := strcase.ToCamel(name)
	typeStr, err := getType(pascalName, schema)
	if err != nil {
		panic(err)
	}
	return jen.Id(pascalName).Op(typeStr).Tag(map[string]string{"json": name + ",omitempty"})
}

func getType(name string, schema *schemas.JSONSchema) (string, error) {
	if len(schema.Type) == 0 {
		var underlyingSchemas []*schemas.JSONSchema
		//nolint:gocritic //TODO(serdar): use switch
		if schema.AllOf != nil {
			if len(schema.AllOf) > 1 {
				return "", fmt.Errorf("length of allOf is greater than 1 in %s", name)
			}
			underlyingSchemas = schema.AllOf
		} else if schema.AnyOf != nil {
			if len(schema.AnyOf) > 2 {
				return "", fmt.Errorf("length of anyOf is greater than 2 in %s", name)
			}
			underlyingSchemas = schema.AnyOf
		} else if schema.Ref != nil {
			if !strings.HasPrefix(*schema.Ref, "#/definitions/") {
				return "", fmt.Errorf("cannot determine the type of %s", name)
			}

			typeStr := strings.TrimPrefix(*schema.Ref, "#/definitions/")
			return typeStr, nil
		} else {
			return "", fmt.Errorf("cannot determine the type of %s", name)
		}

		if len(underlyingSchemas) == 0 || len(underlyingSchemas) > 2 {
			return "", fmt.Errorf("cannot determine the type of %s", name)
		}

		if underlyingSchemas[0].Ref == nil {
			return "", fmt.Errorf("cannot determine the type of %s", name)
		}
		if len(*underlyingSchemas[0].Ref) == 0 {
			return "", fmt.Errorf("cannot determine the type of %s", name)
		}

		isOptional := slices.ContainsFunc(underlyingSchemas, func(s *schemas.JSONSchema) bool {
			return slices.Contains(s.Type, "null")
		})

		if !strings.HasPrefix(*underlyingSchemas[0].Ref, "#/definitions/") {
			return "", fmt.Errorf("cannot determine the type of %s", name)
		}

		typeStr := strings.TrimPrefix(*underlyingSchemas[0].Ref, "#/definitions/")
		if isOptional {
			typeStr = "*" + typeStr
		}

		return typeStr, nil
	}

	isOptional := slices.Contains(schema.Type, "null")

	var typeStr string
	switch schema.Type[0] {
	case "string":
		typeStr = "string"
	case "integer":
		typeStr = "int"
	case "number":
		typeStr = "float64"
	case "boolean":
		typeStr = "bool"
	case "array":
		baseType, err := getType(schema.Title, &schema.Items[0])
		if err != nil {
			return "", err
		}

		typeStr = "[]" + baseType
	case "object":
		typeStr = schema.Title
	default:
		return "", fmt.Errorf("cannot determine the type of %s", name)
	}

	if isOptional {
		typeStr = "*" + typeStr
	}

	return typeStr, nil
}
