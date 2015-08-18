package main

// import "code.google.com/p/go-tour/tree"
import (
	"fmt"
	"golang.org/x/tour/tree"
	// "github.com/golang/tour/tree"
)

/**********
INCORRECT IMPLEMENTATION
***********/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)
	// closing the channel so range can finish
	close(ch)
}

// recWalk walks recursively through the tree and push values to the channel
// at each recursion
func recWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		// send the left part of the tree to be iterated over first
		recWalk(t.Left, ch)
		// push the value to the channel
		ch <- t.Value
		// send the right part of the tree to be iterated over last
		recWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			// not the same size
			fmt.Println("not the same size")
			return false
		case !ok1: //ok2 is also false
			// both channels are empty
			fmt.Println("both channels are empty")
			return true
		case x1 != x2:
			// elements are different
			fmt.Printf("elements are different: x1=%d and x2=%d\n", x1, x2)
			return false
		default:
			// keep iterating
		}
	}
}

func main() {

	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

	ta23 := &tree.Tree{nil, 1, nil}
	ta2 := &tree.Tree{nil, 2, nil}
	ta3 := &tree.Tree{nil, 3, nil}
	ta23.Left = ta2
	ta23.Right = ta3

	tb23 := &tree.Tree{&tree.Tree{nil, 2, nil}, 1, &tree.Tree{nil, 3, nil}}

	fmt.Println("ta23: ", ta23)
	fmt.Println("tb23: ", tb23)

	fmt.Println("ta23 and tb23 same: ", Same(ta23, tb23)) //this test passes

}
