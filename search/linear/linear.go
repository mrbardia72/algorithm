package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 5, 1, 6, 9}
	target := 1

	res := Linear(arr, target)
	fmt.Println("res\t", res)

}

func Linear(arr []int, target int) int {
	for i, item := range arr {
		if item == target {
			return i
		}
	}
	return -1

}
