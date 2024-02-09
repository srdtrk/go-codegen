package schemas

// QueryMsgSchema represents a JSON Schema object type for ExecuteMsg.
// based on https://github.com/CosmWasm/ts-codegen/blob/eeba0cf1e0cbadfb6cfc103412108e763fdb9f6a/packages/wasm-ast-types/types/types.d.ts#L18
type QueryMsgSchema struct {
	Schema string        `json:"$schema,omitempty"` // Section 6.1.
	AllOf  []*JSONSchema `json:"allOf,omitempty"`   // Section 5.22.
	AnyOf  []*JSONSchema `json:"anyOf,omitempty"`   // Section 5.23.
	OneOf  []*JSONSchema `json:"oneOf,omitempty"`   // Section 5.24.
	// Title is "QueryMsg"
	Title string `json:"title,omitempty"` // Section 6.1.
}
