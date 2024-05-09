package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
)

func GenerateInstantiateMsg(f *jen.File, schema *schemas.JSONSchema) {
	if schema == nil {
		types.DefaultLogger().Warn().Msg("No InstantiateMsg found. Skipping...")
		return
	}

	generateStructMsg(f, schema, "InstantiateMsg")
}

func GenerateMigrateMsg(f *jen.File, schema *schemas.JSONSchema) {
	if schema == nil {
		types.DefaultLogger().Info().Msg("No MigrateMsg found. Skipping...")
		return
	}

	generateStructMsg(f, schema, "MigrateMsg")
}

func generateStructMsg(f *jen.File, schema *schemas.JSONSchema, allowedTitle string) {
	if schema == nil {
		panic(fmt.Errorf("schema of %s is nil", allowedTitle))
	}

	if err := validateAsStructMsg(schema, allowedTitle); err != nil {
		panic(err)
	}

	f.Comment(schema.Description)
	f.Type().Id(schema.Title).Struct(
		generateFieldsFromProperties(schema.Properties, true)...,
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
