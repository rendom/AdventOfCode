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

func parse(input string) (rules map[[2]int]bool, updates [][]int) {
	input = strings.TrimRight(input, "\n")
	sections := strings.Split(input, "\n\n")
	rules = make(map[[2]int]bool)

	for _, r := range strings.Split(sections[0], "\n") {
		rr := strings.Split(r, "|")
		before, err := strconv.Atoi(rr[0])
		if err != nil {
			panic(err)
		}

		after, err := strconv.Atoi(rr[1])
		if err != nil {
			panic(err)
		}

		rules[[2]int{before, after}] = true
	}

	for _, r := range strings.Split(sections[1], "\n") {
		pages := []int{}
		for _, number := range strings.Split(r, ",") {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			pages = append(pages, n)
		}
		updates = append(updates, pages)
	}

	return rules, updates
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func solve(input string, partOne bool) int {
	rules, updates := parse(input)
	result := 0
	for _, pages := range updates {
		if value, alreadySorted := checkRow(pages, rules); alreadySorted == partOne {
			result += value
		}
	}

	return result
}

func answerPart1(input string) int {
	return solve(input, true)
}

func checkRow(pages []int, rules map[[2]int]bool) (int, bool) {
	fn := func(a int, b int) int {
		if _, ok := rules[[2]int{a, b}]; ok {
			return -1
		} else if _, ok := rules[[2]int{b, a}]; ok {
			return 1
		}

		return 0
	}

	alreadySorted := slices.IsSortedFunc(pages, fn)

	if alreadySorted {
		return pages[len(pages)/2], true
	}

	slices.SortFunc(pages, fn)
	return pages[len(pages)/2], false
}

func answerPart2(input string) int {
	return solve(input, false)

}
