package main

import "fmt"

func LinearSearch(array []int, target_value int) int {
	for index, value := range array {
		if value == target_value {
			return index
		}
	}
	return -1
}

func TestLinearSearch() {
	var array = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(LinearSearch(array, 0))
}
