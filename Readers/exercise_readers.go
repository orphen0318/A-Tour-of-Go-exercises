package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(rbyte []byte) (int, error) {
	for i := range rbyte {
		rbyte[i] = 'A'
	}

	return len(rbyte), nil
}

func main() {
	reader.Validate(MyReader{})
}
