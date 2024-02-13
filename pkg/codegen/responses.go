package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
	"github.com/srdtrk/go-codegen/pkg/types"
)

func GenerateResponses(f *jen.File, responses map[string]*schemas.JSONSchema) {
	for key, schema := range responses {
		if len(schema.Type) != 1 {
			panic(fmt.Sprintf("response schema for %s must have exactly one type, got %d", key, len(schema.Type)))
		}
		switch schema.Type[0] {
		case schemas.TypeNameObject:
			title := schema.Title
			if title == "" {
				panic(fmt.Sprintf("response schema for %s must have a title", key))
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
		case schemas.TypeNameArray:
			if len(schema.Definitions) == 0 {
				panic(fmt.Sprintf("response schema for %s must have a definition for array", key))
			}

			RegisterDefinitions(schema.Definitions)
		case schemas.TypeNameString:
			// Do nothing
		case schemas.TypeNameNumber:
			// Do nothing
		case schemas.TypeNameInteger:
			// Do nothing
		case schemas.TypeNameBoolean:
			// Do nothing
		case schemas.TypeNameNull:
			types.DefaultLogger().Warn().Msgf("response schema for %s is of type null", key)
		default:
			types.DefaultLogger().Error().Msgf("response schema for %s is of unknown type %s", key, schema.Type[0])
		}

		if len(schema.Definitions) != 0 {
			RegisterDefinitions(schema.Definitions)
		}
	}
}
