package main

import (
	. "fmt"
	. "math"
	"os"
	"strconv"
)

func nSqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z
}

func both(i float64) (float64, float64) {
	return nSqrt(i), Sqrt(i)
}

func main() {
	var num float64

	if len(os.Args[1:]) > 0 {
		num, _ = strconv.ParseFloat(os.Args[1], 64)
	} else {
		num = 5
	}

	Println(both(num))
}
