// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64

package nistec

// This file exports the P256OrdInverse function so it's accessible during tests
// from the unmodified p256_asm_ordinv_test.go from the stdlib, but not as part
// of the public API of filippo.io/nistec.

func P256OrdInverse(k []byte) ([]byte, error) {
	return p256OrdInverse(k)
}
