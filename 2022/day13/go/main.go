package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

//go:embed data-real.txt
var input string

func parseData(d string) [][]any {
	d = strings.TrimRight(d, "\n")
	groups := [][]any{}
	for _, g := range strings.Split(d, "\n\n") {
		group := []any{}
		for _, c := range strings.Split(g, "\n") {
			var item any
			json.Unmarshal([]byte(c), &item)
			group = append(group, item)
		}
		groups = append(groups, group)
	}

	return groups
}

/*
If both values are integers, the lower integer should come first.
If the left integer is lower than the right integer, the inputs are in the right order.
If the left integer is higher than the right integer, the inputs are not in the right order.
Otherwise, the inputs are the same integer; continue checking the next part of the input.

If both values are lists, compare the first value of each list, then the second value,
and so on. If the left list runs out of items first, the inputs are in the right order. <--
If the right list runs out of items first, the inputs are not in the right order.

If the lists are the same length and no comparison makes a decision about the order,
continue checking the next part of the input.

If exactly one value is an integer, convert the integer to a list which contains that integer as its only value,
then retry the comparison. For example, if comparing [0,0,0] and 2,
convert the right value to [2] (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].
*/

func compareValue(a any, b any) int {
	aSlice, aIsSlice := a.([]any)
	bSlice, bIsSlice := b.([]any)

	if !aIsSlice && !bIsSlice {
		if a.(float64) < b.(float64) {
			return 1
		} else if a.(float64) == b.(float64) {
			return 0
		} else {
			return -1
		}
	}

	if !aIsSlice {
		aSlice = []any{a}
	}

	if !bIsSlice {
		bSlice = []any{b}
	}

	aLen := len(aSlice)
	bLen := len(bSlice)

	// if aLen > bLen {
	// 	return -1
	// }

	ln := min(aLen, bLen)
	for i := 0; i < ln; i++ {
		c := compareValue(aSlice[i], bSlice[i])
		if c != 0 {
			return c
		}
	}

	if aLen < bLen {
		return 1
	} else if aLen > bLen {
		return -1
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func answer1(d [][]any) int {
	correct := 0
	for i, g := range d {
		if compareValue(g[0], g[1]) == 1 {
			correct += i + 1
		}
	}

	return correct
}

func answer2(d [][]any) int {
	all := []any{
		float64(2),
		float64(6),
	}

	for _, v := range d {
		all = append(all, v...)
	}

	sort.SliceStable(all, func(i, j int) bool {
		c := compareValue(all[i], all[j])
		return c > 0
	})

	result := 1
	for i, v := range all {
		if v == 2.0 || v == 6.0 {
			result *= i + 1
		}
	}

	return int(result)
}

// 1958 to low
// 6123 to high
func main() {
	b := parseData(input)
	fmt.Printf("Answer 1: %d\n", answer1(b))
	fmt.Printf("Answer 2: %d\n", answer2(b))
}
