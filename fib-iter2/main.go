package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	if arg == "" {
		fmt.Println("Please provide a number argument!")
		return
	}

	fmt.Println("arg is " + arg)
	// num := 5
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Please provide a number argument!")
		return
	}

	j := 0
	for i, k := 1, 0; k < num; i, j, k = i+j, i, k+1 {
		//fmt.Println(i)
	}

	fmt.Printf("fib(%d) is: %d\n", num, j)

}
