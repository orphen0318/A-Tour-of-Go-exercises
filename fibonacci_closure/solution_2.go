package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fir, sec := 1, 0
	return func() int {
		fir, sec = sec, fir+sec
		return fir
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibo[%d]: %d\n", (i + 1), f())
	}
}
