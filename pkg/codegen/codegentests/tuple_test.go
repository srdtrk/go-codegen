package codegentests_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuple(t *testing.T) {
	t.Parallel()

	t.Run("TestTuple", func(t *testing.T) {
		basicTuple := Tuple_of_ClassId_and_TokenId{
			F0: ClassId("myClassId"),
			F1: TokenId("myTokenId"),
		}

		jsonBz, err := json.Marshal(basicTuple)
		require.NoError(t, err)
		require.Equal(t, `["myClassId","myTokenId"]`, string(jsonBz))

		var unmarshalledTuple Tuple_of_ClassId_and_TokenId
		err = json.Unmarshal(jsonBz, &unmarshalledTuple)
		require.NoError(t, err)
		require.Equal(t, basicTuple, unmarshalledTuple)
	})
}
