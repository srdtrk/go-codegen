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
		name := schema.Title
		if schema.Title == "" {
			if len(schema.Properties) != 1 {
				panic(fmt.Errorf("cannot determine the name of the field %v", schema))
			}

			for k, prop := range schema.Properties {
				name = k

				typeName := typePrefix + strcase.ToCamel(k)

				RegisterDefinition(typeName, prop)
			}
		}

		RegisterDefinitions(schema.Definitions)

		// add comment
		fields = append(fields, jen.Comment(schema.Description))
		// add field
		fields = append(fields, generateFieldFromSchema(name, schema, &ptrFalse, typePrefix, true))
	}
	return fields
}
