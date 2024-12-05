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
