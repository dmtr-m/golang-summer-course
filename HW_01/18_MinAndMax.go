package main

import (
	"fmt"
	"slices"
)

func MinAndMax(array []int) (int, int) {
	return slices.Min(array), slices.Max(array)
}

func TestMinAndMax() {
	var array = []int{1, 2, 3, 4, 1, 2, 3, -1, 10}
	fmt.Println(MinAndMax(array))
}
