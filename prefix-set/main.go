package main

import (
	"fmt"
)

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) <= 1 {
		return 0
	} else {
		unique := []int{A[0]}
		prefix := 0
		isFound := false
		for ind, val := range A {
			isFound = false
			for i := 0; i < ind; i++ {
				if val == A[i] {
					isFound = true
				}
			}
			/*
				for _, uni := range unique {
					if val == uni {
						isFound = true
					}
				}
			*/
			if !isFound {
				prefix = ind
				unique = append(unique, val)
			}
		}
		return prefix
	}
}

func main() {
	// fmt.Println("The Solution is: ", Solution([]int{1, 4, 6, 7, 4}))
	fmt.Println("The Solution is: ", Solution([]int{2, 2, 1, 0, 1}))
}
