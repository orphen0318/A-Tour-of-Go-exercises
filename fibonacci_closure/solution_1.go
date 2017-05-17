package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var hisfi []int
	i := 0
	fi := 0
	return func() int {
		if i < 2 {
			fi = i
		} else {
			fi = hisfi[i-2] + hisfi[i-1]
		}
		i++
		hisfi = append(hisfi, fi)

		return fi
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibo[%d]: %d\n", (i + 1), f())
	}
}
