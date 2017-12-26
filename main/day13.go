package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day13.in")

	str := string(b)
	lines := strings.Split(str, "\n")

	m := make(map[int]int)
	scanners := make(map[int]int)
	goingDown := make(map[int]bool)

	maxDepth := -1
	for _, line := range lines {
		splitLine := strings.Split(line, ":")

		d, _ := strconv.Atoi(strings.Trim(splitLine[0], " "))
		r, _ := strconv.Atoi(strings.Trim(strings.Trim(splitLine[1], " "), " "))

		if d > maxDepth {
			maxDepth = d
		}

		scanners[d] = 0
		goingDown[d] = true

		m[d] = r
	}

	severity := 0
	for pos:=0; pos <= maxDepth; pos++ {
		scannerPosition, present := scanners[pos]

		if present && scannerPosition == 0 {
			severity += pos * m[pos]
		}

		for key, _ := range scanners {
			if goingDown[key] { // going down
				if scanners[key] == m[key] - 1 {
					goingDown[key] = false
					scanners[key]--
				} else {
					scanners[key]++
				}
			} else { // going up
				if scanners[key] == 0 {
					goingDown[key] = true
					scanners[key]++
				} else {
					scanners[key]--
				}
			}
		}
	}

	fmt.Println(severity)
}
