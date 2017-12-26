package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day4.in") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	lines := strings.Split(str, "\n")
	count := 0

	for _, l := range lines {
		words := strings.Split(l, " ")

		//var m map[string]int
		m := make(map[string]int)

		valid := true

		for _, w := range words {
			_, ok := m[w]

			if !ok {
				m[w] = 1
			} else {
				m[w]++
				valid = false
			}
		}

		if valid {
			//fmt.Println(l)
			count++
		} else {
			fmt.Println(l)
		}
	}

	fmt.Println(count)
}
