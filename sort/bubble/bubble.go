package main

import (
	"fmt"
)

func main() {
	arr := []int{60, 50, 100, 90, 20, 30}
	res := Bubble(arr)
	fmt.Println("sort bubble:\t", res)

}

func Bubble(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
