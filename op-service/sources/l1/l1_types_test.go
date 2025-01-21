package l1

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zircuit-labs/zkr-monorepo-public/op-service/sources/testdata"
)

type testMetadata struct {
	Name   string `json:"name"`
	Fail   bool   `json:"fail,omitempty"`
	Reason string `json:"reason,omitempty"`
}

func readJsonTestdata(t *testing.T, name string, dest any) {
	f, err := testdata.TestDataFiles.Open(name)
	require.NoError(t, err, "must open %q", name)
	require.NoError(t, json.NewDecoder(f).Decode(dest), "must json-decode %q", name)
	require.NoError(t, f.Close(), "must close %q", name)
}

func TestBlockHeaderJSON(t *testing.T) {
	headersDir, err := testdata.TestDataFiles.ReadDir("data/headers")
	require.NoError(t, err)

	for _, entry := range headersDir {
		if !strings.HasSuffix(entry.Name(), "_metadata.json") {
			continue
		}

		var metadata testMetadata
		readJsonTestdata(t, "data/headers/"+entry.Name(), &metadata)
		t.Run(metadata.Name, func(t *testing.T) {
			var header RPCHeader
			readJsonTestdata(t, "data/headers/"+strings.Replace(entry.Name(), "_metadata.json", "_data.json", 1), &header)

			h := header.computeBlockHash()
			if metadata.Fail {
				require.NotEqual(t, h, header.Hash, "expecting verification error")
			} else {
				require.Equal(t, h, header.Hash, "blockhash should verify ok")
			}
		})
	}
}

func TestBlockJSON(t *testing.T) {
	blocksDir, err := testdata.TestDataFiles.ReadDir("data/blocks")
	require.NoError(t, err)

	for _, entry := range blocksDir {
		if !strings.HasSuffix(entry.Name(), "_metadata.json") {
			continue
		}

		var metadata testMetadata
		readJsonTestdata(t, "data/blocks/"+entry.Name(), &metadata)
		t.Run(metadata.Name, func(t *testing.T) {
			var block RPCBlock
			readJsonTestdata(t, "data/blocks/"+strings.Replace(entry.Name(), "_metadata.json", "_data.json", 1), &block)

			err := block.verify()
			if metadata.Fail {
				require.NotNil(t, err, "expecting verification error")
				require.ErrorContains(t, err, metadata.Reason, "validation failed for incorrect reason")
			} else {
				require.NoError(t, err, "verification should pass")
			}
		})
	}
}
