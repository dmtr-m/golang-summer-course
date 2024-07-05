package main

import "fmt"

func Delete(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func TestDeleteFromSlice() {
	var array = []int{1, 2, 3, 4, 5, 6}
	var index = int(2)
	fmt.Println(Delete(array, index))
}
