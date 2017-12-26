package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day2.in") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	lines := strings.Split(str, "\n")

	var checksum int

	for _, l := range (lines) {
		numbers := strings.Split(l, ",")

		smallest, largest := 99999999999999, 0

		for _, strN := range (numbers) {
			n, err := strconv.Atoi(strN)

			if err != nil {
				fmt.Println(err)
				return
			}

			smallest = min(smallest, n)
			largest = max(largest, n)
		}

		fmt.Println(largest, smallest)
		fmt.Println(largest - smallest)
		checksum += largest - smallest
	}

	fmt.Println(checksum)

	//fmt.Println(str) // print the content as a 'string'
}
