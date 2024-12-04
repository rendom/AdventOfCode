package main

import (
	_ "embed"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func parseVal(n string) ([]int, error) {
	numbers := []int{}
	for _, v := range strings.Split(n, ",") {
		number, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func parseOps(input string) ([][]int, error) {
	r, err := regexp.Compile(`don't\(\)|do\(\)|mul\([0-9,]+\)`)
	if err != nil {
		return nil, err
	}

	m := r.FindAllString(input, -1)

	result := [][]int{}
	add := true
	for _, m := range m {
		if len(m) < 2 {
			return nil, errors.New("Hm.. regex failed")
		}

		switch {
		case m[:3] == "mul":
			var n1 int = 1
			var n2 int = 1
			fmt.Sscanf(m, "mul(%d,%d)", &n1, &n2)

			if add {
				result = append(result, []int{n1, n2})
			}
		case m[:3] == "do(":
			add = true
		case m[:5] == "don't":
			add = false
		}
	}

	return result, nil
}

func parseMul(input string) ([][]int, error) {
	r, err := regexp.Compile(`mul\(([0-9]+,[0-9]+)\)`)
	if err != nil {
		return nil, err
	}

	m := r.FindAllStringSubmatch(input, -1)
	result := [][]int{}
	for _, str := range m {
		if len(str) < 2 {
			return nil, errors.New("Hm.. regex failed")
		}

		numbers := []int{}
		for _, v := range strings.Split(str[1], ",") {
			number, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			numbers = append(numbers, number)
		}

		result = append(result, numbers)
	}

	return result, nil
}

func calcAnswer1(l [][]int) int {
	result := 0
	for _, numbers := range l {
		n := 1
		for _, v := range numbers {
			n *= v
		}

		result += n
	}

	return result
}

func answerPart1(input string) int {
	list, err := parseMul(input)
	if err != nil {
		panic(err)
	}

	return calcAnswer1(list)
}

func answerPart2(input string) int {
	list, err := parseOps(input)
	if err != nil {
		panic(err)
	}

	return calcAnswer1(list)
}
