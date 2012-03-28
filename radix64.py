# Copyright 2012 Kevin Gillette. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

_ord = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
_table = dict((c, i) for i, c in enumerate(_ord))

def encode(n, len=0):
	out = ""
	while n > 0 or len > 0:
		out = _ord[n & 63] + out
		n >>= 6
		len -= 1
	return out

def decode(input):
	n = 0
	for c in input:
		c = _table.get(c)
		if c is None:
			raise ValueError("Invalid character in input: " + c)
		n = n << 6 | c
	return n
