package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount ...
func WordCount(s string) map[string]int {
	// var wordsplitting []string = strings.Fields(s)
	wordsplitting := strings.Fields(s)
	var m map[string]int
	m = make(map[string]int)
	var ok bool
	for _, value := range wordsplitting {
		_, ok = m[value]
		if ok {
			m[value] += 1
		}else{
			m[value] = 1
		}
	}
	// return map[string][int]{"x":1}
	return m
}

func main() {
	wc.Test(WordCount)
}
