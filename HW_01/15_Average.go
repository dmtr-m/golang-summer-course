package main

import "fmt"

func Average(array []int) float32 {
	sum := int(0)
	for _, val := range array {
		sum += val
	}

	return float32(sum) / float32(len(array))
}

func TestAverage() {
	var array = []int{1, 2, 3, 4, 6}
	fmt.Println(Average(array))
}
