package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed data.txt
var input string

var words = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart2(input string) int {
	return parseInput(input, true)
}

func answerPart1(input string) int {
	return parseInput(input, false)
}

func findFirstWord(input string) (int, bool) {
	for nr, w := range words {
		if strings.HasPrefix(input, w) {
			return nr, true
		}
	}

	return -1, false
}

func parseInput(input string, includeWords bool) int {
	input = strings.TrimRight(input, "\n")

	answer := 0

	for _, row := range strings.Split(input, "\n") {
		first := 0
		last := 0
		lineLn := len(row)

		for i := 0; i < lineLn; i++ {
			if unicode.IsDigit(rune(row[i])) {
				last, _ = strconv.Atoi(string(row[i]))

				if first == 0 {
					first = last
				}
			}

			if includeWords {
				if a, ok := findFirstWord(row[i:]); ok {
					if first == 0 {
						first = a
					}

					last = a
				}
			}
		}

		answer += (first*10 + last)
	}

	return answer
}
