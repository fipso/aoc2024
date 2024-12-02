package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputB, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	inputS := string(inputB)
	lines := strings.Split(inputS, "\n")
	var nLeft, nRight []int
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "   ")
		nl, _ := strconv.Atoi(parts[0])
		nr, _ := strconv.Atoi(parts[1])
		nLeft = append(nLeft, nl)
		nRight = append(nRight, nr)
	}

	totalSim := 0

	for _, n := range nLeft {
		count := 0
		for _, m := range nRight {
			if n == m {
				count++
			}
		}
                totalSim += n * count
	}

	fmt.Println("Total Similarity Score: ", totalSim)
}
