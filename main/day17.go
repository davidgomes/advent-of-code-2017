package main

import "fmt"

func main() {
	fmt.Println("Hello day 17")

	buffer := []int{0}
	pos := 0
	inp := 382

	for i := 1; i <= 50000000; i++ {
		for u := 0; u < inp; u++ {
			if pos == len(buffer) - 1 {
				pos = 0
			} else {
				pos++
			}
		}

		split := make([]int, len(buffer) - (pos + 1))
		copy(split, buffer[(pos + 1):])

		for u := 0; u < len(split) - 1; u++ {
			buffer[pos + u + 2] = split[u]
		}

		if len(split) > 0 {
			buffer = append(buffer, split[len(split)-1])
		} else {
			buffer = append(buffer, 0)
		}

		buffer[pos + 1] = i

		if pos == len(buffer) - 1 {
			pos = 0
		} else {
			pos++
		}
	}

	fmt.Println(buffer) // solution 1561
}
