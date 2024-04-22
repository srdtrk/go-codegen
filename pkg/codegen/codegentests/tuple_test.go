package codegentests_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuple(t *testing.T) {
	t.Parallel()

	t.Run("TestBasicTuple", func(t *testing.T) {
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

	t.Run("TestNestedTuple", func(t *testing.T) {
		nestedTuple := Tuple_of_Tuple_of_ClassId_and_TokenId_and_string{
			F0: Tuple_of_ClassId_and_TokenId{
				F0: ClassId("myClassId"),
				F1: TokenId("myTokenId"),
			},
			F1: "hello",
		}

		jsonBz, err := json.Marshal(nestedTuple)
		require.NoError(t, err)
		require.Equal(t, `[["myClassId","myTokenId"],"hello"]`, string(jsonBz))

		var unmarshalledTuple Tuple_of_Tuple_of_ClassId_and_TokenId_and_string
		err = json.Unmarshal(jsonBz, &unmarshalledTuple)
		require.NoError(t, err)
		require.Equal(t, nestedTuple, unmarshalledTuple)
	})
}
