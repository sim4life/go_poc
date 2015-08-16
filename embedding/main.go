package main

import (
	"fmt"
)

type vertex struct {
	x, y float64
}

type vertex3D struct {
	vertex
	z float64
}

func main() {
	var vert3D vertex3D
	//these both work
	vert3D.vertex.x = 2.0
	vert3D.y = 4.0

	fmt.Println("Vertex3D is: ", vert3D)
}
