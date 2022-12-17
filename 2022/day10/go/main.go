package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data-real.txt
var input string

type Instruction struct {
	cmd  string
	data int
}
type frame [6][40]string
type State struct {
	x           int
	framebuffer frame
}

func (s *State) runInstructions(ins []Instruction) []int {
	cycles := 0
	instructionCycle := 0
	signalStr := []int{}

	for k, r := range s.framebuffer {
		for p, _ := range r {
			s.framebuffer[k][p] = "."
		}
	}

	var currentInstruction Instruction
	for len(ins) > 0 {
		cycles++
		instructionCycle++

		s.writePixelToFramebuffer(cycles)

		if (cycles-20)%40 == 0 {
			signalStr = append(signalStr, s.x*cycles)
		}

		currentInstruction = ins[0]
		switch currentInstruction.cmd {
		case "noop":
			if instructionCycle == 1 {
				ins = ins[1:]
				instructionCycle = 0
			}
		case "addx":
			if instructionCycle == 2 {
				s.x += currentInstruction.data
				ins = ins[1:]
				instructionCycle = 0
			}
		}

	}

	return signalStr
}

func (s *State) writePixelToFramebuffer(cycle int) {
	writePixel := (cycle - 1) % 40
	if writePixel+1 >= s.x && writePixel+1 <= s.x+2 {
		s.framebuffer[int(cycle/40)][writePixel] = "#"
	}
}

func parseInstructions(data string) ([]Instruction, error) {
	ins := []Instruction{}
	for _, r := range strings.Split(data, "\n") {
		if r == "" {
			continue
		}

		s := strings.Split(r, " ")
		data := 0
		if len(s) > 1 {
			var err error
			data, err = strconv.Atoi(s[1])
			if err != nil {
				return []Instruction{}, err
			}
		}

		ins = append(ins, Instruction{s[0], data})
	}

	return ins, nil
}

func render(s frame) {
	for _, v := range s {
		for _, s := range v {
			fmt.Print(s)
		}
		fmt.Print("\n")
	}
}

func main() {
	ins, err := parseInstructions(input)
	if err != nil {
		panic(err)
	}

	s := State{x: 1}
	sum := 0
	for _, v := range s.runInstructions(ins) {
		sum += v
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:")
	render(s.framebuffer)

}
