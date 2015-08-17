package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(num int) (fib_num int) {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}

	fib_num1, fib_num := 0, 1

	for i := 2; i <= num; i++ {
		fib_num1, fib_num = fib_num, fib_num1+fib_num
	}

	return
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

	fib_num := fib(num)
	fmt.Println("fib(5) is: ", fib_num)

}
