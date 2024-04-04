package main

import (
	"testing"
)

var (
	arrValInt   = []int{10, 20, 15, 11, 9}
	arrValFloat = []float64{10.23, 9.66, 13.56, 12.56}
)

func Benchmark_Bubble_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(arrValInt)
	}
}

func Benchmark_Bubble_Float_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleGeneric(arrValFloat)
	}
}
