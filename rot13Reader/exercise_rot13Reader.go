package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13r *rot13Reader) Read(rbyte []byte) (int, error) {
	n, err := rot13r.r.Read(rbyte)

	for i, v := range rbyte {
		if (v > 'A' && v < 'N') || (v > 'a' && v < 'n') {
			rbyte[i] += 13
		} else if (v > 'M' && v < 'Z') || (v > 'm' && v < 'z') {
			rbyte[i] -= 13
		}
	}
	return n, err
}

func main() {
	// func NewReader(s string) *Reader, in package "strings"
	// func (r *Reader) Read(b []byte) (n int, err error)

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	n, err := io.Copy(os.Stdout, &r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%v\n", n)
	}
}
