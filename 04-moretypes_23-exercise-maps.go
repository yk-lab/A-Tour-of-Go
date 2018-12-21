package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	l := strings.Fields(s)
	wc := make(map[string]int)
	for _, i := range l {
		c, ok := wc[i]
		if ok {
			wc[i] = c + 1
		} else {
			wc[i] = 1
		}
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
