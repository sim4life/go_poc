package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			//it is non-blocking but if any channel is accessed here
			//then this section also blocks.
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
