// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && (amd64 || arm64 || (ppc64le && go1.19) || s390x)

package nistec

import "filippo.io/nistec/internal/fiat"

// Negate sets p = -q and returns p.
func (p *P256Point) Negate(q *P256Point) *P256Point {
	// fiat.P256Element is a little-endian Montgomery domain fully-reduced
	// element, like p256Element, so they are actually interchangable.
	qy := new(fiat.P256Element)
	*qy.Bits() = q.y
	py := new(fiat.P256Element).Sub(new(fiat.P256Element), qy)

	p.x = q.x
	p.y = *py.Bits()
	p.z = q.z
	return p
}
