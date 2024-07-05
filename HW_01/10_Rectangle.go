package main

import "fmt"

type Rectangle struct {
	height, width int
}

func Area(rectangle Rectangle) int {
	return rectangle.height * rectangle.width
}

func TestRectangle() {
	var rectangle Rectangle
	fmt.Scan(&rectangle.height, &rectangle.width)
	fmt.Println(Area(rectangle))
}
