package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed data-real.txt
var input string

const CLEAR = "\033[H\033[2J"

type State struct {
	Columns [][]byte
}

type Instruction struct {
	FromColumn   int
	FromPosition int
	ToColumn     int
}

func (s *State) move(i Instruction, keepOrder bool) {
	fromColumn := append([]byte{}, s.Columns[i.FromColumn]...)
	move := fromColumn[:i.FromPosition+1]
	if !keepOrder {
		sort.SliceStable(move, func(a, b int) bool {
			return a > b
		})
	}

	s.Columns[i.FromColumn] = s.Columns[i.FromColumn][i.FromPosition+1:]
	s.Columns[i.ToColumn] = append(move, s.Columns[i.ToColumn]...)
}

func (s *State) rows() int {
	max := 0
	for _, c := range s.Columns {
		if l := len(c); l > max {
			max = l
		}
	}

	return max
}

func (s *State) printState() {
	rows := s.rows()
	for i := 0; i < rows; i++ {
		for _, c := range s.Columns {
			offset := (rows - len(c))
			key := i - offset
			if len(c) > key && key >= 0 {
				fmt.Printf("[%s] ", string(c[i-offset]))
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	// state.printState()
	doStuff(false)
	fmt.Println("\n")
	fmt.Println("\n")
	doStuff(true)
}

func doStuff(keepOrder bool) {
	state, instructions, err := parseFile(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, i := range instructions {
		//fmt.Printf("%s\n", CLEAR)
		state.move(i, keepOrder)
		//state.printState()
		//time.Sleep(time.Duration(time.Millisecond * 5))
	}

	answer := ""
	for _, v := range state.Columns {
		answer = answer + string(v[0])
	}
	state.printState()
	fmt.Printf("\nAnswer: %s\n", answer)
}

func parseFile(data string) (State, []Instruction, error) {
	sections := strings.Split(data, "\n\n")
	state := parseState(sections[0])
	instructions, err := parseInstructions(sections[1])
	if err != nil {
		return state, instructions, err
	}

	return state, instructions, nil
}

func parseState(data string) State {
	rows := strings.Split(data, "\n")
	rows = rows[:len(rows)-1]
	columns := len(rows[0])/4 + 1

	state := State{}
	state.Columns = make([][]byte, columns)

	for k, _ := range state.Columns {
		for _, row := range rows {
			valuePosition := (k * 4) + 1
			if string(row[valuePosition]) != " " {
				state.Columns[k] = append(state.Columns[k], row[valuePosition])
			}
		}
	}
	return state
}

func parseInstruction(data string) (Instruction, error) {
	c := strings.Split(data, " ")
	if len(c) != 6 {
		return Instruction{}, fmt.Errorf("Invalid instruction row")
	}

	fromPosition, err := strconv.Atoi(c[1])
	if err != nil {
		return Instruction{}, fmt.Errorf("fromPosition %v", err)
	}

	fromColumn, err := strconv.Atoi(c[3])
	if err != nil {
		return Instruction{}, fmt.Errorf("fromColumn %v", err)
	}

	toColumn, err := strconv.Atoi(c[5])
	if err != nil {
		return Instruction{}, fmt.Errorf("toColumn %v", err)
	}

	return Instruction{
		FromColumn:   fromColumn - 1,
		FromPosition: fromPosition - 1,
		ToColumn:     toColumn - 1,
	}, nil
}

func parseInstructions(data string) ([]Instruction, error) {
	rows := strings.Split(data, "\n")
	instructions := []Instruction{}
	for k, r := range rows {
		if r == "" {
			continue
		}

		instruction, err := parseInstruction(r)
		if err != nil {
			return instructions, fmt.Errorf("Error at row: %d, %v", k+1, err)
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil

}
