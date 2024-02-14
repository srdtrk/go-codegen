package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateInstantiateMsg(f *jen.File, schema *schemas.JSONSchema) {
	generateStructMsg(f, schema, "InstantiateMsg")
}

func GenerateMigrateMsg(f *jen.File, schema *schemas.JSONSchema) {
	generateStructMsg(f, schema, "MigrateMsg")
}

func generateStructMsg(f *jen.File, schema *schemas.JSONSchema, allowedTitle string) {
	if schema == nil {
		return
	}

	if err := validateAsStructMsg(schema, allowedTitle); err != nil {
		panic(err)
	}

	f.Comment(schema.Description)
	f.Type().Id(schema.Title).Struct(
		GenerateFieldsFromProperties(schema.Properties)...,
	)

	RegisterDefinitions(schema.Definitions)
}

func validateAsStructMsg(schema *schemas.JSONSchema, allowedTitle string) error {
	if schema.Title != allowedTitle {
		return fmt.Errorf("title must be InstantiateMsg")
	}
	if schema.Type[0] != "object" {
		return fmt.Errorf("InstantiateMsg type must be object")
	}

	return nil
}
