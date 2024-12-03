package main

import (
	"testing"
)

const example string = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestAnswerPart1(t *testing.T) {
	answerPart1(example)
	answer := 11
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

func TestAnswerPart2(t *testing.T) {
	answer := 31
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
