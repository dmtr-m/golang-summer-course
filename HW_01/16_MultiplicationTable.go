package main

import "fmt"

func MultiplicationTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n * i)
	}
}

func TestMultiplicationTable() {
	var n int
	fmt.Scan(&n)
	MultiplicationTable(n)
}
