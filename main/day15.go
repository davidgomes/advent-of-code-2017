package main

import (
	"fmt"
)

const divisonBy = 2147483647

func main() {
	fmt.Println("Hello Day 15")

	previousA, factorA := 591, 16807
	previousB, factorB := 393, 48271

	matches := 0

	for i := 0; i < 40000000; i++ {
		newValueA := (previousA * factorA) % divisonBy
		newValueB := (previousB * factorB) % divisonBy

		binA := fmt.Sprintf("%032b", int64(newValueA))
		binB := fmt.Sprintf("%032b", int64(newValueB))

		if binA[16:] == binB[16:] {
			matches++
		}

		previousA, previousB = newValueA, newValueB
	}

	fmt.Println(matches)
}
