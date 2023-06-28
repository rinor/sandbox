/*
Package sandbox is a wrapper for OpenBSD's pledge(2) and unveil(2) system calls.
*/
package sandbox

// Unveil ...
func Unveil(path string, flags string) error {
	return unveil(path, flags)
}

// UnveilBlock ...
func UnveilBlock() error {
	return unveilBlock()
}

// Pledge ...
func Pledge(promises, execpromises string) error {
	return pledge(promises, execpromises)
}

// PledgePromises ...
func PledgePromises(promises string) error {
	return pledgePromises(promises)
}

// PledgeExecpromises ...
func PledgeExecpromises(execpromises string) error {
	return pledgeExecpromises(execpromises)
}

// Noop ...
func Noop() bool {
	return noop
}
