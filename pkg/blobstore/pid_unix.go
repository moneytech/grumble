// Copyright (c) 2011 The Grumble Authors
// The use of this source code is goverened by a BSD-style
// license that can be found in the LICENSE-file.

package blobstore

import (
	"syscall"
)

func getPid() uint64 {
	return uint64(syscall.Getpid())
}

func pidRunning(pid uint64) bool {
	return syscall.Kill(int(pid), 0) == 0
}
