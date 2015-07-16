package main

import (
	"fmt"
	"github.com/TheProfs/stringutil"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 = map[string]Vertex{
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Printf(stringutil.Reverse("\nhello, smart world!n\n"))
	//fmt.Printf("hello, smart world!")

	// using slices - s[lo:hi] is lo through hi-1
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	//a := make([]int, 5)    // a len=5 cap=5 [0 0 0 0 0]
	//b := make([]int, 0, 5) // b len=0 cap=5 []

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for _, value := range pow { //skip index
	//for index, _ := range pow { //skip value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	m4 := make(map[string]int)

	m4["Answer"] = 42
	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // 0, false

}
