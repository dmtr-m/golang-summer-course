package main

import (
	"fmt"
)

func Reverse(input string) string {
	input_runes := []rune(input)
	for i, j := 0, len(input_runes)-1; i < j; i, j = i+1, j-1 {
		input_runes[i], input_runes[j] = input_runes[j], input_runes[i]
	}
	return string(input_runes)
}

func TestReverseString() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(Reverse(input))
}
