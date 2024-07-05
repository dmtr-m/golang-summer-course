package main

import "fmt"

func PrintPrimesUpTo(size int) {
	size++
	sieve := make([]bool, size)
	var primes []int

	for number := 2; number < size; number++ {
		if sieve[number] == false {
			primes = append(primes, number)

			for step := 2 * number; step < size; step = step + number {
				sieve[step] = true
			}
		}
	}

	for _, prime := range primes {
		fmt.Println(prime)
	}
}
