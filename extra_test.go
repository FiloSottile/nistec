// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nistec_test

import (
	"bytes"
	"crypto/elliptic"
	"math/big"
	"testing"

	"filippo.io/nistec"
)

type nistPointExtra[T any] interface {
	Bytes() []byte
	SetGenerator() T
	SetBytes([]byte) (T, error)
	Set(T) T
	Add(T, T) T
	Double(T) T
	Negate(T) T
	ScalarMult(T, []byte) (T, error)
	ScalarBaseMult([]byte) (T, error)
	IsInfinity() int
	Equal(T) int
}

func TestNegate(t *testing.T) {
	t.Run("P224", func(t *testing.T) {
		testNegate(t, nistec.NewP224Point, elliptic.P224())
	})
	t.Run("P256", func(t *testing.T) {
		testNegate(t, nistec.NewP256Point, elliptic.P256())
	})
	t.Run("P384", func(t *testing.T) {
		testNegate(t, nistec.NewP384Point, elliptic.P384())
	})
	t.Run("P521", func(t *testing.T) {
		testNegate(t, nistec.NewP521Point, elliptic.P521())
	})
}

func testNegate[P nistPointExtra[P]](t *testing.T, newPoint func() P, c elliptic.Curve) {
	p := newPoint().SetGenerator()
	p.Add(p, p) // make a test point with z != 1

	p1 := newPoint().Negate(p)

	minusOne := new(big.Int).Sub(c.Params().N, big.NewInt(1)).Bytes()
	p2, err := newPoint().ScalarMult(p, minusOne)
	fatalIfErr(t, err)
	if !bytes.Equal(p1.Bytes(), p2.Bytes()) {
		t.Error("-P != [-1]P")
	}

	p.Negate(p)
	if !bytes.Equal(p.Bytes(), p1.Bytes()) {
		t.Error("-P (aliasing) != -P")
	}
}

func TestEqual(t *testing.T) {
	t.Run("P224", func(t *testing.T) {
		testEqual(t, nistec.NewP224Point, elliptic.P224())
	})
	t.Run("P256", func(t *testing.T) {
		testEqual(t, nistec.NewP256Point, elliptic.P256())
	})
	t.Run("P384", func(t *testing.T) {
		testEqual(t, nistec.NewP384Point, elliptic.P384())
	})
	t.Run("P521", func(t *testing.T) {
		testEqual(t, nistec.NewP521Point, elliptic.P521())
	})
}

func testEqual[P nistPointExtra[P]](t *testing.T, newPoint func() P, c elliptic.Curve) {
	inf := newPoint()
	g := newPoint().SetGenerator()
	g.Add(g, g)
	if inf.Equal(inf) != 1 {
		t.Error("inf != inf")
	}
	if g.Equal(g) != 1 {
		t.Error("G != G")
	}
	if inf.Equal(g) != 0 || g.Equal(inf) != 0 {
		t.Error("G == inf")
	}
	p1 := newPoint().Set(g)
	p2 := newPoint()
	p3 := newPoint()
	sc := make([]byte, (c.Params().N.BitLen()+7)>>3)
	for i := 1; i < 30; i++ {
		sc[len(sc)-1] = byte(i)
		_, err := p2.ScalarMult(g, sc)
		fatalIfErr(t, err)
		sc[len(sc)-1] = byte(i * 2)
		_, err = p3.ScalarBaseMult(sc)
		fatalIfErr(t, err)
		if p1.Equal(p2) != 1 || p2.Equal(p1) != 1 {
			t.Error("n*G != n*G (1)")
		}
		if p1.Equal(p3) != 1 || p3.Equal(p1) != 1 {
			t.Error("n*G != n*G (2)")
		}
		if p2.Equal(p3) != 1 || p3.Equal(p2) != 1 {
			t.Error("n*G != n*G (3)")
		}
		p1.Add(p1, g)
		if p1.Equal(g) != 0 || g.Equal(p1) != 0 {
			t.Error("n*G == G")
		}
	}
}

func TestInfinity(t *testing.T) {
	t.Run("P224", func(t *testing.T) {
		testInfinity(t, nistec.NewP224Point, elliptic.P224())
	})
	t.Run("P256", func(t *testing.T) {
		testInfinity(t, nistec.NewP256Point, elliptic.P256())
	})
	t.Run("P384", func(t *testing.T) {
		testInfinity(t, nistec.NewP384Point, elliptic.P384())
	})
	t.Run("P521", func(t *testing.T) {
		testInfinity(t, nistec.NewP521Point, elliptic.P521())
	})
}

func testInfinity[P nistPointExtra[P]](t *testing.T, newPoint func() P, c elliptic.Curve) {
	inf := newPoint()
	g := newPoint().SetGenerator()
	g.Add(g, g)
	if inf.IsInfinity() != 1 {
		t.Error("inf != inf")
	}
	if g.IsInfinity() != 0 {
		t.Error("G == inf")
	}
	p1 := newPoint().Set(g)
	p2 := newPoint()
	sc := make([]byte, (c.Params().N.BitLen()+7)>>3)
	for i := 1; i < 30; i++ {
		sc[len(sc)-1] = byte(i)
		_, err := p2.ScalarMult(g, sc)
		fatalIfErr(t, err)
		if p2.IsInfinity() != 0 {
			t.Error("n*G == inf (1)")
		}
		p2.Negate(p2)
		if p2.IsInfinity() != 0 {
			t.Error("n*G == inf (2)")
		}
		p2.Add(p2, p1)
		if p2.IsInfinity() != 1 {
			t.Error("n*G - n*G != inf")
		}
		p1.Add(p1, g)
		if p1.IsInfinity() != 0 {
			t.Error("n*G == inf (3)")
		}
	}
}

func fatalIfErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
