package main

import (
	"fmt"
	"math"
)

type MyInt int

func (mi MyInt) Abs() int {
	if mi < 0 {
		return int(-1 * mi)
	} else {
		return int(mi)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var mi MyInt = -10
	fmt.Println(mi.Abs())

	v1 := Vertex{3, 4}
	// all works
	fmt.Println(v1.Abs())
	fmt.Println((&v1).Abs())
	fmt.Println((*(&v1)).Abs())
}
