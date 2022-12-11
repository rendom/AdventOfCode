package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed data-real.txt
var input string

type Pos struct {
	y int
	x int
}

type State struct {
	Visited map[Pos]int // y, x, count

	Head Pos
	Tail Pos
}

func (s *Pos) Move(dir string, steps int) {
	switch dir {
	case "R":
		s.x += steps
	case "L":
		s.x -= steps
	case "U":
		s.y -= steps
	case "D":
		s.y += steps
	}
}

func isTouching(a Pos, b Pos) bool {
	cx := math.Abs(float64(a.x - b.x))
	cy := math.Abs(float64(a.y - b.y))

	if cx <= 1 && cy <= 1 {
		return true
	}

	return false
}

func (s *State) Move(dir string, steps int) {
	for i := 1; i <= steps; i++ {
		s.Head.Move(dir, 1)
		// s.Visited[s.Head]++
		if !isTouching(s.Head, s.Tail) {
			if s.Head.x > s.Tail.x {
				s.Tail.x++
			} else if s.Head.x < s.Tail.x {
				s.Tail.x--
			}

			if s.Head.y > s.Tail.y {
				s.Tail.y++
			} else if s.Head.y < s.Tail.y {
				s.Tail.y--
			}
			s.Visited[s.Tail]++
		}
	}
}

func parse(data string) (State, error) {
	s := State{
		Head:    Pos{0, 0},
		Tail:    Pos{0, 0},
		Visited: map[Pos]int{},
	}

	s.Visited[Pos{0, 0}] = 1

	for _, r := range strings.Split(data, "\n") {
		if r == "" {
			continue
		}

		d := strings.Split(r, " ")

		steps, err := strconv.Atoi(d[1])
		if err != nil {
			return s, err
		}

		s.Move(d[0], steps)
	}

	return s, nil
}

func answer1(s State) int {
	sum := 0
	for _, v := range s.Visited {
		if v > 0 {
			sum++
		}
	}

	return sum
}

func main() {
	state, err := parse(input)
	if err != nil {
		fmt.Printf("Unable to parse, %v\n", err)
		return
	}

	fmt.Printf("Answer 1: %d\n", answer1(state))
}
