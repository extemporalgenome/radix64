// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package radix64

import (
	"errors"
)

const ord = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

var table [256]byte

const (
	// starting offsets
	upper = '9' - '0' + 1
	lower = upper + 'Z' - 'A' + 1
)

const costMask = ^uint64(63)

var (
	ErrBufLimReached = errors.New("Buffer limit reached")
	ErrInvalidByte   = errors.New("Input cannot be decoded")
)

func Encode(n uint64, b []byte) error {
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = ord[n&63]
		n >>= 6
	}
	if n > 0 {
		return ErrBufLimReached
	}
	return nil
}

func Decode(b []byte) (n uint64, err error) {
	for _, c := range b {
		c = table[c]
		if c&64 == 0 {
			return 0, ErrInvalidByte
		} else {
			n = n<<6 | uint64(c&63)
		}
	}
	return
}

func Cost(n uint64) (bytes, remainder int) {
	bytes = 1
	for n&costMask != 0 {
		n >>= 6
		bytes++
	}
	remainder = 6
	for n != 0 {
		n >>= 1
		remainder--
	}
	return
}

func init() {
	for i := 0; i < len(ord); i++ {
		table[ord[i]] = byte(i | 64)
	}
}
