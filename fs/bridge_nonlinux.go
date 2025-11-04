//go:build !linux

package fs

import "github.com/4831c0/go-fuse/v2/fuse"

func (b *rawBridge) Statx(cancel <-chan struct{}, in *fuse.StatxIn, out *fuse.StatxOut) fuse.Status {
	return fuse.ENOSYS
}
