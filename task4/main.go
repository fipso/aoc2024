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

LINE_LOOP:
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

		//valid := true
		for skipIndex := -1; skipIndex < len(lvls); skipIndex++ {
			// Copy lvls but skip current Index
			lvlsCopy := make([]int, 0)
			for i, v := range lvls {
				if i == skipIndex {
					continue
				}

				lvlsCopy = append(lvlsCopy, v)
			}

			inc := false
			valid := true
			for i, lvl := range lvlsCopy {
				if i < 1 {
					continue
				}

				diff := math.Abs(float64(lvl - lvlsCopy[i-1]))
				if diff < 1 || diff > 3 {
					valid = false
					break
				}

				if i == 1 {
					inc = lvl > lvlsCopy[i-1]
				} else if (inc && lvl < lvlsCopy[i-1]) || (!inc && lvl > lvlsCopy[i-1]) {
					valid = false
					break
				}
			}

			fmt.Println(valid, skipIndex, lvlsCopy)
			if valid {
				valids++
				continue LINE_LOOP
			}

		}

	}

	fmt.Println(valids)
}
