package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.txt
var input string

func getAnswer(data string, check int) (int, error) {
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		until := i + check
		if until > dataLen {
			return 0, fmt.Errorf("Unable to find answer")
		}

		if unique(data[i:until]) {
			return until, nil
		}
	}

	return 0, fmt.Errorf("Unable to find answer")
}

func unique(data string) bool {
	exists := map[rune]bool{}
	for _, c := range data {
		if _, ok := exists[c]; ok {
			return false
		}

		exists[c] = true
	}
	return true
}

func main() {
	answer1, _ := getAnswer(input, 4)
	answer2, _ := getAnswer(input, 14)
	fmt.Printf("Answer 1: %d\nAnswer 2: %d\n", answer1, answer2)
}
