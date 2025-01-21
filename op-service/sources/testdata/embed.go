package testdata

import "embed"

//go:embed data/**/*.json
var TestDataFiles embed.FS
