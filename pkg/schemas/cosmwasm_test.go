package schemas_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func TestIDLSchemaFromFile(t *testing.T) {
	_, err := schemas.IDLSchemaFromFile("testdata/cw-ica-controller.json")
	require.NoError(t, err)
}
