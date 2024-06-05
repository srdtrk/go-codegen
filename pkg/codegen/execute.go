package codegen

import (
	"fmt"
	"slices"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
)

// Generates the code for ExecuteMsg
func GenerateExecuteMsg(f *jen.File, schema *schemas.JSONSchema) {
	if schema == nil {
		types.DefaultLogger().Info().Msg("No ExecuteMsg found. Skipping...")
		return
	}

	generateEnumMsg(f, schema, []string{"ExecuteMsg", "ExecuteMsg_for_Empty"})
}

// Generates the code for SudoMsg
func GenerateSudoMsg(f *jen.File, schema *schemas.JSONSchema) {
	if schema == nil {
		types.DefaultLogger().Info().Msg("No SudoMsg found. Skipping...")
		return
	}

	generateEnumMsg(f, schema, []string{"SudoMsg", "SudoMsg_for_Empty"})
}

// Generates the code for QueryMsg
func GenerateQueryMsg(f *jen.File, schema *schemas.JSONSchema) {
	if schema == nil {
		types.DefaultLogger().Info().Msg("No QueryMsg found. Skipping...")
		return
	}

	generateEnumMsg(f, schema, []string{"QueryMsg", "QueryMsg_for_Empty"})
}

func generateEnumMsg(f *jen.File, schema *schemas.JSONSchema, allowedTitles []string) {
	if schema == nil {
		panic(fmt.Errorf("schema of %s is nil", allowedTitles))
	}

	if err := validateAsEnumMsg(schema, allowedTitles); err != nil {
		panic(err)
	}

	f.Comment(schema.Description)
	f.Type().Id(schema.Title).Struct(
		generateFieldsFromOneOf(schema.OneOf, schema.Title+"_")...,
	)

	RegisterDefinitions(schema.Definitions)
}

func validateAsEnumMsg(schema *schemas.JSONSchema, allowedTitles []string) error {
	if !slices.Contains(allowedTitles, schema.Title) {
		return fmt.Errorf("title %s is not one of %s", schema.Title, allowedTitles)
	}
	if len(schema.OneOf) == 0 {
		types.DefaultLogger().Warn().Msg(fmt.Sprintf("%s is an empty enum", schema.Title))
	}

	return nil
}
