package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day7.in")

	if err != nil {
		fmt.Print(err)
		return
	}

	str := string(b)
	lines := strings.Split(str, "\n")

	m := make(map[string]map[string]bool)
	allChildren := make([]string, 0)

	for _, l := range lines {
		arrowSplit := strings.Split(l, "->")


		if len(arrowSplit) == 2 {
			children := strings.Split(strings.Trim(arrowSplit[1], " "), ",")
			name := strings.Split(arrowSplit[0], " ")[0]

			childrenMap := make(map[string]bool)

			for _, c := range children {
				childrenMap[strings.Trim(c, " ")] = true
			}

			m[name] = childrenMap
		}

		allChildren = append(allChildren, strings.Split(arrowSplit[0], " ")[0])
	}

	// Find root
	var root string
	for _, name := range allChildren {
		isRoot := true

		for _, v := range m {

			if v[name] {
				isRoot = false
			}
		}

		if isRoot {
			root = name
		}
	}

	fmt.Println(root)
}
