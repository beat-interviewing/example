package main

import "testing"

func TestFibFn(t *testing.T) {

	fn := FibFn()

	sequence := []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	
	for i, expected := range sequence {

		t.Run(i, func(t *testing.T) {
			actual := fn()
			if (expected != actual) {
				t.Errorf("Wrong value %d in sequence. Expected %d", actual, expected)
			}
		})
	}
}