package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

var used map[int]bool
var comps [][]int
var bestScore int

func dfs2(last []int, side int, score int) {
	//fmt.Println(last)

	//fmt.Println(score, last)
	if score > bestScore {
		bestScore = score
	}

	var toCheckSide int

	if side == 0 {
		toCheckSide = 1
	} else {
		toCheckSide = 0
	}

	for i, comp := range comps {
		isUsed, pres := used[i]

		if (pres && !isUsed) || !pres {
			if last[toCheckSide] == comp[0] {
				used[i] = true
				dfs2(comp, 0, score+comp[0]+comp[1])
				used[i] = false
			} else if last[toCheckSide] == comp[1] {
				used[i] = true
				dfs2(comp, 1, score+comp[0]+comp[1])
				used[i] = false
			}
		}
	}
}

func main() {
	fmt.Println("Hello World")

	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day24.in")

	str := string(b)
	lines := strings.Split(str, "\n")

	comps = make([][]int, 0)
	bestScore = 0

	for _, line := range lines {
		var i, o int
		fmt.Sscanf(line, "%d/%d", &i, &o)

		comps = append(comps, []int{i, o})
	}

	for i, comp := range comps {
		used = make(map[int]bool, 0)

		if comp[0] == 0 {
			used[i] = true
			dfs2(comp, 0, comp[0]+comp[1])
			used[i] = false
		} else if comp[1] == 0 {
			used[i] = true
			dfs2(comp, 1, comp[0]+comp[1])
			used[i] = false
		}
	}

	fmt.Println(bestScore)
}
