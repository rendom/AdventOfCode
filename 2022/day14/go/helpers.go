package main

import (
	"math"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getStart(m M) C {
	startX := math.MaxInt
	startY := math.MaxInt

	for p, _ := range m {
		startX = min(p.x, startX)
		startY = min(p.y, startY)
	}

	return C{startX, startY}
}

func getEnd(m M) C {
	startX := 0
	startY := 0

	for p, _ := range m {
		startX = max(p.x, startX)
		startY = max(p.y, startY)
	}

	return C{startX, startY}
}

func isCollition(m M, c C) bool {
	n, _ := m[c]
	return n != 0
}

func move(pos C, dir C) C {
	pos.x += dir.x
	pos.y += dir.y
	return pos
}
