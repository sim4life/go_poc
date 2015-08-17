package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci(num int) int {
	if num == 1 {
		return 1
	} else if num == 0 {
		return 0
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
	f := fibonacci(9)
	fmt.Println(f)

}
