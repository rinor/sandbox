//go:build linux

package sandbox

import (
	"syscall"
	"unsafe"
)

const (
	// OpenBSD syscalls, mapped to unused syscall numbers in Linux
	nanos_sys_pledge = 335
	nanos_sys_unveil = 336

	// uname -s
	nanos_sysname = "Nanos"
)

var (
	noop = true
)

func init() {
	var uts syscall.Utsname
	if err := syscall.Uname(&uts); err != nil {
		return
	}

	int8ToString := func(s []int8) string {
		b := make([]byte, 0, len(s))
		for _, v := range s {
			if v == 0x00 {
				break
			}
			b = append(b, byte(v))
		}
		return string(b)
	}

	noop = int8ToString(uts.Sysname[:]) != nanos_sysname
}

func pledge(promises, execpromises string) error {
	return pledgePromises(promises)
}

func pledgePromises(promises string) error {
	if noop {
		return nil
	}
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
	if noop {
		return nil
	}
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
	if noop {
		return nil
	}
	// Both pointers must be nil.
	var pathUnsafe, flagsUnsafe unsafe.Pointer
	_, _, e := syscall.Syscall(nanos_sys_unveil, uintptr(pathUnsafe), uintptr(flagsUnsafe), 0)
	if e != 0 {
		return e
	}
	return nil
}
