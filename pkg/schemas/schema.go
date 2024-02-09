package schemas

import (
	"encoding/json"
	"fmt"
)

// JSONSchema represents a JSON Schema object type.
// based on https://github.com/CosmWasm/ts-codegen/blob/eeba0cf1e0cbadfb6cfc103412108e763fdb9f6a/packages/wasm-ast-types/types/types.d.ts#L25
type JSONSchema struct {
	Schema               string                 `json:"$schema,omitempty"`              // Section 6.1.
	Ref                  string                 `json:"$ref,omitempty"`                 // Section 7.
	AdditionalItems      *JSONSchema            `json:"additionalItems,omitempty"`      // Section 5.9.
	Items                *JSONSchemaList        `json:"items,omitempty"`                // Section 5.9.
	Required             []string               `json:"required,omitempty"`             // Section 5.15.
	Properties           map[string]*JSONSchema `json:"properties,omitempty"`           // Section 5.16.
	PatternProperties    map[string]*JSONSchema `json:"patternProperties,omitempty"`    // Section 5.17.
	AdditionalProperties *bool                  `json:"additionalProperties,omitempty"` // Section 5.18.
	Type                 string                 `json:"type,omitempty"`                 // Section 5.21.
	AllOf                []*JSONSchema          `json:"allOf,omitempty"`                // Section 5.22.
	AnyOf                []*JSONSchema          `json:"anyOf,omitempty"`                // Section 5.23.
	OneOf                []*JSONSchema          `json:"oneOf,omitempty"`                // Section 5.24.
	Title                string                 `json:"title,omitempty"`                // Section 6.1.
	Description          string                 `json:"description,omitempty"`          // Section 6.1.
	Definitions          map[string]*JSONSchema `json:"definitions,omitempty"`
}

type JSONSchemaList []JSONSchema

// UnmarshalJSON accepts list or single schema.
func (l *JSONSchemaList) UnmarshalJSON(raw []byte) error {
	if len(raw) > 0 && raw[0] == '[' {
		var s []JSONSchema
		if err := json.Unmarshal(raw, &s); err != nil {
			return fmt.Errorf("failed to unmarshal type list: %w", err)
		}

		*l = s

		return nil
	}

	var s JSONSchema
	if err := json.Unmarshal(raw, &s); err != nil {
		return fmt.Errorf("failed to unmarshal type list: %w", err)
	}

	*l = []JSONSchema{s}

	return nil
}
