package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateInstantiateMsg(f *jen.File, schema *schemas.JSONSchema) {
	if err := validateAsInstantiateMsg(schema); err != nil {
		panic(err)
	}

	f.Comment(schema.Description)
	f.Type().Id(schema.Title).Struct(
		GenerateFieldsFromProperties(schema.Properties)...,
	)

	GenerateDefinitions(f, schema.Definitions)
}

func validateAsInstantiateMsg(schema *schemas.JSONSchema) error {
	if schema.Title != "InstantiateMsg" {
		return fmt.Errorf("title must be InstantiateMsg")
	}
	if schema.Type[0] != "object" {
		return fmt.Errorf("InstantiateMsg type must be object")
	}
	if schema.Properties == nil {
		return fmt.Errorf("InstantiateMsg must have properties")
	}

	return nil
}
