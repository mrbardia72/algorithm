package main

import "testing"

var (
	items = []int{10, 20, 15, 11, 9}
	find  = 11
)

func Benchmark_Liner_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Binary(items, find)
	}
}
