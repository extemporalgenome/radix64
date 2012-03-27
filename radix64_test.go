// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package radix64

import (
	"math"
	"testing"
)

const (
	enc = 1 << iota
	dec
	err
)

var equivs = []struct {
	m uint8
	n uint64
	s string
}{
	{enc | dec, 0, "0"},
	{enc | dec, 63, "_"},
	{enc | dec, 64, "10"},
	{enc | dec, 1057222719, "_0_0_"},
	// Padding
	{enc | dec, 0, "000"},
	{enc | dec, 4095, "00__"},
	// Errors
	{enc | dec | err, 4095, "\xff"},
}

var costs = []struct {
	n    uint64
	b, r int
}{
	{0, 1, 6},
	{31, 1, 1},
	{3, 1, 4},
	{63, 1, 0},
	{64, 2, 5},
	{4095, 2, 0},
}

func TestEncDec(t *testing.T) {
	for _, eq := range equivs {
		b := make([]byte, len(eq.s))
		if eq.m&enc != 0 {
			e := Encode(eq.n, b)
			if eq.m&err == 0 {
				if e != nil {
					t.Errorf("Received unexpected error encoding %d: %s\n", eq.n, e)
				} else if string(b) != eq.s {
					t.Errorf("Expected %d to encode into %q; got %q\n", eq.n, eq.s, b)
				}
			} else if e != ErrBufLimReached {
				t.Error("Expected an ErrBufLimReached when encoding", eq.n, "into buffer of length", len(b))
			}
		}
		if eq.m&dec != 0 {
			n, e := Decode([]byte(eq.s))
			if eq.m&err == 0 {
				if e != nil {
					t.Errorf("Received unexpected error decoding %q\n", e)
				} else if n != eq.n {
					t.Errorf("Expected %q to decode into %d; got %d\n", eq.s, eq.n, n)
				}
			} else if e != ErrInvalidByte {
				t.Errorf("Expected an ErrInvalidByte when decoding %q\n", eq.s)
			}
		}
	}
}

func TestCost(t *testing.T) {
	for _, c := range costs {
		if b, r := Cost(c.n); b != c.b || r != c.r {
			t.Errorf("Expected cost of %d to be (%d, %d); got (%d, %d)\n", c.n, c.b, c.r, b, r)
		}
	}
}

func BenchmarkEncSmall(b *testing.B) {
	buf := make([]byte, 1)
	for i := 0; i < b.N; i++ {
		Encode(63, buf)
	}
}

func BenchmarkEncMedium(b *testing.B) {
	buf := make([]byte, 5)
	n := uint64(math.Pow(64, 5)) - 1
	for i := 0; i < b.N; i++ {
		Encode(n, buf)
	}
}

func BenchmarkEncLarge(b *testing.B) {
	buf := make([]byte, 10)
	n := uint64(math.Pow(64, 10)) - 1
	for i := 0; i < b.N; i++ {
		Encode(n, buf)
	}
}

func BenchmarkEncPadding(b *testing.B) {
	buf := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		Encode(0, buf)
	}
}

func BenchmarkDecSmall(b *testing.B) {
	buf := []byte("_")
	for i := 0; i < b.N; i++ {
		Decode(buf)
	}
}

func BenchmarkDecMedium(b *testing.B) {
	buf := []byte("_____")
	for i := 0; i < b.N; i++ {
		Decode(buf)
	}
}

func BenchmarkDecLarge(b *testing.B) {
	buf := []byte("__________")
	for i := 0; i < b.N; i++ {
		Decode(buf)
	}
}
