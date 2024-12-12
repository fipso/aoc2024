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
	lineLen = len(strings.Split(text, "\n")[0])

	// Find XMAS normal
	wc := findWord("XMAS")
	wc += findWord("SAMX")

	// Find diagonals
	touchedIndexes := make(map[int]bool)
	for i, c := range text {
		if !isRightChar(c) {
			continue
		}
		if touchedIndexes[i] {
			continue
		}
                touchedIndexes[i] = true

                // Explorer neighbours
                if i >= lineLen {
                  topLeft := getDiagNeighbourIndex(i, false, false)
                  topRight := getDiagNeighbourIndex(i, true, false)
                } 

                bottomLeft := getDiagNeighbourIndex(i, false, true)
                bottomRight := getDiagNeighbourIndex(i, true, true)

	}

	fmt.Println("words", wc)

}

func isRightChar(c rune) bool {
	return c == 'X' || c == 'M' || c == 'A' || c == 'S'
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
