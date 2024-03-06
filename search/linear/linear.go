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
	arr := []int{2, 3, 5, 1, 6, 9}
	target := 1

	res, err := Linear(arr, target)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("res\t", res)

}

func Linear(arr []int, target int) (int, error) {
	for i, item := range arr {
		if item == target {
			return i, nil
		}
	}
	return -1, ErrNotFound

}
