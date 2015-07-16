package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	broken := 0
	//var new_z float64;
	old_z, new_z := x, x
	if new_z = x / 2; x > 4 {
	}

	for i := 1; i <= 10; i++ {
		new_z = new_z - (math.Pow(new_z, 2)-x)/(2*new_z)
		//fmt.Println(old_z)
		//fmt.Println(new_z)
		if old_z-new_z < 0.0000000001 {
			broken = i
			break
		}
		old_z = new_z
	}
	if broken > 0 {
		fmt.Println("it was broken: ", broken)
	}
	return new_z
}

func main() {
	fmt.Println(Sqrt(55))
}
