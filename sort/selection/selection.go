package main

import (
	"fmt"
	"github.com/mrbardia72/algorithm/constraints"
)

func main() {
	var items = []int{2, 1, 1, 2, 2, 3, 5, 8, 6}
	res := Selection(items)
	fmt.Println("sort selection:\t", res)

}

func Selection[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		minx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minx] {
				minx = j
			}
		}
		arr[i], arr[minx] = arr[minx], arr[i]
	}

	return arr
}
