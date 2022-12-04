package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var input string

func generateSet() map[rune]int {
	c := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := map[rune]int{}
	for k, v := range c {
		m[v] = k + 1
	}

	return m
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	d := loadRucksacks(input)
	set := generateSet()

	sum := 0
	for _, sack := range d {
		result := findBajs(sack[0], sack[1])
		for r, _ := range result {
			if _, ok := set[r]; !ok {
				fmt.Println("Not found in set")
			}
			sum += set[r]
		}
	}

	return sum
}

func part2() int {
	data := strings.TrimRight(input, "\n")
	sacks := strings.Split(data, "\n")
	set := generateSet()

	group := []map[rune]bool{}
	sum := 0
	gc := 0
	for _, s := range sacks {
		gc++
		groupItem := map[rune]bool{}
		for _, c := range s {
			groupItem[c] = true
		}
		group = append(group, groupItem)

		if gc%3 == 0 {
			for c := range group[0] {
				if group[1][c] && group[2][c] {
					sum += set[c]
				}
			}
			group = []map[rune]bool{}
		}
	}

	return sum
}

func findBajs(a []rune, b []rune) map[rune]int {
	result := map[rune]int{}

	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				result[bv]++
			}
		}
	}

	return result
}

func loadRucksacks(data string) [][2][]rune {
	data = strings.TrimRight(data, "\n")

	rucksacks := strings.Split(data, "\n")
	result := [][2][]rune{}
	for _, r := range rucksacks {
		if len(r) == 1 {
			continue
		}

		split := len(r) / 2
		rd := [2][]rune{
			[]rune(r[0:split]),
			[]rune(r[split:len(r)]),
		}

		result = append(result, rd)
	}

	return result
}
