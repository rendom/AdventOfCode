package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func parse(input string) [][]int {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")

	result := [][]int{}
	for _, row := range rows {
		vals := strings.Split(row, " ")
		row := []int{}
		for _, v := range vals {
			val, _ := strconv.Atoi(v)
			row = append(row, val)
		}

		result = append(result, row)
	}

	return result
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}

	return v
}

func validatePart1Row(row []int) bool {
	dir := row[0] > row[1]
	for i := 1; i < len(row); i++ {
		if !validate(row[i-1], row[i], dir) {
			return false
		}
	}

	return true
}

func validate(first int, sec int, dir bool) bool {
	ldir := first > sec
	if dir != ldir {
		return false
	}

	diff := abs(first - sec)
	if diff < 1 || diff > 3 {
		return false
	}

	return true
}

func getAnswer(input string, allowOneFailure bool) int {
	list := parse(input)

	ok := 0
	for _, row := range list {
		if validatePart1Row(row) {
			ok++
		} else if allowOneFailure {
			for r := range len(row) {
				tmp := make([]int, len(row))
				copy(tmp, row)
				tmp = append(tmp[:r], tmp[r+1:]...)
				fmt.Println(r, row, tmp, row[r])
				if validatePart1Row(tmp) {
					ok++
					break
				}
			}
		}
	}

	return ok
}

func answerPart1(input string) int {
	return getAnswer(input, false)
}

func answerPart2(input string) int {
	return getAnswer(input, true)
}
