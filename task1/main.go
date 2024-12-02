package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

	sort.Ints(nLeft)
	sort.Ints(nRight)

	totalDiff := 0

	// Diff
	for i := 0; i < len(nLeft); i++ {
		totalDiff += int(math.Abs(float64(nLeft[i]) - float64(nRight[i])))
	}

	fmt.Println("Total diff: ", totalDiff)
}
