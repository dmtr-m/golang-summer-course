package main

import (
	"fmt"
	"strings"
)

func IsVowel(character rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, character)
}

func TestIsVowel() {
	var input string
	fmt.Scanln(&input)

	character := rune(input[0])
	if IsVowel(character) {
		fmt.Println("The character is a vowel")
	} else {
		fmt.Println("The character is a consonant")
	}
}
