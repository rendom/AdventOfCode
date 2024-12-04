package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.txt
var input string

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart1(input string) int {
	return 0
}

func answerPart2(input string) int {
	return 0
}
