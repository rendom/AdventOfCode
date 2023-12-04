package main

import (
	"testing"
)

const example string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestAnswerPart1(t *testing.T) {
	answer := 4361
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

func TestAnswerPart2(t *testing.T) {
	answer := 467835
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
