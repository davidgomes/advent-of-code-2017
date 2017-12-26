package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

// bmfehlanpkjcidog

func main() {
	fmt.Println("Hello Day 16")

	m := make(map[string]int)

	length := 16
	for i := 0; i < length; i++ {
		m[string(97 + i)] = i
	}

	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day16.in")

	str := string(b)
	cmds := strings.Split(str, ",")

	arr := make([]string, length)

	for key, val := range m {
		arr[val] = key
	}

	fmt.Println(arr)

	for _, cmd := range cmds {
		for key, val := range m {
			arr[val] = key
		}

		if string(cmd[0]) == "s" {
			var x int
			fmt.Sscanf(cmd, "s%d", &x)

			spinned := make([]string, x)

			copy(spinned, arr[len(arr) - x:])

			notSpinned := make([]string, len(arr) - x)

			for i := 0; i < len(arr) - x; i++ {
				notSpinned[i] = arr[i]
			}

			u := 0
			for i := x; i < x + len(notSpinned); i++ {
				arr[i] = notSpinned[u]
				u++
			}

			for i := 0; i < len(spinned); i++ {
				arr[i] = spinned[i]
			}

			for i := 0; i < len(arr); i++ {
				m[arr[i]] = i
			}
		} else if string(cmd[0]) == "x" {
			var a, b int
			fmt.Sscanf(cmd, "x%d/%d", &a, &b)

			fmt.Println(a, b)

			m[arr[a]], m[arr[b]] = m[arr[b]], m[arr[a]]
		} else if string(cmd[0]) == "p" {
			var s string
			fmt.Sscanf(cmd, "p%s", &s)

			splitS := strings.Split(s, "/")

			a := splitS[0]
			b := splitS[1]

			m[a], m[b] = m[b], m[a]
		}

		for key, val := range m {
			arr[val] = key
		}

		fmt.Println(arr)
	}

	for key, val := range m {
		arr[val] = key
	}

	fmt.Println(arr)
}