# sandbox

Package **sandbox** is a wrapper around OpenBSD's pledge(2) and unveil(2) system calls.

Can be used with [Nanos](https://github.com/nanovms/nanos) `sandbox` _klib_ - https://docs.ops.city/ops/klibs#sandbox

Can be safely used on other non-OpenBSD operating systems, where the syscalls are `noop`.