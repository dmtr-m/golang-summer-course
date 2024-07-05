package main

import "fmt"

func Merge(array_1, array_2 []int) []int {
	var result []int

	var pointer_1 int = 0
	var pointer_2 int = 0
	for pointer_1 < len(array_1) && pointer_2 < len(array_2) {
		if array_1[pointer_1] <= array_2[pointer_2] {
			result = append(result, array_1[pointer_1])
			pointer_1++
		} else {
			result = append(result, array_2[pointer_2])
			pointer_2++
		}
	}

	if pointer_1 < len(array_1) {
		result = append(result, array_1[pointer_1:]...)
	}
	if pointer_2 < len(array_2) {
		result = append(result, array_2[pointer_2:]...)
	}

	return result
}

func TestMerge() {
	var array_1 = []int{1, 1}
	var array_2 = []int{10, 10}
	fmt.Println(Merge(array_1, array_2))
}
