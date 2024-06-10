package main

import (
	"fmt"
)

func reverse[T int | string](slice []T) []T {
	newSlice := make([]T, len(slice))
	reverseIndex := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newSlice[reverseIndex] = slice[i]
		reverseIndex--
	}
	return newSlice
}

func main() {

	fmt.Println("Array de Ints ->", reverse([]int{1, 2, 3, 4, 5}))
	fmt.Println("Array de string ->", reverse([]string{"a", "b", "c", "d", "e"}))
}
