package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func parse(input string) [2][]int {
	input = strings.TrimRight(input, "\n")
	rows := strings.Split(input, "\n")

	result := [2][]int{}

	for _, row := range rows {
		cols := strings.Split(row, "   ")

		c1, _ := strconv.Atoi(cols[0])
		c2, _ := strconv.Atoi(cols[1])

		result[0] = append(result[0], c1)
		result[1] = append(result[1], c2)
	}

	return result
}

func abs(val int) int {
	if val > 0 {
		return val
	}

	return val * -1
}

func answerPart1(input string) int {
	lists := parse(input)
	slices.Sort(lists[0])
	slices.Sort(lists[1])

	sum := 0
	for idx, c1 := range lists[0] {
		sum += abs(c1 - lists[1][idx])
	}

	return sum
}

func findLen(find int, list []int) int {
	l := 0

	for _, v := range list {
		if v == find {
			l++
		}
	}

	return l
}

func answerPart2(input string) int {
	lists := parse(input)
	sum := 0

	for _, c1 := range lists[0] {
		sum += c1 * findLen(c1, lists[1])
	}

	return sum
}
