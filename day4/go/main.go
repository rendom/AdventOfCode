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
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	return solver(checkRange)
}

func part2() int {
	return solver(overlap)
}

func solver(compare func([]int, []int) bool) (r int) {
	groups, err := parseData(input)
	if err != nil {
		fmt.Println(err)
		return r
	}

	for _, group := range groups {
		if compare(group[0], group[1]) {
			r++
		}
	}

	return r
}

func overlap(r []int, m []int) bool {
	if r[0] <= m[1] && r[1] >= m[0] {
		return true
	}

	return false
}

func checkRange(r []int, m []int) bool {
	if r[0] >= m[0] && r[1] <= m[1] {
		return true
	}

	if m[0] >= r[0] && m[1] <= r[1] {
		return true
	}

	return false
}

func parseData(data string) ([][][]int, error) {
	data = strings.TrimRight(data, "\n")
	rows := strings.Split(data, "\n")

	result := [][][]int{}
	for _, row := range rows {
		groups := strings.Split(row, ",")
		group := [][]int{}
		for _, r := range groups {
			groupRange := []int{}
			for _, d := range strings.Split(r, "-") {
				digit, err := strconv.Atoi(d)
				if err != nil {
					return [][][]int{}, err
				}
				groupRange = append(groupRange, digit)
			}
			group = append(group, groupRange)
		}
		result = append(result, group)
	}

	return result, nil
}
