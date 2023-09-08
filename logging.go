package fuse

import (
	"github.com/anacrolix/log"
)

var Logger log.Logger

func initLogger() {
	Logger = log.Default.WithContextText("fuse").WithNames("fuse")
}
