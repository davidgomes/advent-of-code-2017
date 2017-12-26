package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

// 308663 257593 285566 176447 664

type Queue [][]int

func (q *Queue) Push(n []int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n []int) {
	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}

const maxNum = 2000

var matrix [maxNum][maxNum]int8

var directions map[string][]int

/*func dfs(currentPosition []int, finalPosition []int, steps int) {
	if currentPosition[0] >= 1700 || currentPosition[1] >= 1700 || currentPosition[0] <= 0 || currentPosition[1] <= 0 || matrix[currentPosition[0]][currentPosition[1]] == 1 {
		return
	}

	matrix[currentPosition[0]][currentPosition[1]] = 1

	if currentPosition[0] == finalPosition[0] &&
		currentPosition[1] == finalPosition[1] {
		fmt.Println(steps)
		os.Exit(0)
	}

	for _, movement := range directions {
		nextPosition := make([]int, 2)
		copy(nextPosition, currentPosition)
		nextPosition[0] = currentPosition[0] + movement[0]
		nextPosition[1] = currentPosition[1] + movement[1]
		dfs(nextPosition, finalPosition, steps+1)
	}
}*/

func main() {
	for y := 0; y < maxNum; y++ { // rows
		for x := 0; x < maxNum; x++ { // columns
			matrix[y][x] = -1
		}
	}

	directions = map[string][]int{// x, y
		"ne": []int{1, -1},
		"nw": []int{-1, -1},
		"s": []int{0, 2},
		"n": []int{0, -2},
		"se": []int{1, 1},
		"sw": []int{-1, 1},
	}

	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day11.in")

	if err != nil {
		fmt.Print(err)
		return
	}

	str := string(b)
	instructions := strings.Split(str, ",")
	finalPosition := []int{maxNum / 2, maxNum / 2} // x, y

	for _, instruction := range instructions {
		/*fmt.Println(instruction)*/
		finalPosition[0] += directions[instruction][0]
		finalPosition[1] += directions[instruction][1]
	}

	fmt.Println(finalPosition)

	q := Queue{}
	q.Push([]int{maxNum / 2, maxNum / 2, 0})

	for q.Len() > 0 {
		currentPosition := q.Pop()

		//fmt.Println(currentPosition[0], currentPosition[1])

		if currentPosition[0] >= 1700 || currentPosition[1] >= 1700 || currentPosition[0] <= 0 || currentPosition[1] <= 0 || matrix[currentPosition[0]][currentPosition[1]] == 1 {
			continue
		}

		matrix[currentPosition[0]][currentPosition[1]] = 1

		if currentPosition[0] == finalPosition[0] && currentPosition[1] == finalPosition[1] {
			fmt.Println("found")
			fmt.Println(currentPosition[2])
			os.Exit(0)
		}

		for _, movement := range directions {
			nextPosition := make([]int, 3)
			copy(nextPosition, currentPosition)
			nextPosition[0] = currentPosition[0] + movement[0]
			nextPosition[1] = currentPosition[1] + movement[1]
			nextPosition[2] = currentPosition[2] + 1
			q.Push(nextPosition)
		}
	}

	//dfs(currentPosition, finalPosition, 0)
}
