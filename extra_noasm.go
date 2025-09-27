// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build purego || (!amd64 && !arm64 && !(ppc64le && go1.19) && !s390x)

package nistec

import "filippo.io/nistec/internal/fiat"

// Negate sets p = -q and returns p.
func (p *P256Point) Negate(q *P256Point) *P256Point {
	p.x.Set(q.x)
	p.y.Sub(new(fiat.P256Element), q.y)
	p.z.Set(q.z)
	return p
}

// IsInfinity returns 1 if p is the point-at-infinity, 0 otherwise.
func (p *P256Point) IsInfinity() int {
	return p.z.IsZero()
}

// Equal returns 1 if p and q represent the same point, 0 otherwise.
func (p *P256Point) Equal(q *P256Point) int {
	pinf := p.z.IsZero()
	qinf := q.z.IsZero()
	bothinf := pinf & qinf
	noneinf := (1 - pinf) & (1 - qinf)
	px := new(fiat.P256Element).Mul(p.x, q.z)
	qx := new(fiat.P256Element).Mul(q.x, p.z)
	py := new(fiat.P256Element).Mul(p.y, q.z)
	qy := new(fiat.P256Element).Mul(q.y, p.z)
	return bothinf | (noneinf & px.Equal(qx) & py.Equal(qy))
}
