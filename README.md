# filippo.io/nistec

```
import "filippo.io/nistec"
```

This package implements the NIST P elliptic curves, according to FIPS 186-4
and SEC 1, Version 2.0, exposing the necessary APIs to build a wide array of
higher-level primitives.

It's an exported version of `crypto/internal/nistec` in the standard library,
which powers `crypto/elliptic`, `crypto/ecdsa`, and `crypto/ecdh`.
The git history has been preserved, and new upstream changes are applied periodically.

This package uses fiat-crypto or specialized assembly and Go code for its
backend field arithmetic (not math/big) and exposes constant-time, heap
allocation-free, byte slice-based safe APIs. Group operations use modern and
safe complete addition formulas where possible. The point at infinity is
handled and encoded according to SEC 1, Version 2.0, and invalid curve points
can't be represented. This makes it particularly suitable to be used as a
prime order group implementation.

Use the `purego` build tag to exclude the assembly and rely entirely on formally
verified fiat-crypto arithmetic and complete addition formulas.

Read the docs at [pkg.go.dev/filippo.io/nistec](https://pkg.go.dev/filippo.io/nistec).

This repository does not accept contributions.
Any changes should be submitted upstream to the Go project.
