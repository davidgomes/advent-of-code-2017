package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"crypto/sha1"
	"encoding/binary"
)

func main() {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd + "/main/day6.in")

	if err != nil {
		fmt.Print(err)
		return
	}

	str := string(b)
	sections := strings.Split(str, " ")

	banks := make([]uint32, len(sections))

	for i, l := range sections {
		n, err := strconv.Atoi(l)

		if err != nil {
			fmt.Print(err)
			return
		}

		banks[i] = uint32(n)
	}

	seen := make(map[string]bool)
	count := 0

	for {
		// Hash and see if we've been here before
		h := sha1.New()

		for _, n := range banks {
			bs := make([]byte, 4)
			binary.LittleEndian.PutUint32(bs, n)
			h.Write(bs)
		}

		hash := string(h.Sum(nil))

		if seen[hash] {
			fmt.Println(count)
			break
		} else {
			seen[hash] = true
			count++
		}

		// Find largest number
		largest := 0
		for i:=1; i < len(banks); i++ {
			if banks[i] > banks[largest] {
				largest = i
			}
		}

		// Distribute extra blocks
		toDistribute := banks[largest]

		var current int
		if largest == len(banks) - 1 {
			current = 0
		} else {
			current = largest + 1
		}

		banks[largest] = 0
		for toDistribute > 0 {
			banks[current]++
			toDistribute--
			if current == len(banks) - 1 {
				current = 0
			} else {
				current = current + 1
			}
		}
	}
}
