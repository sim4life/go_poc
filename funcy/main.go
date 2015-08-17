package main

import (
	"fmt"
)

func vardiacFun(fun string, params ...int) {
	fmt.Printf(" \"%s\" and Parameter type: %T\n", fun, params)
	fmt.Println(params)
	for _, elem := range params { //loop of the parameter
		fmt.Println(elem)
	}
}

func main() {
	vardiacFun("some", 1, 3, 4, 5, 6)
}
