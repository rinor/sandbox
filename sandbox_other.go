//go:build !(openbsd || linux)

package sandbox

const (
	noop = true
)

func unveil(path string, flags string) error {
	return nil
}

func unveilBlock() error {
	return nil
}

func pledge(promises, execpromises string) error {
	return nil
}

func pledgePromises(promises string) error {
	return nil
}

func pledgeExecpromises(execpromises string) error {
	return nil
}
