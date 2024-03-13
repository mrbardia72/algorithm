package main

import "testing"

func Fuzz_Bubble_Test(f *testing.F) {
	arrx := []int{10, 25, 2, 6, 100, 8}
	for _, tc := range arrx {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, s int) {
		Bubble(arrx)
	})
}
