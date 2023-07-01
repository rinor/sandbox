# sandbox

Package **sandbox** is a wrapper around OpenBSD's pledge(2) and unveil(2) system calls.

Can be used with [Nanos](https://github.com/nanovms/nanos) `sandbox` _klib_ - https://docs.ops.city/ops/klibs#sandbox

Can be safely used on other non-OpenBSD operating systems, where the syscalls are `noop`.

## Nanos ops config

```json
{
  "Klibs": [
    "sandbox"
  ],
  "ManifestPassthrough": {
    "sandbox": {
      "pledge": {},
      "unveil": {}
    }
  }
}
```

### pledge

```go
package main

import (
	"log"
	"os"

	"github.com/rinor/sandbox"
)

func main() {
	if sandbox.Noop() {
		log.Print("PLEDGE: calls won't have any effect (noop)")
	} else {
		log.Print("PLEDGE: calls will fail if not implemented in kernel")
	}

	log.Print("PLEDGE: calling PledgePromises(stdio error rpath)")
	err := sandbox.PledgePromises("stdio error rpath")
	if err != nil {
		log.Fatalf("PLEDGE: PledgePromises - %q", err)
	}

	log.Print("PLEDGE: Readir should work - (rpath - enabled)")
	_, err = os.ReadDir(".")
	if err != nil {
		log.Fatalf("%q", err)
	}
	log.Print("PLEDGE: Readir OK")

	log.Print("PLEDGE: Disabling rpath")
	err = sandbox.PledgePromises("stdio error")
	if err != nil {
		log.Fatalf("PLEDGE: PledgePromises - %q", err)
	}

	log.Print("PLEDGE: Readir should fail - (rpath - disabled)")
	_, err = os.ReadDir(".")
	if err != nil {
		log.Fatalf("PLEDGE: Readir %q", err)
	}
}
```

### unveil

```go
package main

import (
	"log"
	"os"

	"github.com/rinor/sandbox"
)

func main() {
	if sandbox.Noop() {
		log.Print("UNVEIL: calls won't have any effect (noop)")
	} else {
		log.Print("UNVEIL: calls will fail if not implemented in kernel")
	}

	log.Print("UNVEIL: calling Unveil(/, rwxc)")
	err := sandbox.Unveil("/", "rwxc")
	if err != nil {
		log.Fatalf("UNVEIL: unveil - %q", err)
	}

	log.Print("UNVEIL: Readir should work - (unveil r - enabled)")
	_, err = os.ReadDir("/")
	if err != nil {
		log.Fatalf("%q", err)
	}
	log.Print("UNVEIL: Readir OK")

	log.Print("UNVEIL: Disabling r")
	err = sandbox.Unveil("/", "wxc")
	if err != nil {
		log.Fatalf("UNVEIL: unveil - %q", err)
	}

	log.Print("UNVEIL: Readir should fail - (unveil r - disabled)")
	_, err = os.ReadDir("/")
	if err != nil {
		log.Printf("UNVEIL: Readir %q", err)
	}

	log.Print("UNVEIL: calling unveilBlock")
	err = sandbox.UnveilBlock()
	if err != nil {
		log.Fatalf("UNVEIL: unveilBlock - %q", err)
	}

	log.Print("UNVEIL: unveil calls should fail - (unveilBlock called)")
	err = sandbox.Unveil("/", "rwxc")
	if err != nil {
		log.Fatalf("UNVEIL: unveil - %q", err)
	}
}
```