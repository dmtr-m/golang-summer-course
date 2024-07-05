package main

import "fmt"

func Sum(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}

	return sum
}

func TestSumOfArray() {
	var number_elements int
	fmt.Scan(&number_elements)

	array := make([]int, number_elements)
	for i := 0; i < number_elements; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println(Sum(array))
}
