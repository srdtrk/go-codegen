package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func generateFieldsFromOneOf(oneOf []*schemas.JSONSchema, typePrefix string) []jen.Code {
	ptrFalse := false
	fields := []jen.Code{}
	for _, schema := range oneOf {
		if schema.Title == "" && len(schema.Properties) != 1 {
			panic(fmt.Errorf("cannot determine the name of the field %v", schema))
		}

		var (
			name    string
			jsonKey string
		)
		// NOTE: there is only one property in this map
		for k, prop := range schema.Properties {
			name, jsonKey = validatedTitleAndJSONKey(schema.Title, k)

			typeName := typePrefix + strcase.ToCamel(jsonKey)

			RegisterDefinition(typeName, prop)
		}

		RegisterDefinitions(schema.Definitions)

		// add comment
		fields = append(fields, jen.Comment(schema.Description))
		// add field
		fields = append(fields, generateFieldFromSchema(name, jsonKey, schema, &ptrFalse, typePrefix, true))
	}
	return fields
}

func validatedTitleAndJSONKey(title, key string) (string, string) {
	if title == "" && key == "" {
		panic(fmt.Errorf("cannot determine the name of the field"))
	}
	if title == "" {
		title = strcase.ToCamel(key)
	}
	if key == "" {
		key = strcase.ToSnake(title)
	}
	return title, key
}
