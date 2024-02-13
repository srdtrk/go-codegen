package codegen

import (
	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
)

func GenerateResponses(f *jen.File, responses map[string]*schemas.JSONSchema) {
	for _, schema := range responses {
		title := schema.Title
		if title == "" {
			panic("query response must have a title")
		}
		duplicate, found := GetDefinition(schema.Title)
		if found {
			if duplicate.Title != schema.Title || duplicate.Description != schema.Description {
				types.DefaultLogger().Warn().Msgf("found duplicate definition `%s` with differing implementations", schema.Title)
				types.DefaultLogger().Warn().Msgf("renaming the duplicate definition to `%s`", schema.Title+"_2")

				schema.Title += "_2"
			}
		}

		RegisterDefinition(schema.Title, schema)

		if len(schema.Definitions) != 0 {
			RegisterDefinitions(schema.Definitions)
		}
	}
}
