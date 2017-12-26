package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day5.in") // just pass the file name

	if err != nil {
		fmt.Print(err)
		return
	}

	str := string(b) // convert content to a 'string'
	lines := strings.Split(str, "\n")

	nums := make([]int, len(lines))

	for i, l := range lines {
		n, err := strconv.Atoi(l)

		if err != nil {
			fmt.Print(err)
			return
		}

		nums[i] = n
	}

	i := 0
	count := 0
	for {
		if i < 0 || i >= len(nums) {
			fmt.Println(count)
			return
		}

		i, nums[i] = i + nums[i], nums[i] + 1

		count++
	}
}
