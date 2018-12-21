package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := []int{}
	return func() int {
		var f int
		if len(fib) == 0 {
			f = 0
		} else if len(fib) == 1 {
			f = 1
		} else {
			f = fib[len(fib)-2] + fib[len(fib)-1]
		}
		fib = append(fib, f)
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
