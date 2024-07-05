package main

import "fmt"

func IsPalindrome(s string) bool {
	s_runes := []rune(s)
	for i, j := 0, len(s_runes)-1; i < j; i, j = i+1, j-1 {
		if s_runes[i] != s_runes[j] {
			return false
		}
	}
	return true
}

func TestPalindrome() {
	var input string
	fmt.Scan(&input)
	fmt.Println(IsPalindrome(input))
}
