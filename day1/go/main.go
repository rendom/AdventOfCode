package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

func main() {
	answer, err := part1()
	if err != nil {
		fmt.Printf("Part 1: Error, %v\n", err)
		return
	}

	fmt.Printf("Part 1: %d\n", answer)

	answer, err = part2()
	if err != nil {
		fmt.Printf("Part 2: Error, %v\n", err)
		return
	}
	fmt.Printf("Part 2: %d\n", answer)
}

func part1() (int, error) {
	cals, err := sumSliceFromData(input)
	if err != nil {
		return 0, err
	}

	return cals[0], nil
}

func part2() (int, error) {
	cals, err := sumSliceFromData(input)
	if err != nil {
		return 0, err
	}

	amountOfElfs := 3
	if l := len(cals); l < 3 {
		amountOfElfs = l
	}

	cals = cals[0:amountOfElfs]

	sum := 0
	for _, v := range cals {
		sum = sum + v
	}

	return sum, nil
}

func sumSliceFromData(d string) ([]int, error) {
	d = strings.TrimRight(d, "\n")
	byElf := strings.Split(d, "\n\n")
	result := []int{}

	for _, calsRaw := range byElf {
		cals := strings.Split(calsRaw, "\n")
		sum := 0
		for _, cal := range cals {
			v, err := strconv.Atoi(cal)
			if err != nil {
				return result, fmt.Errorf("Unable to convert string to int, %v", err)
			}

			sum = sum + v
		}

		result = append(result, sum)
	}

	if len(byElf) == 0 {
		return result, fmt.Errorf("No elfs found")
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[j] < result[i]
	})

	return result, nil
}
