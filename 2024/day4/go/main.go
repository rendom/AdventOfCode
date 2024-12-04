package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var DIR = [][2]int{
	{-1, 0},
	{1, 0},

	{0, 1},
	{0, -1},

	{1, 1},
	{-1, -1},

	{1, -1},
	{-1, 1},
}

//go:embed data.txt
var input string

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func findWordInDir(cursor [2]int, m [][]rune, dir [2]int, findWord string) bool {
	maxY := len(m)
	maxX := len(m[0])

	for _, c := range findWord {
		if cursor[0] < 0 || cursor[1] < 0 {
			return false
		}

		if cursor[0] >= maxY || cursor[1] >= maxX {
			return false
		}

		if c != m[cursor[0]][cursor[1]] {
			return false
		}

		cursor[0] += dir[0]
		cursor[1] += dir[1]
	}

	return true
}

func searchMap(m [][]rune, findWord string) int {
	result := 0

	for y, rows := range m {
		for x, _ := range rows {
			for _, move := range DIR {
				cursor := [2]int{y, x}
				if findWordInDir(cursor, m, move, findWord) {
					result++
				}
			}
		}
	}

	return result
}

func parseInput(input string) [][]rune {
	input = strings.TrimRight(input, "\n")
	result := [][]rune{}
	for _, row := range strings.Split(input, "\n") {
		result = append(result, []rune(row))
	}

	return result
}

func answerPart1(input string) int {
	m := parseInput(input)
	return searchMap(m, "XMAS")
}

func answerPart2(input string) int {
	return 0
}