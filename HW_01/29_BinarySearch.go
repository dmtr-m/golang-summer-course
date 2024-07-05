package main

import "fmt"

func BinarySearch(array []int, search_value int) int {
	var left_bound = 0
	var right_bound = len(array)

	for left_bound+1 < right_bound {
		mid := (left_bound + right_bound) / 2
		if array[mid] < search_value {
			left_bound = mid
		} else {
			right_bound = mid
		}
	}

	if array[right_bound] != search_value {
		return -1
	}

	return right_bound
}

func TestBinarySearch() {
	var array = []int{1, 2, 3, 5, 5, 6, 7, 8}
	var value = int(5)
	fmt.Println(BinarySearch(array, value))
}
