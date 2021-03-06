package main

import (
	"fmt"
)

type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *IntSet) Get(i int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *IntSet) Remove(i int) {
	delete(set.set, i)
}

func (set *IntSet) isEmpty() bool {
	for _, val := range set.set {
		if val == true {
			return false
		}
	}
	return true
}

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) <= 1 {
		return 0
	} else {
		set := NewIntSet()
		for _, val := range A {
			set.Add(val)
		}
		for ind, value := range A {
			set.Remove(value)
			if set.isEmpty() {
				return ind
			}
		}
		return 0

	}
}

func someFunc(set *IntSet) {
	fmt.Printf("set[1] is: %v\n", set.Get(1))
	fmt.Printf("set[2] is: %v\n", set.Get(2))
	set.Add(2)
	fmt.Printf("set[1] is: %v\n", set.Get(1))
	fmt.Printf("set[2] is: %v\n", set.Get(2))

}
func main() {
	//fmt.Println("The Solution is: ", Solution([]int{1, 4, 6, 7, 4}))
	//fmt.Println("The Solution is: ", Solution([]int{2, 2, 1, 0, 1}))

	some := NewIntSet()

	some.Add(1)
	fmt.Printf("some[1] is: %v\n", some.Get(1))
	fmt.Printf("some[2] is: %v\n", some.Get(2))
	someFunc(some)
	fmt.Printf("some[1] is: %v\n", some.Get(1))
	fmt.Printf("some[2] is: %v\n", some.Get(2))

}
