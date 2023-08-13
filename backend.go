package fuse

import (
	"github.com/anacrolix/log"
	"os"
	"strings"
)

type Backend string

const (
	fuseTBackend   = "FUSE-T"
	osxfuseBackend = "OSXFUSE"
)

func (be Backend) IsFuseT() bool {
	return be == fuseTBackend
}

func (be Backend) IsUnset() bool {
	return be == ""
}

var forcedBackend = func() (ret Backend) {
	ret = Backend(strings.ToUpper(strings.TrimSpace(os.Getenv("FUSE_FORCE_BACKEND"))))
	if !ret.IsUnset() {
		Logger.Levelf(log.Info, "forcing backend %v", ret)
	}
	return
}()
