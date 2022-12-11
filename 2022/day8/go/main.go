package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data-real.txt
var input string

type T struct {
	Height   int
	VisibleX bool
	VisibleY bool
}

type State struct {
	Map [][]*T
}

func (s *State) addTree(y int, x int, height int) *T {
	if s.Map[y][x] == nil {
		s.Map[y][x] = &T{}
	}

	s.Map[y][x].Height = height

	return s.Map[y][x]
}

func initMax(size int, value int) []int {
	r := make([]int, size)
	for i := 0; i < len(r); i++ {
		r[i] = value
	}

	return r
}

func initState(rowLen int, colLen int) State {
	m := make([][]*T, rowLen)
	for i := 0; i < rowLen; i++ {
		m[i] = make([]*T, colLen)
	}

	state := State{
		Map: m,
	}

	return state
}

func parseMap(rawData string) (State, error) {
	rawData = strings.TrimRight(rawData, "\n")
	rows := strings.Split(rawData, "\n")

	rowLen := len(rows)
	colLen := len(rows[0])
	state := initState(rowLen, colLen)

	rowMaxTop := initMax(colLen, -1)
	rowMaxBottom := initMax(colLen, -1)
	for rowIdx := 0; rowIdx < len(rows); rowIdx++ {
		// check X
		colMaxLeft := -1
		colMaxRight := -1
		for colIdx := 0; colIdx < len(rows[rowIdx]); colIdx++ {
			height, max, visible, _ := handleValue(rows[rowIdx][colIdx], colMaxLeft)
			colMaxLeft = max
			t := state.addTree(rowIdx, colIdx, height)
			t.VisibleX = t.VisibleX || visible

			height, max, visible, _ = handleValue(rows[rowIdx][len(rows[rowIdx])-colIdx-1], colMaxRight)
			colMaxRight = max
			t = state.addTree(rowIdx, len(rows[rowIdx])-colIdx-1, height)
			t.VisibleX = t.VisibleX || visible

			height, max, visible, _ = handleValue(rows[rowIdx][colIdx], rowMaxTop[colIdx])
			rowMaxTop[colIdx] = max
			t = state.addTree(rowIdx, colIdx, height)
			t.VisibleY = t.VisibleY || visible

			height, max, visible, _ = handleValue(rows[len(rows)-rowIdx-1][colIdx], rowMaxBottom[colIdx])
			rowMaxBottom[colIdx] = max
			t = state.addTree(len(rows)-rowIdx-1, colIdx, height)
			t.VisibleY = t.VisibleY || visible
		}
	}

	return state, nil
}

// height, max, visible, error
func handleValue(column byte, max int) (int, int, bool, error) {
	height, err := strconv.Atoi(string(column))
	if err != nil {
		return 0, 0, false, err
	}

	if height > max {
		return height, height, true, nil
	}

	return height, max, false, nil
}

func handleValueInt(height int, max int) (int, bool) {
	if height < max {
		return max, true
	}

	return height, false
}

func (s *State) scenicScore(x int, y int) int {
	sum := 1
	for _, v := range s.treeDistance(x, y) {
		sum = sum * v
	}

	return sum
}

func (s *State) treeDistance(x int, y int) [4]int {
	visibility := [4]int{0, 0, 0, 0}
	tree := s.Map[y][x]

	topMax, bottomMax := tree.Height, tree.Height
	visibleTop, visibleBottom := true, true

	rowLen := len(s.Map)
	for r := 1; r <= rowLen; r++ {
		if visibleTop && y-r >= 0 {
			top := s.Map[y-r][x]
			topMax, visibleTop = handleValueInt(top.Height, topMax)
			visibility[0] += 1
		}

		if visibleBottom && y+r < rowLen {
			bottom := s.Map[y+r][x]
			bottomMax, visibleBottom = handleValueInt(bottom.Height, bottomMax)
			visibility[1] += 1
		}

		if !visibleTop && !visibleBottom {
			break
		}
	}

	visibleRight, visibleLeft := true, true
	rightMax, leftMax := tree.Height, tree.Height

	colLen := len(s.Map[0])
	for c := 1; c <= colLen; c++ {
		if visibleLeft && x-c >= 0 {
			left := s.Map[y][x-c]
			leftMax, visibleLeft = handleValueInt(left.Height, leftMax)
			visibility[2] += 1
		}

		if visibleRight && x+c < colLen {
			right := s.Map[y][x+c]
			rightMax, visibleRight = handleValueInt(right.Height, rightMax)
			visibility[3] += 1
		}

		if !visibleRight && !visibleLeft {
			break
		}
	}

	return visibility
}

func getAnswer1(s State) int {
	sum := 0
	for _, r := range s.Map {
		for _, c := range r {
			if c.VisibleX || c.VisibleY {
				sum += 1
			}
		}
	}

	return sum
}

func getAnswer2(s State) int {
	max := 0
	for r := 0; r < len(s.Map); r++ {
		for c := 0; c < len(s.Map[r]); c++ {
			val := s.scenicScore(c, r)
			if val > max {
				max = val
			}
		}
	}
	return max
}

func main() {
	state, err := parseMap(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Answer 1: %d\n", getAnswer1(state))
	fmt.Printf("Answer 2: %d\n", getAnswer2(state))
}
