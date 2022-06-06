package main

import "fmt"

func FibFn() func() uint64 {

	return func() uint64 {
		return 0
	}
}

func main() {
	fib := FibFn()
	for j := 0; j < 10; j++ {
		fmt.Printf("%v ", fib())
	}
}
