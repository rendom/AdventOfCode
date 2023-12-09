package main

import (
	_ "embed"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.txt
var input string

type Game struct {
	Joker bool
}

type Player struct {
	RawHand string
	Cards   []int
	Bid     int
	Bonus   int
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart1(input string) int {
	return calcAnswer(input, false)
}

func answerPart2(input string) int {
	return calcAnswer(input, true)
}

func calcAnswer(input string, joker bool) int {
	players := parseInput(input, joker)
	ranked := sortPlayersByRank(players)

	answer := 0
	for k, p := range ranked {
		answer += (len(ranked) - k) * p.Bid
	}
	return answer
}

func compareHands(a, b Player) int {
	for i := 0; i < 5; i++ {
		if a.Cards[i] < b.Cards[i] {
			return 1
		} else if a.Cards[i] > b.Cards[i] {
			return -1
		}
	}

	return 0
}

func sortPlayersByRank(players []Player) []Player {
	slices.SortStableFunc(players, func(a, b Player) int {
		if a.Bonus < b.Bonus {
			return 1
		} else if a.Bonus == b.Bonus {
			return compareHands(a, b)
		} else {
			return -1
		}
	})

	return players
}

func getCardWeights(joker bool) map[string]int {
	w := []string{
		"2", "3", "4", "5", "6", "7", "8",
		"9", "T", "J", "Q", "K", "A",
	}

	if joker {
		w = []string{
			"J", "2", "3", "4", "5", "6", "7", "8",
			"9", "T", "Q", "K", "A",
		}
	}

	m := map[string]int{}
	for i, v := range w {
		m[v] = i
	}

	return m
}

func countCards(cards []int) [][2]int {
	m := map[int]int{}

	for _, c := range cards {
		m[c]++
	}

	r := [][2]int{}
	for c, v := range m {
		r = append(r, [2]int{v, c})
	}

	sort.Slice(r, func(a, b int) bool {
		return r[a][0] > r[b][0]
	})

	return r
}

func getCountByValue(counts [][2]int, weight int) int {
	for _, c := range counts {
		if c[1] == weight {
			return c[0]
		}
	}

	return 0
}

func fillCardsWithJoker(c [][2]int) [][2]int {
	jokers := getCountByValue(c, 0)

	if jokers == 0 || jokers == 5 {
		return c
	}

	counts := [][2]int{}
	for _, v := range c {
		if v[1] == 0 {
			continue
		}

		counts = append(counts, [2]int{v[0] + jokers, v[1]})
		jokers = 0
	}
	return counts
}

func getCardsBonus(cards []int, joker bool) int {
	counts := countCards(cards)
	bonus := 0

	if joker {
		counts = fillCardsWithJoker(counts)
	}

	switch {
	case counts[0][0] == 5:
		bonus = 6
	case counts[0][0] == 4:
		bonus = 5
	case counts[0][0] == 3 && counts[1][0] == 2:
		bonus = 4
	case counts[0][0] == 3:
		bonus = 3
	case counts[0][0] == 2 && counts[1][0] == 2:
		bonus = 2
	case counts[0][0] == 2:
		bonus = 1
	}

	return bonus
}

func cardsToWeight(cards string, joker bool) []int {
	weights := getCardWeights(joker)
	r := []int{}

	for _, c := range cards {
		r = append(r, weights[string(c)])
	}

	return r
}

func parseInput(input string, joker bool) []Player {
	players := []Player{}
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}

		s := strings.Split(row, " ")

		player := Player{}
		player.RawHand = s[0]
		player.Cards = cardsToWeight(s[0], joker)
		player.Bid, _ = strconv.Atoi(s[1])
		player.Bonus = getCardsBonus(player.Cards, joker)
		players = append(players, player)
	}

	return players
}
