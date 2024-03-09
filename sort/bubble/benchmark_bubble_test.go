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

// go test -bench=. -v
// go test -bench=. -count 5
// Display memory allocation statistics > go test -bench=. -benchmem

// sample output:
//Benchmark_Bubble_Test
//Benchmark_Bubble_Test-4         214033183                5.425 ns/op
//PASS
//ok      github.com/mrbardia72/algorithm/sort/bubble     1.732s

// The -4 suffix denotes the number of CPUs used to run the benchmark
// On the right side of the function name, you have two values, 214033183 and 5.425 ns/op. The former indicates the
//total number of times the loop was executed, while the latter is the average amount of time each iteration took to
//complete, expressed in nanoseconds per operation.
