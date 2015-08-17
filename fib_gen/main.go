package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib_generator() chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = j, i+j {
			c <- i
		}
	}()

	return c
}

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

	c := fib_generator()
	for n := 0; n <= num/2; n++ {
		fmt.Print(" ", <-c)
	}
	fmt.Println()

	for n := 0; n < num/2; n++ {
		fmt.Print(" ", <-c)
	}
	if num%2 > 0 {
		fmt.Print(" ", <-c)
	}
	fmt.Println()

}
