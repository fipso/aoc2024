package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var text string
var lineLen int

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	text = string(b)
	lineLen = len(strings.Split(text, "\n"))

	// Find XMAS normal
	wc := findWord("XMAS")
	wc += findWord("SAMX")

	fmt.Println("words", wc)

}

func findWord(w string) int {
	wc := regexp.MustCompile(w).FindAllString(text, -1)
	return len(wc)
}

func getDiagNeighbourIndex(index int, right bool, down bool) int {
	if !right && !down {
		return index - lineLen - 1
	}
	if right && !down {
		return index - lineLen + 1
	}
	if right && down {
		return index + lineLen + 1
	}
	if !right && down {
		return index + lineLen - 1
	}

	return -1
}
