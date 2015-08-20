package main

import (
	"fmt"
)

type circle struct {
	// length int
	coorL int
	coorR int
}

func Solution(A []int) int {
	// write your code in Go 1.4
	intersects := 0
	var circles = make([]circle, len(A))
	for ind, val := range A {
		circles[ind] = circle{ind - val, ind + val}
	}

	for ind, val := range circles {
		for i := 0; i < ind; i++ {
			if val.coorL <= circles[i].coorR && val.coorR >= circles[i].coorL {
				intersects++
			}
		}
		if intersects > 10000000 {
			return -1
		}
	}
	return intersects
}

func main() {
	fmt.Println("The Solution is: ", Solution([]int{1, 5, 2, 1, 4, 0}))
}
