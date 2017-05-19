package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // ErrNegativeSqrt(x) is constructor.
	}
	for {
		z := 1.0
		if math.Abs(z-(z-(z*z-x)/(2*z))) < 1e-10 {
			return z, nil
		} else {
			z = z - (z*z-x)/(2*z)
		}

		// only loop 10 times to find Sqrt
		// z := 1.0
		// for i := 0; i < 10; i++ {
		// 	z = z - (z*z-x)/(2*z)
		// }
		// return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
