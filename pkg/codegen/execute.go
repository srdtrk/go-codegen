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
		GenerateFieldsFromOneOf(schema.OneOf, schema.Title+"_")...,
	)

	GenerateDefinitions(f, schema.Definitions)
}

func validateAsExecuteMsg(schema *schemas.JSONSchema) error {
	if schema.Title != "ExecuteMsg" && schema.Title != "ExecuteMsg_for_Empty" {
		return fmt.Errorf("title must be InstantiateMsg")
	}
	if len(schema.OneOf) == 0 {
		return fmt.Errorf("ExecuteMsg is empty")
	}

	return nil
}
