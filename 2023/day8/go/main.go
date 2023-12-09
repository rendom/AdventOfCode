package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var input string

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart1(input string) int {
	i, n := parseInput(input)
	return findGoal(i, findNodeByValue(n, "AAA"), "ZZZ")
}

func answerPart2(input string) int {
	i, n := parseInput(input)

	starts := []*Node{}
	for _, v := range n {
		if strings.HasSuffix(v.Value, "A") {
			starts = append(starts, v)
		}
	}

	return findGoalPart2(i, starts)
}

func findNodeByValue(nodes []*Node, value string) *Node {
	for _, n := range nodes {
		if n.Value == value {
			return n
		}
	}

	return nil
}

func findGoal(ins []string, start *Node, goal string) int {
	curr := start
	c := 0
	for goal != curr.Value {
		for _, i := range ins {
			c++
			if curr == nil {
				return -1
			}

			if i == "L" {
				curr = curr.Left
			} else {
				curr = curr.Right
			}
		}
	}

	return c
}

func findGoalPart2(ins []string, start []*Node) int {
	curr := start
	c := 0

	walk := true
	for walk {
		for _, i := range ins {
			c++
			for k, c := range curr {
				if c == nil {
					return -1
				}

				if i == "L" {
					curr[k] = c.Left
				} else {
					curr[k] = c.Right
				}
			}

			if hasNodesReachedDestination(curr) {
				walk = false
			}
		}
	}

	return c
}

func hasNodesReachedDestination(n []*Node) bool {
	for _, c := range n {
		if !strings.HasSuffix(c.Value, "Z") {
			return false
		}
	}
	return true
}

func parseInput(input string) (instructions []string, nodes []*Node) {
	s := strings.Split(input, "\n\n")
	instructions = strings.Split(s[0], "")

	m := map[string]*Node{}
	lr := map[string][]string{}

	for _, r := range strings.Split(s[1], "\n") {
		if r == "" {
			continue
		}

		s = strings.Split(r, " = (")

		n := &Node{}
		n.Value = s[0]
		nodes = append(nodes, n)

		m[n.Value] = n

		lrs := strings.Split(s[1], ", ")
		lr[n.Value] = []string{lrs[0], string(lrs[1][:len(lrs[1])-1])}
	}

	for k, n := range nodes {
		nodes[k].Left = m[lr[n.Value][0]]
		nodes[k].Right = m[lr[n.Value][1]]
	}

	return instructions, nodes
}
