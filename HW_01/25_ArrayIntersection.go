package main

import "fmt"

func Intersection(array_1, array_2 []int) []int {
	values_counter := make(map[int]int)
	for _, value := range array_1 {
		if _, contains := values_counter[value]; !contains {
			values_counter[value] = 1
		} else {
			values_counter[value] += 1
		}
	}

	var result []int

	for _, value := range array_2 {
		if values_counter[value] > 0 {
			result = append(result, value)
			values_counter[value] -= 1
		}
	}

	return result
}

func TestArrayIntersection() {
	var array_1 = []int{1, 2, 3, 4, 1, 2, 3, 1, 0}
	var array_2 = []int{1, 2, 3, 2, 2, 1, 2, 3, 3}
	fmt.Println(Intersection(array_1, array_2))
}
