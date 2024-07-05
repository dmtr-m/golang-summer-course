package main

import "fmt"

func RemoveDuplicates(array []int) []int {
	var result []int

	contains_list := make(map[int]bool)
	for _, value := range array {
		if _, contains := contains_list[value]; !contains {
			result = append(result, value)
			contains_list[value] = true
		}
	}

	return result
}

func TestRemoveDuplicates() {
	var array = []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7}
	fmt.Println(RemoveDuplicates(array))
}
