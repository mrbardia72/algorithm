package main

import (
	"fmt"
	"github.com/mrbardia72/algorithm/constraints"
)

func main() {
	arr := []int{60, 50, 100, 90, 20, 30}
	res := Bubble(arr)
	fmt.Println("sort bubble:\t", res)

	arrGenric := []float64{60.2, 50.32, 100.56, 90.19, 20.55, 30.69}
	resGenric := BubbleGeneric(arrGenric)
	fmt.Println("sort bubble genric:\t", resGenric)

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

func BubbleGeneric[T constraints.Ordered](arr []T) []T {
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
