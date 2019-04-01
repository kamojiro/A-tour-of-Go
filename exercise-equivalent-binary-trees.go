package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

//trees.go

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	var walk func(*tree.Tree, chan int)
	walk = func(t *tree.Tree, ch chan int){
		l := t.Left
		r := t.Right
		if l.String() != "()" {
			walk(l, ch)
		}
		ch <- t.Value
		if r.String() != "()" {
			walk(r, ch)
		}
	}
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var ans bool = true
	Walk(t1, ch1)
	Walk(t2, ch2)
	for i := range ch1 {
		if i != <- ch2 {
			ans = false
		}
	}
	return ans
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println( Same( tree.New(1), tree.New(2)))
	fmt.Println( Same( tree.New(1), tree.New(1)))

}
