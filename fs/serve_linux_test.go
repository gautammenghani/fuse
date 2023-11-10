package fs_test

import (
	"io"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func platformStatfs(st *syscall.Statfs_t) *statfsResult {
	return &statfsResult{
		Blocks:  st.Blocks,
		Bfree:   st.Bfree,
		Bavail:  st.Bavail,
		Files:   st.Files,
		Ffree:   st.Ffree,
		Bsize:   int64(st.Bsize),
		Namelen: int64(st.Namelen),
		Frsize:  int64(st.Frsize),
	}
}

func platformStat(fi os.FileInfo) *statResult {
	r := &statResult{
		Mode: fi.Mode(),
	}
	st := fi.Sys().(*syscall.Stat_t)
	r.Ino = st.Ino
	r.Nlink = uint64(st.Nlink)
	r.UID = st.Uid
	r.GID = st.Gid
	r.Blksize = int64(st.Blksize)
	return r
}

var _lockOFDHelper = helpers.Register("lock-ofd", &lockHelp{
	lockFn: func(fd uintptr, req *lockReq) error {
		lk := unix.Flock_t{
			Type:   unix.F_WRLCK,
			Whence: int16(io.SeekStart),
			Start:  req.Start,
			Len:    req.Len,
		}
		cmd := unix.F_OFD_SETLK
		if req.Wait {
			cmd = unix.F_OFD_SETLKW
		}
		return unix.FcntlFlock(fd, cmd, &lk)
	},
	unlockFn: func(fd uintptr, req *lockReq) error {
		lk := unix.Flock_t{
			Type:   unix.F_UNLCK,
			Whence: int16(io.SeekStart),
			Start:  req.Start,
			Len:    req.Len,
		}
		cmd := unix.F_OFD_SETLK
		if req.Wait {
			cmd = unix.F_OFD_SETLKW
		}
		return unix.FcntlFlock(fd, cmd, &lk)
	},
	queryFn: func(fd uintptr, lk *unix.Flock_t) error {
		cmd := unix.F_OFD_GETLK
		return unix.FcntlFlock(fd, cmd, lk)
	},
})

func init() {
	lockOFDHelper = _lockOFDHelper
}
