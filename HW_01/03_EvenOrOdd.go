package main

import (
	"fmt"
)

func IsEven(x int) bool {
	return x%2 == 0
}

func TestEvenOrOdd() {
	var number_1 int
	fmt.Scan(&number_1)
	fmt.Println(IsEven(number_1))
}
