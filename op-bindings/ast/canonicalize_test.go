package ast

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/zkr-monorepo-public/op-bindings/solc"
)

func TestCanonicalize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		filename string
	}{
		{
			"simple",
			"simple.json",
		},
		{
			"remap public variables",
			"public-variables.json",
		},
		{
			"values in storage",
			"values-in-storage.json",
		},
		{
			"custom types",
			"custom-types.json",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f, err := os.Open(path.Join("testdata", tt.filename))
			require.NoError(t, err)
			dec := json.NewDecoder(f)
			var testData struct {
				In  *solc.StorageLayout `json:"in"`
				Out *solc.StorageLayout `json:"out"`
			}
			require.NoError(t, dec.Decode(&testData))
			require.NoError(t, f.Close())

			// Run 100 times to make sure that we aren't relying
			// on random map iteration order.
			for i := 0; i < 100; i++ {
				require.Equal(t, testData.Out, CanonicalizeASTIDs(testData.In, ""))
			}
		})
	}
}
