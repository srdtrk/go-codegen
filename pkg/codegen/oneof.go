package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateFieldsFromOneOf(oneOf []*schemas.JSONSchema, typePrefix string) []jen.Code {
	ptrFalse := false
	fields := []jen.Code{}
	for _, schema := range oneOf {
		name := schema.Title
		if schema.Title == "" {
			if len(schema.Properties) != 1 {
				panic(fmt.Errorf("cannot determine the name of the field %v", schema))
			}

			for k := range schema.Properties {
				name = k
			}
		}
		// add comment
		fields = append(fields, jen.Comment(schema.Description))
		// add field
		fields = append(fields, GenerateFieldFromSchema(name, schema, &ptrFalse, typePrefix))
	}
	return fields
}
