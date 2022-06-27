package main

import "testing"

func BenchmarkFibFn(b *testing.B) {
	fn := FibFn()
	for i := 0; i < b.N; i++ {
		fn()
	}
}
