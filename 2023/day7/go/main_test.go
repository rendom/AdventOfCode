package main

import (
	"testing"
)

const example string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestAnswerPart1(t *testing.T) {
	answer := 6440
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

func TestAnswerPart2(t *testing.T) {
	answer := 5905
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
