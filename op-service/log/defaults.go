package log

import (
	"os"

	"github.com/zircuit-labs/l2-geth-public/log"
)

func SetupDefaults() {
	SetGlobalLogHandler(log.LogfmtHandlerWithLevel(os.Stdout, log.LevelInfo))
}
