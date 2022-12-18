package main

import (
	_ "embed"
	"fmt"
)

//go:embed data-real.txt
var input string

type C struct {
	x int
	y int
}

type M map[C]int

const (
	ROCK = iota + 1
	SAND
	AIR
	SPAWNER
)

func unitsOfSand(m M) int {
	s := 0
	for _, v := range m {
		if v == SAND {
			s++
		}
	}

	return s
}

func pourSand(m M, start C) M {
	dir := []C{
		{0, 1},
		{-1, 1},
		{1, 1},
	}

	end := getEnd(m)
	rest := false
	currPos := start
	for !rest {
		moved := false
		for _, d := range dir {
			tmp := move(currPos, d)
			if !isCollition(m, tmp) {
				currPos = tmp
				moved = true
				break
			}
		}

		if !moved {
			m[currPos] = SAND
			currPos = start
		}

		if currPos.y >= end.y {
			rest = true
		}
	}

	return m
}

func answer1() int {
	return 0
}

func answer2() int {
	return 0
}

func main() {
	m, err := parseData(input)
	if err != nil {
		fmt.Printf("Parsing error: %v", err)
		return
	}

	pourSand(m, C{500, 0})
	render(m)

	fmt.Printf("Answer 1: %d\n", unitsOfSand(m))
	fmt.Printf("Answer 2: %d\n", answer2())
}
