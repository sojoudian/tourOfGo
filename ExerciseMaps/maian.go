package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		counts[v]++
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
