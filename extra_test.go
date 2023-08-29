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
	Add(T, T) T
	Double(T) T
	Negate(T) T
	ScalarMult(T, []byte) (T, error)
	ScalarBaseMult([]byte) (T, error)
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
