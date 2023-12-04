package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Map map[Pos]C

const SYMBOL = "symbol"
const EMPTY = "empty"
const DIGIT = "digit"

type Pos struct {
	x int
	y int
}

type C struct {
	Width           int
	Type            string
	Value           string
	ConnectedDigits []Pos
}

//go:embed data.txt
var input string

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart2(input string) int {
	m, err := parseInput(input)
	if err != nil {
		return -1
	}

	sum := 0
	for pos, v := range m {
		if v.Type == SYMBOL && v.Value == "*" {
			sum += findConnectedDigits(m, pos)
		}
	}

	return sum
}

func findConnectedDigits(m Map, pos Pos) int {
	values := []int{}
	for mPos, c := range m {
		if c.Type != DIGIT {
			continue
		}

		if mPos.y >= pos.y-1 && mPos.y <= pos.y+1 && ((mPos.x >= pos.x-1 && mPos.x <= pos.x+1) ||
			(mPos.x+c.Width-1 >= pos.x-1 && mPos.x+c.Width-1 <= pos.x+1)) {
			val, err := strconv.Atoi(c.Value)
			if err != nil {
				panic(err)
			}
			values = append(values, val)
		}
	}

	if len(values) < 2 {
		return 0
	}

	sum := 1
	for _, v := range values {
		sum *= v
	}
	return sum
}

func answerPart1(input string) int {
	m, err := parseInput(input)
	if err != nil {
		return -1
	}

	sum, err := getSumOfConnectedDigits(m)
	if err != nil {
		return -1
	}

	return sum
}

func getSumOfConnectedDigits(m Map) (int, error) {
	sum := 0
	for p, v := range m {
		if v.Type == DIGIT && isConnected(m, p, v.Width) {
			v, err := strconv.Atoi(v.Value)
			if err != nil {
				return 0, err
			}
			sum += v
		}
	}

	return sum, nil
}

func isConnected(m Map, checkPos Pos, width int) bool {
	dir := []Pos{
		{-1, 0},
		{1, 0},

		{0, 1},
		{0, -1},

		{1, 1},
		{-1, 1},

		{1, -1},
		{-1, -1},
	}

	isConnected := false
	for w := 0; w < width; w++ {
		for _, d := range dir {
			check := Pos{
				x: checkPos.x + w + d.x,
				y: checkPos.y + d.y,
			}

			if v, ok := m[check]; ok {
				if v.Type == SYMBOL {
					isConnected = true
				}
			}
		}
	}

	return isConnected
}

func parseInput(input string) (Map, error) {
	input = strings.TrimRight(input, "\n")

	m := Map{}
	for y, row := range strings.Split(input, "\n") {
		width := 0
		value := ""

		rowLen := len(row)
		for x := 0; x <= rowLen+1; x++ {
			t := "none"
			currVal := '?'
			if x < rowLen {
				currVal = rune(row[x])
				t = getTypeFromChar(currVal)
			}

			if t == DIGIT {
				if t == DIGIT {
					value = value + string(currVal)
					width++
				}

				if rowLen <= x+1 || getTypeFromChar(rune(row[x+1])) != DIGIT {
					m[Pos{(x - width) + 1, y}] = C{
						Width: width,
						Type:  t,
						Value: value,
					}

					width = 0
					value = ""
				}
				continue
			}

			m[Pos{x - width, y}] = C{
				Width: width,
				Type:  t,
				Value: string(currVal),
			}
		}
	}

	return m, nil
}

func getTypeFromChar(c rune) string {
	t := ""
	if c == '.' {
		t = EMPTY
	} else if unicode.IsDigit(c) {
		t = DIGIT
	} else {
		t = SYMBOL
	}

	return t
}
