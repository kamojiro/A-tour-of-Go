package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

// Scale ...
func (v *Vertex) Scale(f float64)  {
	v.X = v.X*f
	v.Y = v.Y * f
}

// Abs 
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3,4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs()) //+vは構造体
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
