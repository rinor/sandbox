//go:build nanos && !openbsd

package sandbox

import (
	"syscall"
	"unsafe"
)

const (
	// OpenBSD syscalls, mapped to unused syscall numbers in Linux
	nanos_sys_pledge = 335
	nanos_sys_unveil = 336

	noop = false
)

func pledge(promises, execpromises string) error {
	return pledgePromises(promises)
}

func pledgePromises(promises string) error {
	// This variable holds the execpromises and is always nil.
	var exptr unsafe.Pointer
	pptr, err := syscall.BytePtrFromString(promises)
	if err != nil {
		return err
	}
	_, _, e := syscall.Syscall(nanos_sys_pledge, uintptr(unsafe.Pointer(pptr)), uintptr(exptr), 0)
	if e != 0 {
		return e
	}
	return nil
}

func pledgeExecpromises(execpromises string) error {
	return nil
}

func unveil(path string, flags string) error {
	pathPtr, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	flagsPtr, err := syscall.BytePtrFromString(flags)
	if err != nil {
		return err
	}
	_, _, e := syscall.Syscall(nanos_sys_unveil, uintptr(unsafe.Pointer(pathPtr)), uintptr(unsafe.Pointer(flagsPtr)), 0)
	if e != 0 {
		return e
	}
	return nil
}

func unveilBlock() error {
	// Both pointers must be nil.
	var pathUnsafe, flagsUnsafe unsafe.Pointer
	_, _, e := syscall.Syscall(nanos_sys_unveil, uintptr(pathUnsafe), uintptr(flagsUnsafe), 0)
	if e != 0 {
		return e
	}
	return nil
}
