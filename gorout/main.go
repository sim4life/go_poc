package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	time.Sleep(100 * time.Millisecond)
	say("hello")
}

/*
I don't understand, without main() Sleep(), the output of
hello
world
hello
world
hello
*/
