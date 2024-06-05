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

	// MigrateMsg can be either a struct or an enum
	generateStructOrEnumMsg(f, schema, "MigrateMsg")
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
		return fmt.Errorf("title must be %v", allowedTitle)
	}

	if schema.Type == nil {
		return fmt.Errorf("%v type must be defined", allowedTitle)
	}

	if schema.Type[0] != "object" {
		return fmt.Errorf("%v type must be object", allowedTitle)
	}

	return nil
}

// Generate the msg regardless if it is a struct or enum this is mostly needed for migration msgs
func generateStructOrEnumMsg(f *jen.File, schema *schemas.JSONSchema, allowedTitle string) {
	if schema == nil {
		panic(fmt.Errorf("schema of %s is nil", allowedTitle))
	}

	err := validateAsStructMsg(schema, allowedTitle)
	if err == nil {
		generateStructMsg(f, schema, allowedTitle)
	} else {
		generateEnumMsg(f, schema, []string{allowedTitle, allowedTitle + "_for_Empty"})
	}
}
