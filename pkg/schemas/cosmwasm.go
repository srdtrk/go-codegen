package schemas

import (
	"encoding/json"
	"fmt"
	"os"
)

// IDLSchema is the schema produced by cosmwasm_schema::write_api
type IDLSchema struct {
	ContractName    string                 `json:"contract_name"`
	ContractVersion string                 `json:"contract_version"`
	IdlVersion      string                 `json:"idl_version"`
	Instantiate     *JSONSchema            `json:"instantiate"`
	Execute         *JSONSchema            `json:"execute"`
	Query           *JSONSchema            `json:"query"`
	Migrate         *JSONSchema            `json:"migrate"`
	Sudo            *JSONSchema            `json:"sudo"`
	Responses       map[string]*JSONSchema `json:"responses"`
}

// IDLSchemaFromFile reads a JSON file and returns the IDL schema
func IDLSchemaFromFile(filePath string) (*IDLSchema, error) {
	schemaFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var idlSchema IDLSchema
	err = json.Unmarshal(schemaFile, &idlSchema)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %w", err)
	}

	return &idlSchema, nil
}
