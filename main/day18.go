package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

var m map[string]int

func getRegister(reg string) int {
	_, present := m[reg]

	if present {
		return m[reg]
	} else {
		m[reg] = 0
		return 0
	}
}

func main() {
	fmt.Println("Hello day 18")

	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day18.in")

	str := string(b)
	lines := strings.Split(str, "\n")

	m = make(map[string]int)

	var lastPlayed int

	i := 0
	for i < len(lines) {
		line := lines[i]
		var cmd, a, tmp, b string

		fmt.Sscanf(line, "%s %s", &cmd, &a)
		fmt.Println(line)

		incr := true

		if cmd == "set" {
			fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
			bi, err := strconv.Atoi(b)

			if err == nil {
				m[a] = bi
			} else {
				m[a] = getRegister(b)
			}
		} else if cmd == "add" {
			fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
			bi, err := strconv.Atoi(b)

			if err == nil {
				m[a] += bi
			} else {
				m[a] = getRegister(a) + bi
			}
		} else if cmd == "mul" {
			fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
			bi, err := strconv.Atoi(b)

			if err == nil {
				m[a] *= bi
			} else {
				m[a] *= getRegister(b)
			}
		} else if cmd == "mod" {
			fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
			bi, err := strconv.Atoi(b)

			if err == nil {
				m[a] %= bi
			} else {
				m[a] %= getRegister(b)
			}
		} else if cmd == "snd" {
			lastPlayed = getRegister(a)
		} else if cmd == "rcv" {
			if getRegister(a) != 0 {
				fmt.Println(lastPlayed)
				os.Exit(0)
			}
		} else if cmd == "jgz" {
			if getRegister(a) != 0 {
				fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
				bi, err := strconv.Atoi(b)

				fmt.Println(b)

				if err == nil {
					i += bi
				} else {
					i += getRegister(b)
				}

				incr = false
			}
		}

		fmt.Println(m)

		if incr {
			i++
		}
	}

	fmt.Println(m)
}
