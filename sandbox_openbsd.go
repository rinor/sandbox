//go:build openbsd && !nanos

package sandbox

import (
	"golang.org/x/sys/unix"
)

const (
	noop = false
)

func unveil(path string, flags string) error {
	return unix.Unveil(path, flags)
}

func unveilBlock() error {
	return unix.UnveilBlock()
}

func pledge(promises, execpromises string) error {
	return unix.Pledge(promises, execpromises)
}

func pledgePromises(promises string) error {
	return unix.PledgePromises(promises)
}

func pledgeExecpromises(execpromises string) error {
	return unix.Execpromises(execpromises)
}
