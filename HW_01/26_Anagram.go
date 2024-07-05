package main

import (
	"fmt"
	"strings"
)

func IsAnagram(string_1, string_2 string) bool {
	string_1 = strings.ToLower(string_1)
	string_2 = strings.ToLower(string_2)

	var letters_string_1 = make(map[rune]int)
	for _, letter := range string_1 {
		letters_string_1[letter] += 1
	}

	for _, letter := range string_2 {
		if letters_string_1[letter] > 0 {
			letters_string_1[letter] -= 1
		} else {
			return false
		}
	}

	for _, value := range letters_string_1 {
		if value > 0 {
			return false
		}
	}

	return true
}

func TestIsAnagram() {
	var string_1 = "aaaa"
	var string_2 = "aaaa"
	fmt.Println(IsAnagram(string_1, string_2))
}
