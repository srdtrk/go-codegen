package codegen

import (
	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateFieldsFromOneOf(oneOf []*schemas.JSONSchema) []jen.Code {
	ptrFalse := false
	fields := []jen.Code{}
	for _, schema := range oneOf {
		// add comment
		fields = append(fields, jen.Comment(schema.Description))
		// add field
		fields = append(fields, GenerateFieldFromSchema(schema.Title, schema, &ptrFalse))
	}
	return fields
}
