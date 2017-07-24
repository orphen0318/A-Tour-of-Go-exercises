package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func WalkControl(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	c1 := make(chan int)
	c2 := make(chan int)

	go WalkControl(t1, c1)
	go WalkControl(t2, c2)

	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}
	return true
}

func main() {

	// Test the Walk function
	ch := make(chan int)
	go WalkControl(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}

	// Test the Same function
	fmt.Println("two tree is identify: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("two tree is identify: ", Same(tree.New(1), tree.New(2)))
}
