package main

import (
	"testing"
)

var (
	arr = []int{10, 20, 15, 11, 9}
)

func Benchmark_Bubble_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bubble(arr)
	}
}
