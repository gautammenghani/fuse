package fuse

import (
	"github.com/anacrolix/log"
)

var Logger = log.Default.WithContextText("fuse").WithNames("fuse")
