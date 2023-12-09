package main

import (
	"testing"
)
const exampleP1 string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const exampleP2 string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestAnswerPart1(t *testing.T) {
	answer := 142
	if r := answerPart1(exampleP1); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

func TestAnswerPart2(t *testing.T) {
	answer := 281
	if r := answerPart2(exampleP2); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
