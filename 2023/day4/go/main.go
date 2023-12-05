package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

//go:embed data.txt
var input string

type Card struct {
	GameNr  int
	Correct []int
	Picks   []int
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart2(input string) int {
	cards, err := parseInput(input)
	if err != nil {
		return -1
	}

	for i := 0; i < len(cards); i++ {
		correctPicks := cards[i].getCorrectPicks()
		for y := 1; y <= len(correctPicks); y++ {
			card := getCardByNumber(cards[i].GameNr+y, cards)
			if card != nil {
				cards = append(cards, *card)
			}
		}
	}

	return len(cards)
}

func getCardByNumber(search int, cards []Card) *Card {
	for _, c := range cards {
		if c.GameNr == search {
			return &c
		}
	}

	return nil
}

func answerPart1(input string) int {
	cards, err := parseInput(input)
	if err != nil {
		return -1
	}

	sum := 0
	for _, c := range cards {
		sum += c.calculateCardPointsAnswer1()
	}

	return sum
}

func (c *Card) calculateCardPointsAnswer1() int {
	w := c.getCorrectPicks()
	sum := 0
	for i := 0; i < len(w); i++ {
		if i == 0 {
			sum = 1
		} else {
			sum *= 2
		}
	}

	return sum
}

func (c *Card) getCorrectPicks() []int {
	correctPicks := []int{}
	for _, v := range c.Picks {
		if slices.Contains(c.Correct, v) {
			correctPicks = append(correctPicks, v)
		}
	}

	return correctPicks
}

func parseInput(input string) ([]Card, error) {
	input = strings.TrimRight(input, "\n")
	cards := []Card{}
	for _, row := range strings.Split(input, "\n") {
		g, err := parseGame(row)
		if err != nil {
			return []Card{}, err
		}

		cards = append(cards, g)
	}

	return cards, nil
}

func parseGame(input string) (Card, error) {
	f := strings.Split(input, ": ")
	if len(f) != 2 {
		return Card{}, errors.Errorf("Invalid format on game")
	}

	header := strings.Split(f[0], " ")
	nr, err := strconv.Atoi(strings.Trim(header[len(header)-1], " "))
	if err != nil {
		return Card{}, errors.Errorf("Header err: %v", err)
	}

	c := Card{GameNr: nr}
	for k, v := range strings.Split(f[1], " | ") {
		for _, nrv := range strings.Split(v, " ") {
			if nrv == "" {
				continue
			}

			nr, err = strconv.Atoi(nrv)
			if err != nil {
				return Card{}, errors.Errorf("nr err: %v", err)
			}

			if k == 0 {
				c.Correct = append(c.Correct, nr)
			} else {
				c.Picks = append(c.Picks, nr)
			}
		}
	}

	return c, nil
}
