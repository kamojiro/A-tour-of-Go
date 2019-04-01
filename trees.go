package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func main() {
	t := tree.New(1)
	
	fmt.Println(t.Left, t.Value, t.Right)
	r := t.Right
	fmt.Printf("%T\n", r)
	fmt.Println(r.String() == "()")
	fmt.Println(r == nil)

}
