package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data-real.txt
var input string

const START = 'S'
const END = 'E'

type Node struct {
	p      Point
	parent *Node
}

type Point struct {
	x int
	y int
}

func canVisit(a, c rune) bool {
	// End's height = z
	if c == END {
		c = 'z'
	}

	return c-1 <= a || a == 'S'
}

func findRunePosition(rows []string, find rune) (int, int, error) {
	for y, row := range rows {
		for x, c := range row {
			if find == c {
				return x, y, nil
			}
		}
	}

	return -1, -1, fmt.Errorf("unable to find %s", find)
}

func findNeighborInMap(m []string, pos Point) []Node {
	d := []Point{
		{-1, 0}, // Left
		{1, 0},  // Right
		{0, 1},  // Top
		{0, -1}, // Bottom
	}

	yMax := len(m)
	xMax := len(m[0])
	result := []Node{}
	for _, p := range d {
		checkX := pos.x + p.x
		checkY := pos.y + p.y

		if checkX < 0 || checkX >= xMax || checkY < 0 || checkY >= yMax {
			continue
		}

		can := canVisit(
			rune(m[pos.y][pos.x]),
			rune(m[checkY][checkX]),
		)

		if can {
			result = append(result, Node{p: Point{pos.x + p.x, pos.y + p.y}})
		}
	}

	return result
}

func bfs(m []string, start Point, end Point) (*Node, error) {
	queue := Queue{}
	visits := map[Point]bool{}
	queue.append(Node{p: start})
	visits[start] = true

	for len(queue) > 0 {
		current := queue.shift()

		if current.p == end {
			return &current, nil
		}

		nbs := findNeighborInMap(m, current.p)
		for _, nb := range nbs {
			if _, visited := visits[nb.p]; !visited {
				visits[nb.p] = true
				nb.parent = &current
				queue.append(nb)
			}
		}
	}

	return nil, fmt.Errorf("Unable to find end")
}

func getLength(m []string, node Node) int {
	n := 0
	current := &node
	for current.parent != nil {
		current = current.parent
		n++
	}

	return n
}

func getData() []string {
	return strings.Split(
		strings.TrimRight(input, "\n"),
		"\n",
	)
}

func answer1() {
	m := getData()
	x, y, err := findRunePosition(m, 'S')
	if err != nil {
		fmt.Printf("Unable to find start\n")
		return
	}
	start := Point{x, y}

	x, y, err = findRunePosition(m, 'E')
	if err != nil {
		fmt.Println("Unable to find end")
		return
	}
	end := Point{x, y}

	shortest, err := bfs(m, start, end)
	if err != nil {
		fmt.Printf("Unable to find solution\n")
		return
	}

	fmt.Printf("Answer 1: %d\n", getLength(m, *shortest))
}

func main() {
	answer1()
}
