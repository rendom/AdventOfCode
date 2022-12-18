package main

import (
	"fmt"
	"strconv"
	"strings"
)

func render(m M) {
	m[C{500, 0}] = SPAWNER
	start := getStart(m)
	end := getEnd(m)

	symbols := map[int]string{
		ROCK:    "#",
		SAND:    "o",
		AIR:     ".",
		SPAWNER: "+",
	}

	for y := start.y; y <= end.y; y++ {
		for x := start.x; x <= end.x; x++ {
			s := symbols[AIR]
			if material, ok := m[C{x, y}]; ok {
				s = symbols[material]
			}
			fmt.Print(s)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func parseData(d string) (M, error) {
	m := M{}
	lines := [][]C{}
	d = strings.TrimRight(d, "\n")
	for _, r := range strings.Split(d, "\n") {
		line := []C{}
		for _, lp := range strings.Split(r, " -> ") {
			values := strings.Split(lp, ",")
			if len(values) < 2 {
				return m, fmt.Errorf("Unable to parse rocks, invalid value")
			}

			x, err := strconv.Atoi(values[0])
			if err != nil {
				return m, fmt.Errorf("Unable to parse rocks, invalid format on x-axis")
			}

			y, err := strconv.Atoi(values[1])
			if err != nil {
				return m, fmt.Errorf("Unable to parse rocks, invalid format on y-axis")
			}
			line = append(line, C{x, y})
		}
		lines = append(lines, line)
	}

	return placeMaterialFromLines(m, lines, ROCK), nil
}

func fillRange(m M, a C, b C, material int) M {
	if a.x == b.x {
		from := min(a.y, b.y)
		to := max(a.y, b.y)

		for y := from; y <= to; y++ {
			m[C{a.x, y}] = material
		}
	} else {
		from := min(a.x, b.x)
		to := max(a.x, b.x)

		for x := from; x <= to; x++ {
			m[C{x, a.y}] = material
		}
	}

	return m

}

func placeMaterialFromLines(m M, lines [][]C, material int) M {
	for _, r := range lines {
		l := len(r)
		for i := 0; i < l; i++ {
			if i+1 >= l {
				continue
			}

			m = fillRange(m, r[i], r[i+1], material)
		}
	}

	return m
}
