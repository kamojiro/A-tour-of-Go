package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// (v *Vertex) ...
func (v *Vertex) Scale(f float64)  {
	// * を外すと、main の結果は f 倍されない
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())

}
