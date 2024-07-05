package main

import (
	"fmt"
	"slices"
)

func Max(array []int) int {
	return slices.Max(array)
}

func TestMaxOfThree() {
	numbers := make([]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println(Max(numbers))
}
