package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotFound = errors.New("not found")
)

func main() {
	arr := []int{20, 30, 50, 60, 70, 80, 90, 100}
	target := 90

	res, err := Binary(arr, target)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("res\t", res)
}

func Binary(arr []int, target int) (int, error) {
	start := 0
	end := len(arr) - 1
	var mid int

	for start <= end {
		mid = start + (end-start)/2
		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			return mid, nil
		}
	}

	return -1, ErrNotFound
}
