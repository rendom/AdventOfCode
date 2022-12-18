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
			if currPos == start {
				rest = true
			}
			currPos = start
		}

		if currPos.y >= end.y {
			rest = true
		}
	}

	return m
}

func answer1() int {
	m, err := parseData(input)
	if err != nil {
		fmt.Printf("Parsing error: %v", err)
		return 0
	}

	pourSand(m, C{500, 0})
	return unitsOfSand(m)
}

func answer2() int {
	m, err := parseData(input)
	if err != nil {
		fmt.Printf("Parsing error: %v", err)
		return 0
	}

	end := getEnd(m)
	m = fillRange(m, C{0, end.y + 2}, C{99999, end.y + 2}, ROCK)
	pourSand(m, C{500, 0})
	return unitsOfSand(m)
}

func main() {

	fmt.Printf("Answer 1: %d\n", answer1())
	fmt.Printf("Answer 2: %d\n", answer2())
}
