package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

var grid [][]string
var path string
var steps int

func canGo(x int, y int) bool {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return false
	}

	return grid[y][x] != " " && grid[y][x] != ""
}

func move(position []int, direction string) {
	x, y := position[0], position[1]

	if grid[y][x] != "|" && grid[y][x] != "+" && grid[y][x] != "-" {
		path += grid[y][x]
	}

	fmt.Println(x, y)

	if grid[y][x] == " " || grid[y][x] == "" {
		fmt.Println("i fucked up")

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if position[0] == x && position[1] == y {
					fmt.Printf("#")
				} else {
					fmt.Printf("%s", grid[y][x])
				}
			}
			fmt.Println()
		}

		fmt.Println()
	}

	steps++
	if direction == "down" {
		if canGo(x, y+1) {
			move([]int{x, y + 1}, "down")
		} else {
			if canGo(x-1, y) {
				move([]int{x - 1, y}, "left")
			} else if canGo(x+1, y) {
				move([]int{x + 1, y}, "right")
			}
		}
	} else if direction == "up" {
		if canGo(x, y-1) {
			move([]int{x, y - 1}, "up")
		} else {
			if canGo(x-1, y) {
				move([]int{x - 1, y}, "left")
			} else if canGo(x+1, y) {
				move([]int{x + 1, y}, "right")
			}
		}
	} else if direction == "left" {
		if canGo(x-1, y) {
			move([]int{x - 1, y}, "left")
		} else {
			if canGo(x, y+1) {
				move([]int{x, y + 1}, "down")
			} else if canGo(x, y-1) {
				move([]int{x, y - 1}, "up")
			}
		}
	} else if direction == "right" {
		if canGo(x+1, y) {
			move([]int{x + 1, y}, "right")
		} else {
			if canGo(x, y+1) {
				move([]int{x, y + 1}, "down")
			} else if canGo(x, y-1) {
				move([]int{x, y - 1}, "up")
			}
		}
	}
}

func main() {
	fmt.Println("Hello day 19")

	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day19.in")

	str := string(b)
	lines := strings.Split(str, "\n")

	longestLine := 0
	for _, line := range lines {
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}

	grid = make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, longestLine)
		for x, c := range line {
			grid[y][x] = string(c)
		}
	}

	steps = 0
	startingPoint := []int{89, 0}
	move(startingPoint, "down")
	fmt.Println(path, steps)
}
