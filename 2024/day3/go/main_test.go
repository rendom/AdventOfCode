package main

import (
	"testing"
)

const example string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestAnswerPart1(t *testing.T) {
	answerPart1(example)
	answer := 161
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

func TestAnswerPart2(t *testing.T) {
	answer := 0
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
