package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

var mm map[string]int

func getRegister2(reg string) int {
	_, present := mm[reg]

	if present {
		return mm[reg]
	} else {
		mm[reg] = 0
		return 0
	}
}

func main() {
	pwd, _ := os.Getwd()
	b, _ := ioutil.ReadFile(pwd + "/main/day23.in")

	str := string(b)
	lines := strings.Split(str, "\n")

	mm = make(map[string]int)

	var result int

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
				mm[a] = bi
			} else {
				mm[a] = getRegister2(b)
			}
		} else if cmd == "sub" {
			fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
			bi, err := strconv.Atoi(b)

			if err == nil {
				mm[a] -= bi
			} else {
				mm[a] -= getRegister2(b)
			}
		} else if cmd == "mul" {
			fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
			bi, err := strconv.Atoi(b)
			result++

			if err == nil {
				mm[a] *= bi
			} else {
				mm[a] *= getRegister2(b)
			}

			fmt.Println(result)
		} else if cmd == "jnz" { // 76794
			ai, err := strconv.Atoi(a)
			//fmt.Println("jump command", getRegister2(a), ai)

			if (err == nil && ai != 0) || getRegister2(a) != 0 {
				fmt.Println("jumping")
				fmt.Sscanf(line, "%s %s %s", &tmp, &tmp, &b)
				bi, err := strconv.Atoi(b)

				if err == nil {
					i += bi
				} else {
					i += getRegister2(b)
				}

				incr = false
			}
		}

		fmt.Println(mm)

		if incr {
			i++
		}
	}

	fmt.Println(result)
}
