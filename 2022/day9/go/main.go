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
	Body []*Pos
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

func n(n int) int {
	if n < 0 {
		return -1
	} else if n > 0 {
		return 1
	}

	return 0
}

func (s *State) Move(dir string, steps int) {
	tailIdx := len(s.Body) - 1
	for i := 1; i <= steps; i++ {
		s.Head.Move(dir, 1)
		follow := s.Head
		for idx, part := range s.Body {
			dx := follow.x - part.x
			dy := follow.y - part.y

			if math.Abs(float64(dx)) > 1 || math.Abs(float64(dy)) > 1 {
				part.x += n(dx)
				part.y += n(dy)
			}

			if idx == tailIdx {
				s.Visited[*part]++
			}

			follow = *part
		}
	}
}

func doStuff(data string, size int) (State, error) {
	body := []*Pos{}
	for i := 0; i < size; i++ {
		body = append(body, &Pos{0, 0})
	}

	s := State{
		Head:    Pos{0, 0},
		Body:    body,
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

func (s *State) calcVisited() int {
	sum := 0
	for _, v := range s.Visited {
		if v > 0 {
			sum++
		}
	}

	return sum
}

func answer1() {
	s, err := doStuff(input, 1)
	if err != nil {
		fmt.Printf("Unable to parse, %v\n", err)
		return
	}

	fmt.Printf("Answer 1: %d\n", s.calcVisited())
}

func answer2() {
	s, err := doStuff(input, 9)
	if err != nil {
		fmt.Printf("Unable to parse, %v\n", err)
		return
	}

	fmt.Printf("Answer 2: %d\n", s.calcVisited())
}

func main() {
	answer1()
	answer2()
}
