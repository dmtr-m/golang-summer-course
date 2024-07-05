package main

import "fmt"

func Fibonacci(n int) []int {
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 1}
	}

	var fibonacci_sequence = []int{1, 1}
	for i := 2; i < n; i++ {
		fibonacci_sequence = append(fibonacci_sequence, int(fibonacci_sequence[i-1]+fibonacci_sequence[i-2]))
	}

	return fibonacci_sequence
}

func TestFibbonacci() {
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(10))
}
