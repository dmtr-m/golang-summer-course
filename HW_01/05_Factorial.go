package main

import "fmt"

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func TestFactorial() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Factorial(n))
}
