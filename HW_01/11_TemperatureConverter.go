package main

import "fmt"

func Celsius2Fahrenheit(celsius float32) float32 {
	return (celsius * 1.8) + 32
}

func TestTemperatureConverter() {
	var celsius float32
	fmt.Scan(&celsius)
	fmt.Println(Celsius2Fahrenheit(celsius))
}
