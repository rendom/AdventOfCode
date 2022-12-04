package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	ROCK    = 1
	PAPER   = 2
	SCISSOR = 3

	WIN  = 6
	DRAW = 3
	LOSS = 0
)

var ACTION = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSOR,
	"Y": PAPER,
	"X": ROCK,
	"Z": SCISSOR,
}

//go:embed data.txt
var input string

type Match [2]int

func main() {
	part1()
	part2()
}

func part1() {
	matches, err := parseMatches(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for _, m := range matches {
		sum += calcPoints(m)
	}

	fmt.Printf("Answer 1: %d\n", sum)
}

func part2() {
	matches, err := parseMatches(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Y means you need to end the round in a draw,
	// X means you need to lose,
	// Z means you need to win. Good luck!"

	drawKey := PAPER
	lossKey := ROCK
	winKey := SCISSOR

	m := map[int]map[int]int{
		ROCK: map[int]int{
			winKey:  PAPER,
			lossKey: SCISSOR,
			drawKey: ROCK,
		},
		PAPER: map[int]int{
			winKey:  SCISSOR,
			lossKey: ROCK,
			drawKey: PAPER,
		},
		SCISSOR: map[int]int{
			winKey:  ROCK,
			lossKey: PAPER,
			drawKey: SCISSOR,
		},
	}

	for k, match := range matches {
		matches[k][1] = m[match[0]][match[1]]
	}

	sum := 0
	for _, m := range matches {
		sum += calcPoints(m)
	}

	fmt.Printf("Answer 2: %d\n", sum)
}

func matchOutcome(m Match) int {
	if m[1] == m[0] {
		return DRAW
	}

	if (m[1] == ROCK && m[0] == SCISSOR) ||
		(m[1] == PAPER && m[0] == ROCK) ||
		(m[1] == SCISSOR && m[0] == PAPER) {
		return WIN
	}

	return LOSS
}

func calcPoints(m Match) int {
	return m[1] + matchOutcome(m)
}

func parseMatches(matchData string) ([]Match, error) {
	matchData = strings.TrimRight(matchData, "\n")
	rounds := strings.Split(matchData, "\n")
	matches := []Match{}
	for r, c := range rounds {
		competitors := strings.Split(c, " ")
		var comp [2]int
		for k, v := range competitors {
			if _, ok := ACTION[v]; !ok {
				return matches, fmt.Errorf(
					"Invalid action (%s) in round %d",
					v,
					r,
				)
			}
			comp[k] = ACTION[v]
		}

		matches = append(matches, comp)
	}

	return matches, nil
}
