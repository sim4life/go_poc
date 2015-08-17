package main

import (
	"fmt"
)

func Mycounter() func() {
	theCount := 0
	increment := func() {
		theCount++
		fmt.Println("The count is", theCount)
	}
	return increment
}

func Mycounter2(initialCount int) (func(), func() int) {
	theCount := initialCount
	increment := func() {
		theCount++
		fmt.Println("The inc count is", theCount)
	}
	get := func() int {
		fmt.Println("The get count is", theCount)
		return theCount
	}
	return increment, get
}

func main() {
	// incFunc := Mycounter()
	incFunc, getFunc := Mycounter2(10)
	for i := 0; i < 5; i++ {
		incFunc() //use () to execute increment
	}
	fmt.Println("The final value is ", getFunc())

}
