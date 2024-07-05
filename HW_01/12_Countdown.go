package main

import "fmt"

func Countdown(time int) {
	for i := time; i >= 0; i-- {
		fmt.Println(i)
	}
}

func TestCountdown() {
	var time int
	fmt.Scan(&time)
	Countdown(time)
}
