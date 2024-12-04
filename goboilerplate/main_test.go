package main

import (
	"testing"
)

const example string = `
	0000
`

func TestAnswerPart1(t *testing.T) {
	answerPart1(example)
	answer := 1
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

func TestAnswerPart2(t *testing.T) {
	answer := 2
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
