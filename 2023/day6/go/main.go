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

const A = 2

type Race struct {
	Time     int
	Distance int
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func m(v []int) int {
	r := 1
	for _, s := range v {
		r *= s
	}

	return r
}

func (r *Race) getSolutions() []int {
	results := []int{}
	for n := 1; n < r.Time; n++ {
		distance := (r.Time - n) * n
		if distance > r.Distance {
			results = append(results, n)
		}
	}
	return results
}

func answerPart2(input string) int {
	r := parseInputP2(input)

	answer := 1
	for _, v := range r {
		answer *= len(v.getSolutions())
	}

	return answer

}

func answerPart1(input string) int {
	r := parseInput(input)

	answer := 1
	fmt.Println(r[0].getSolutions())
	for _, v := range r {
		answer *= len(v.getSolutions())
	}

	return answer
}

func parseInputP2(input string) []Race {
	race := Race{}
	input = strings.TrimRight(input, "\n")
	for currR, row := range strings.Split(input, "\n") {
		value := ""
		for _, v := range row {
			if unicode.IsDigit(rune(v)) {
				value = value + string(v)
			}
		}

		nr, _ := strconv.Atoi(value)
		if currR == 0 {
			race.Time = nr
		} else {
			race.Distance = nr
		}
	}

	return []Race{race}
}

func parseInput(input string) []Race {
	races := []Race{}
	for y, row := range strings.Split(input, "\n") {
		split := strings.Split(row, " ")
		currR := 0
		for _, v := range split {
			if v != "" && unicode.IsDigit(rune(v[0])) {
				nr, _ := strconv.Atoi(v)
				if y == 0 {
					races = append(races, Race{Time: nr})
				} else {
					races[currR].Distance = nr
					currR++
				}
			}
		}
	}

	return races
}
