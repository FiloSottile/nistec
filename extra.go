// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nistec

import "filippo.io/nistec/internal/fiat"

// Negate sets p = -q and returns p.
func (p *P224Point) Negate(q *P224Point) *P224Point {
	p.x.Set(q.x)
	p.y.Sub(new(fiat.P224Element), q.y)
	p.z.Set(q.z)
	return p
}

// Negate sets p = -q and returns p.
func (p *P384Point) Negate(q *P384Point) *P384Point {
	p.x.Set(q.x)
	p.y.Sub(new(fiat.P384Element), q.y)
	p.z.Set(q.z)
	return p
}

// Negate sets p = -q and returns p.
func (p *P521Point) Negate(q *P521Point) *P521Point {
	p.x.Set(q.x)
	p.y.Sub(new(fiat.P521Element), q.y)
	p.z.Set(q.z)
	return p
}

// IsInfinity returns 1 if p is the point-at-infinity, 0 otherwise.
func (p *P224Point) IsInfinity() int {
	return p.z.IsZero()
}

// IsInfinity returns 1 if p is the point-at-infinity, 0 otherwise.
func (p *P384Point) IsInfinity() int {
	return p.z.IsZero()
}

// IsInfinity returns 1 if p is the point-at-infinity, 0 otherwise.
func (p *P521Point) IsInfinity() int {
	return p.z.IsZero()
}

// Equal returns 1 if p and q represent the same point, 0 otherwise.
func (p *P224Point) Equal(q *P224Point) int {
	pinf := p.z.IsZero()
	qinf := q.z.IsZero()
	bothinf := pinf & qinf
	noneinf := (1 - pinf) & (1 - qinf)
	px := new(fiat.P224Element).Mul(p.x, q.z)
	qx := new(fiat.P224Element).Mul(q.x, p.z)
	py := new(fiat.P224Element).Mul(p.y, q.z)
	qy := new(fiat.P224Element).Mul(q.y, p.z)
	return bothinf | (noneinf & px.Equal(qx) & py.Equal(qy))
}

// Equal returns 1 if p and q represent the same point, 0 otherwise.
func (p *P384Point) Equal(q *P384Point) int {
	pinf := p.z.IsZero()
	qinf := q.z.IsZero()
	bothinf := pinf & qinf
	noneinf := (1 - pinf) & (1 - qinf)
	px := new(fiat.P384Element).Mul(p.x, q.z)
	qx := new(fiat.P384Element).Mul(q.x, p.z)
	py := new(fiat.P384Element).Mul(p.y, q.z)
	qy := new(fiat.P384Element).Mul(q.y, p.z)
	return bothinf | (noneinf & px.Equal(qx) & py.Equal(qy))
}

// Equal returns 1 if p and q represent the same point, 0 otherwise.
func (p *P521Point) Equal(q *P521Point) int {
	pinf := p.z.IsZero()
	qinf := q.z.IsZero()
	bothinf := pinf & qinf
	noneinf := (1 - pinf) & (1 - qinf)
	px := new(fiat.P521Element).Mul(p.x, q.z)
	qx := new(fiat.P521Element).Mul(q.x, p.z)
	py := new(fiat.P521Element).Mul(p.y, q.z)
	qy := new(fiat.P521Element).Mul(q.y, p.z)
	return bothinf | (noneinf & px.Equal(qx) & py.Equal(qy))
}
