package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	total := 0

	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	s := string(b)

	allDos := regexp.MustCompile(`do\(\)`).FindAllStringSubmatchIndex(s, -1)
	allDonts := regexp.MustCompile(`don't\(\)`).FindAllStringSubmatchIndex(s, -1)

	for _, dontIndex := range allDonts {
		// Find next doIndex
		nextDoIndex := -1
		for _, doIndex := range allDos {
			if doIndex[0] > dontIndex[0] {
				nextDoIndex = doIndex[0]
				break
			}
		}

		// Fill area between don't and do with spaces
		if nextDoIndex != -1 {
			s = s[:dontIndex[0]] + strings.Repeat(" ", nextDoIndex-dontIndex[0]) + s[nextDoIndex:]
		} else {
			s = s[:dontIndex[0]] + strings.Repeat(" ", len(s)-dontIndex[0])
		}

	}

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	match := r.FindAllStringSubmatch(s, -1)
	for _, m := range match {
		in := m[0][4:]
		in = in[:len(in)-1]
		parts := strings.Split(in, ",")
		n1, _ := strconv.Atoi(parts[0])
		n2, _ := strconv.Atoi(parts[1])

		total += n1 * n2
	}

	println(total)
}
