package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

type Game struct {
	ID   int
	Sets []Bags
}

type Bags struct {
	Red   int
	Green int
	Blue  int
}

const (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

// only 12 red cubes, 13 green cubes, and 14 blue cubes
// max ^

func main() {
	answer1 := answerPart1(input)
	// 2456 to low
	fmt.Printf("Answer part 1: %d\n", answer1)
}

func answerPart1(input string) int {
	games, err := parseInput(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}

	sum := sumGameIds(getValidGamesForPart1(games))
	return sum
}

func sumGameIds(games []Game) int {
	sum := 0
	for _, v := range games {
		sum += v.ID
	}
	return sum
}

func getValidGamesForPart1(games []Game) []Game {
	result := []Game{}
	for _, g := range games {
		isOk := true
		for _, s := range g.Sets {
			if s.Red > MAX_RED || s.Blue > MAX_BLUE || s.Green > MAX_GREEN {
				isOk = false
				break
			}
		}

		if isOk {
			result = append(result, g)
		}
	}

	return result
}

func parseInput(input string) ([]Game, error) {
	input = strings.TrimRight(input, "\n")
	games := []Game{}
	for _, row := range strings.Split(input, "\n") {
		g, err := parseLineToGame(row)
		if err != nil {
			return []Game{}, err
		}

		games = append(games, g)
	}

	return games, nil
}

func parseLineToGame(row string) (Game, error) {
	s1 := strings.Split(row, ": ")
	if len(s1) < 2 {
		return Game{}, errors.New("Invalid row, unable to parse game header")
	}

	g := Game{}
	s2 := strings.Split(s1[0], " ")
	if len(s2) < 2 {
		return Game{}, errors.New("Invalid row, unable to parse ID")
	}

	id, err := strconv.Atoi(s2[1])
	if err != nil {
		return Game{}, err
	}

	g.ID = id
	g.Sets, err = parseBags(s1[1])
	if err != nil {
		return Game{}, err
	}
	return g, nil
}

func parseBags(row string) ([]Bags, error) {
	result := []Bags{}
	sets := strings.Split(row, "; ")
	for _, set := range sets {
		r := Bags{}
		bags := strings.Split(set, ", ")
		for _, bag := range bags {
			s := strings.Split(bag, " ")
			amount, err := strconv.Atoi(s[0])
			if err != nil {
				return []Bags{}, err
			}

			switch s[1] {
			case "red":
				r.Red = amount
			case "green":
				r.Green = amount
			case "blue":
				r.Blue = amount
			}
		}

		result = append(result, r)
	}

	return result, nil
}
