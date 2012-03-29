package main

import (
	"flag"
	"fmt"
	"github.com/extemporalgenome/radix64"
	"os"
	"strconv"
)

func main() {
	var (
		encode, decode, stats bool
		width                 int
		n                     uint64
		in                    string
		out                   []byte
		err                   error
	)
	flag.BoolVar(&encode, "e", false, "encode")
	flag.IntVar(&width, "w", 0, "padding width (optional, for fixed-size encoding)")
	flag.BoolVar(&decode, "d", false, "decode")
	flag.BoolVar(&stats, "s", false, "stats")
	flag.Parse()
	if in = flag.Arg(0); in == "" {
		fmt.Fprintln(os.Stderr, "Expected an argument")
		os.Exit(2)
	}
	if width < 0 {
		width = 0
	}

	switch {
	default:
		// if neither encode nor decode are specified, try both
		fallthrough
	case encode:
		n, err = strconv.ParseUint(in, 10, 64)
		if err == nil {
			if stats {
				goto stats
			}
			if width == 0 {
				width, _ = radix64.Cost(n)
			}
			out = make([]byte, width+1)
			if err = radix64.Encode(n, out[:width]); err != nil {
				goto error
			}
			out[width] = '\n'
			os.Stdout.Write(out)
			return
		} else if encode != decode {
			goto error
		}
		// encode == decode (try both)
		fallthrough
	case decode:
		n, err = radix64.Decode([]byte(in))
		if err != nil {
			goto error
		} else if stats {
			goto stats
		}
		fmt.Println(n)
		return
	}
	return
error:
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
stats:
	bytes, remainder := radix64.Cost(n)
	fmt.Println(n, "requires", bytes, "bytes, of which", remainder, "bits are padding")
}
