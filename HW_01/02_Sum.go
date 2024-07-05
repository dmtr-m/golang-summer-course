package main

import (
	"fmt"
)

func Sum2vals(x, y int) int {
	return x + y
}

func TestSum() {
	var number_1, number_2 int
	fmt.Scan(&number_1, &number_2)
	fmt.Println(Sum2vals(number_1, number_2))
}
