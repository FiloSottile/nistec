// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fiat

// Bits returns a reference to the underlying little-endian fully-reduced
// Montgomery representation of e. Handle with care.
func (e *P256Element) Bits() *[4]uint64 {
	var _ p256MontgomeryDomainFieldElement = e.x
	return (*[4]uint64)(&e.x)
}
