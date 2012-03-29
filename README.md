# radix64

radix64 is similar to base64 encoding, but has a different intent:
base64 is intended to encode arbitrary, variable-length binary data,
while radix64 is designed to semantically encode unsigned integers
(not their binary representation) into fixed-width, URL-safe strings.

The encoded character set is a rearrangement of the base64url
alphabet as defined in RFC4648, though radix64 uses 0-9 as the lowest
valued ordinals, making zero-padding recognizable.

radix64 is a good solution for transmitting large arrays of
fixed-size integers over a protocol that isn't entirely binary-safe
(application/x-www-form-urlencoded). In cases where a JSON array may
have been used to encode integers over a fixed range, radix64 can be
considerably more compact, especially considering that may characters
in JSON will get urlencoded.

radix64 currently has native Go, Python, and Javascript
implementations.

# Limitations

The Go and JavaScript implementations currently do not check for
integer overflow when decoding.

The Go implementation decodes into uint54 values -- 10 input bytes
will never overflow.

The JavaScript implementation may be engine dependent. V8 (through
Chrome) dealt with int literals as signed 32-bit integers, and so it
may only be able to cleanly decode 5 bytes.

# Examples

## Go

	import "github.com/extemporalgenome/radix64"

	b := make([]byte, 6)
	err := radix64.Encode(45948488810, b)
	// string(b) == "golang"
	n, err := radix64.Decode([]byte("go"))
	// n == 2738

## Python

	>>> import radix64
	>>> radix64.encode(27864775857)
	'Python'
	>>> radix64.decode('py')
	3324

## JavaScript

	enc64(1244) == 'JS';
	dec64('-JS-') == 16332606;
