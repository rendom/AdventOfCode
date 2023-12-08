package main

import (
	"testing"
)

const example string = `Time:      7  15   30
Distance:  9  40  200`

func TestAnswerPart1(t *testing.T) {
	answer := 288
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

// 81+94
func TestAnswerPart2(t *testing.T) {
	answer := 71503
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
