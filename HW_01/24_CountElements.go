package main

import "fmt"

func Count(array []int, target_element int) int {
	values_counter := make(map[int]int)
	for _, value := range array {
		if _, contains := values_counter[value]; !contains {
			values_counter[value] = 1
		} else {
			values_counter[value] += 1
		}
	}

	return values_counter[target_element]
}

func TestCountElements() {
	var array = []int{0, 1, 1, 1, 2, 3, 4, 5, 6}
	fmt.Println(Count(array, 1))
	fmt.Println(Count(array, 7))
}
