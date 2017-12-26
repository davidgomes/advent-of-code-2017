package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

var visited map[int]bool
var m map[int]map[int]bool
var cache map[int]int
var res int

func dfs(current int) {
	if !visited[current] {
		//fmt.Printf("increasing because I found %d\n", current)
		res += 1
	} else {
		return
	}

	visited[current] = true

	for child, _ := range m[current] {
		if current != child {
			dfs(child)
		}
	}
}

func main() {
	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day12.in")

	str := string(b)
	lines := strings.Split(str, "\n")

	m = make(map[int]map[int]bool)

	for i:=0; i<=len(lines); i++ {
		m[i] = make(map[int]bool)
	}

	for _, line := range lines {
		splitLine := strings.Split(line, "<->") // []string

		n, _ := strconv.Atoi(strings.Trim(splitLine[0], " "))

		for _, strChild := range strings.Split(splitLine[1], ",") {
			child, _ := strconv.Atoi(strings.Trim(strChild, " "))

			m[n][child] = true
		}
	}

	visited = make(map[int]bool)
	cache = make(map[int]int)
	res = 0
	dfs(0)
	fmt.Println(res)
}
