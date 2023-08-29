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
