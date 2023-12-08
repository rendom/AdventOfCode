package main

import (
	"testing"
)

// [[9999 4] [19999 4] [1005 4]]
// const example string = `seeds: 79 14 55 13 1000 10
const example string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

// DESTSTART | SOURCESTART | LENGTH

// Humi 46 10
// Check 56 93
// NO match 56 93
// matched: 56 - 56 | og: 56 - 93
// loc 60 0

func TestAnswerPart1(t *testing.T) {
	answer := 35
	if r := answerPart1(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}

// 81+94
func TestAnswerPart2(t *testing.T) {
	answer := 46
	if r := answerPart2(example); r != answer {
		t.Fatalf("Wrong answer, got %d expected %d", r, answer)
	}
}
