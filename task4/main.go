package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	valids := 0
	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		var lvls []int
		parts := strings.Split(line, " ")
		for _, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			lvls = append(lvls, n)
		}

		inc := false
		fails := 0
		for i, lvl := range lvls {
			if i < 1 {
				continue
			}

			diff := math.Abs(float64(lvl - lvls[i-1]))
			if diff < 1 || diff > 3 {
				fails++
				continue
			}

			if i == 1 {
				inc = lvl > lvls[i-1]
			} else if (inc && lvl < lvls[i-1]) || (!inc && lvl > lvls[i-1]) {
				fails++
			}

		}

		if fails > 1 {
			continue
		}

		valids++
	}

	fmt.Println(valids)
}
