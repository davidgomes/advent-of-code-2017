package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

/*
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
 */

func main() {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day8.in")

	if err != nil {
		fmt.Print(err)
		return
	}

	str := string(b)
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		fmt.Println(splitLine)
	}
}
