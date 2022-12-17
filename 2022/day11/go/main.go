package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed data-real.txt
var input string

type Monkey struct {
	monkeyNr  int
	items     []int
	operation [3]string

	test     int
	testPass int
	testFail int
}

func getOperationValue(o string, item int) (int, error) {
	if o == "old" {
		return item, nil
	}

	i, err := strconv.Atoi(o)
	return i, err
}

func (m *Monkey) doOperation(v int) (int, error) {
	val1, err := getOperationValue(m.operation[0], v)
	if err != nil {
		return 0, err
	}

	val2, err := getOperationValue(m.operation[2], v)
	if err != nil {
		return 0, err
	}

	switch m.operation[1] {
	case "*":
		return val1 * val2, nil
	case "+":
		return val1 + val2, nil
	}

	return 0, fmt.Errorf("invalid operation")
}

func (m *Monkey) runTest(v int) int {
	if v%m.test == 0 {
		return m.testPass
	} else {
		return m.testFail
	}
}

func parseMonkey(rows []string) (Monkey, error) {
	m := Monkey{}
	var err error
	m.monkeyNr, err = strconv.Atoi(string(rows[0][7]))
	if err != nil {
		return m, fmt.Errorf("Unable to parse monkey nr, %v", err)
	}

	m.items, err = parseItems(rows[1][18:])
	if err != nil {
		return m, err
	}

	m.operation, err = parseOperation(rows[2])
	if err != nil {
		return m, err
	}

	m.test, err = strconv.Atoi(rows[3][21:])
	if err != nil {
		return m, fmt.Errorf("Unable to parse test value, %v", err)
	}

	m.testPass, err = strconv.Atoi(rows[4][29:])
	if err != nil {
		return m, fmt.Errorf("Unable to parse test pass value, %v", err)
	}

	m.testFail, err = strconv.Atoi(rows[5][30:])
	if err != nil {
		return m, fmt.Errorf("Unable to parse test fail value, %v", err)
	}

	return m, nil
}

func parseItems(s string) ([]int, error) {
	items := strings.Split(s, ", ")
	r := []int{}
	for idx, v := range items {
		i, err := strconv.Atoi(v)
		if err != nil {
			return r, fmt.Errorf("Unable to parse items (idx: %d), %v", idx, err)
		}

		r = append(r, i)
	}

	return r, nil
}

func parseOperation(s string) ([3]string, error) {
	re := regexp.MustCompile(`Operation: new = ([a-z0-9]+) (\*|\+) ([a-z0-9]+)`)
	m := re.FindStringSubmatch(s)
	if len(m) == 0 {
		return [3]string{}, fmt.Errorf("Unable to parse operation")
	}

	return [3]string{m[1], m[2], m[3]}, nil
}

// func parseOperation(s string) (func(v int) int, error) {
// 	re := regexp.MustCompile(`Operation: new = ([a-z0-9]+) (\*|\+) ([a-z0-9]+)`)
// 	m := re.FindStringSubmatch(s)
// 	if len(m) == 0 {
// 		return func(v int) int {
// 			return 0
// 		}, fmt.Errorf("Unable to parse operation")
// 	}
//
// 	if m[0] == "old" {
// 	}
//
// 	return func(v int) int {
// 		return 0
// 	}, nil
// }

func parseData(data string) ([]Monkey, error) {
	rows := strings.Split(data, "\n")
	r := []Monkey{}
	for i := 0; i < len(rows); i += 7 {
		m, err := parseMonkey(rows[i : i+7])
		if err != nil {
			return r, err
		}

		r = append(r, m)
	}

	return r, nil
}

func run(monkeys []Monkey) {
	inspected := make([]int, len(monkeys))

	for i := 0; i < 20; i++ {
		for k, m := range monkeys {
			for len(monkeys[k].items) > 0 {
				v, err := m.doOperation(monkeys[k].items[0])
				if err != nil {
					panic(err)
				}

				inspected[k]++

				monkeys[k].items = monkeys[k].items[1:]

				item := (v / 3)
				throwTo := m.runTest(item)
				monkeys[throwTo].items = append(monkeys[throwTo].items, item)
			}
		}
	}

	for k, v := range monkeys {
		fmt.Printf("Monkey %d, Inspected %d, items: %v\n", v.monkeyNr, inspected[k], v.items)
	}

	sort.SliceStable(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})

	fmt.Println("Answer", inspected[0]*inspected[1])
}

func run2(monkeys []Monkey) {
	inspected := make([]int, len(monkeys))

	mmm := 1
	for _, v := range monkeys {
		mmm *= v.test
	}

	for i := 0; i < 10000; i++ {
		for k, m := range monkeys {
			for len(monkeys[k].items) > 0 {
				v, err := m.doOperation(monkeys[k].items[0])
				if err != nil {
					panic(err)
				}

				inspected[k]++

				monkeys[k].items = monkeys[k].items[1:]

				item := v % mmm
				throwTo := m.runTest(item)
				monkeys[throwTo].items = append(monkeys[throwTo].items, item)
			}

		}
	}

	sort.SliceStable(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})
	fmt.Printf("Answer %d (should be %d)\n", inspected[0]*inspected[1], 2713310158)
}
func main() {
	r, _ := parseData(input)
	run(r)

	b, err := parseData(input)
	if err != nil {
		panic(err)
	}
	run2(b)
}
