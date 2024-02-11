package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateExecuteMsg(f *jen.File, schema *schemas.JSONSchema) {
	if err := validateAsExecuteMsg(schema); err != nil {
		panic(err)
	}

	f.Comment(schema.Description)
	f.Type().Id(schema.Title).Struct(
		GenerateFieldsFromProperties(schema.Properties)...,
	)

	GenerateDefinitions(f, schema.Definitions)
}

func validateAsExecuteMsg(schema *schemas.JSONSchema) error {
	if schema.Title != "ExecuteMsg" && schema.Title != "ExecuteMsg_for_Empty" {
		return fmt.Errorf("title must be InstantiateMsg")
	}
	if schema.OneOf == nil {
		return fmt.Errorf("ExecuteMsg must be an enum")
	}

	return nil
}
