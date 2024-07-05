package main

import "fmt"

func Contains(array []int, value int) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

func TestContains() {
	var array = []int{1, 2, 3, 4, 6}
	var value = int(1)
	fmt.Println(Contains(array, value))
}
