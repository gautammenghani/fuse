package fuse

// Force init order here.

func init() {
	initLogger()
	initForcedBackend()
}
