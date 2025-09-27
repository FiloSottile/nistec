// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && (amd64 || arm64 || ppc64le || s390x)

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

// IsInfinity returns 1 if p is the point-at-infinity, 0 otherwise.
func (p *P256Point) IsInfinity() int {
	return p.isInfinity()
}

// Equal returns 1 if p and q represent the same point, 0 otherwise.
func (p *P256Point) Equal(q *P256Point) int {
	pinf := p256Equal(&p.z, &p256Zero)
	qinf := p256Equal(&q.z, &p256Zero)
	bothinf := pinf & qinf
	noneinf := (1 - pinf) & (1 - qinf)

	// xp = Xp / Zp²
	// yp = Yp / Zp³
	// xq = Xq / Zq²
	// yq = Yq / Zq³
	// If Zp != 0 and Zq != 0, then:
	//    xp == yp  <=>  Xp*Zq² == Xq*Zp²
	//    xq == yq  <=>  Yp*Zq³ == Yq*Zp³
	px := new(p256Element)
	qx := new(p256Element)
	py := new(p256Element)
	qy := new(p256Element)
	pz := new(p256Element)
	qz := new(p256Element)
	p256Sqr(pz, &p.z, 1)
	p256Sqr(qz, &q.z, 1)
	p256Mul(px, &p.x, qz)
	p256Mul(qx, &q.x, pz)
	samex := p256Equal(px, qx)
	p256Mul(pz, pz, &p.z)
	p256Mul(qz, qz, &q.z)
	p256Mul(py, &p.y, qz)
	p256Mul(qy, &q.y, pz)
	samey := p256Equal(py, qy)
	return bothinf | (noneinf & samex & samey)
}
