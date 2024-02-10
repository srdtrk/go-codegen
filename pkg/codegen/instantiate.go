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

	comment := schema.Description
	if comment == "" {
		comment = "The message to instantiate a new contract instance."
	}

	f.Comment(comment)
	f.Type().Id(schema.Title).Struct(
		GenerateFieldsFromProperties(schema.Properties)...,
	)
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
