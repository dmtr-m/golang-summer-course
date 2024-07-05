package main

import "fmt"

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func TestBubbleSort() {
	var array = []int{1, 2, 6, 5, 3, 2, 0, -1, 2, 9}
	BubbleSort(array)
	fmt.Println(array)
}
