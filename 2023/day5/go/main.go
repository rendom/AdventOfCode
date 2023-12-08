package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

//go:embed data.txt
var input string

type Seeds []Seed

type Ranges []Range
type Range [2]int // Start, length

type DestinationMap struct {
	DestRangeStart   int
	SourceRangeStart int
	RangeLength      int
}

type Seed struct {
	SeedID          int
	SeedRangeLength int

	Soil        Ranges
	Fertilizer  Ranges
	Water       Ranges
	Light       Ranges
	Temperature Ranges
	Humidity    Ranges
	Location    Ranges
}

type Puzzle struct {
	Seeds                 Seeds
	SeedToSoil            []DestinationMap
	SoilToFertilizer      []DestinationMap
	FertilizerToWater     []DestinationMap
	WaterToLight          []DestinationMap
	LightToTemperature    []DestinationMap
	TemperatureToHumidity []DestinationMap
	HumidityToLocation    []DestinationMap
}

func main() {
	fmt.Printf("Answer part 1: %d\n", answerPart1(input))
	fmt.Printf("Answer part 2: %d\n", answerPart2(input))
}

func answerPart2(input string) int {
	p := parseInput(input, true)
	p.fillAllSeeds()
	return minLocationFromSeeds(p.Seeds)
}

func answerPart1(input string) int {
	p := parseInput(input, false)
	p.fillAllSeeds()
	return minLocationFromSeeds(p.Seeds)
}

func minLocationFromSeeds(s Seeds) int {
	answer := math.MaxInt
	for _, v := range s {
		for _, l := range v.Location {
			answer = min(l[0], answer)
		}
	}
	return answer
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func (p *Puzzle) fillAllSeeds() {
	l := len(p.Seeds)
	for i := 0; i < l; i++ {
		p.Seeds[i].fillSeedValues(*p)
	}
}

// 37405178 är tydligen fel
func (s *Seed) fillSeedValues(p Puzzle) {
	s.Soil = resolveChain(Ranges{Range{s.SeedID, s.SeedRangeLength}}, p.SeedToSoil)
	s.Fertilizer = resolveChain(s.Soil, p.SoilToFertilizer)
	s.Water = resolveChain(s.Fertilizer, p.FertilizerToWater)
	s.Light = resolveChain(s.Water, p.WaterToLight)
	s.Temperature = resolveChain(s.Light, p.LightToTemperature)
	s.Humidity = resolveChain(s.Temperature, p.TemperatureToHumidity)
	s.Location = resolveChain(s.Humidity, p.HumidityToLocation)
}

func resolveChain(ir Ranges, dms []DestinationMap) Ranges {
	destResult := Ranges{}

	checkRanges := ir

	for _, dm := range dms {
		leftOvers := Ranges{}
		for _, r := range checkRanges {
			destination, lo, status := resolveDm(r, dm)

			if len(lo) > 0 {
				leftOvers = append(leftOvers, lo...)
			}

			if status == -1 {
				continue
			}

			destResult = append(destResult, destination)
		}

		checkRanges = leftOvers
	}

	if len(checkRanges) > 0 {
		destResult = append(destResult, checkRanges...)
	}

	if len(destResult) == 0 {
		return ir
	}

	return destResult
}

func resolveDm(ir Range, dm DestinationMap) (Range, Ranges, int) {
	leftOver := Ranges{}
	rangeEnd := ir[0] + ir[1]

	dmSourceEnd := dm.SourceRangeStart + dm.RangeLength
	start, end := overlap(ir[0], rangeEnd, dm.SourceRangeStart, dmSourceEnd)

	if start == -1 && end == -1 {
		return ir, Ranges{ir}, -1
	}

	if ir[0] < start {
		leftOver = append(leftOver, Range{ir[0], (dm.SourceRangeStart - ir[0] - 1)})
	}

	if rangeEnd > dmSourceEnd {
		ns := dmSourceEnd + 1
		nln := (ir[0] + ir[1]) - dmSourceEnd - 1 // - ir[1]
		leftOver = append(leftOver, Range{ns, nln})
	}

	diffStart := start - dm.SourceRangeStart
	destinationID := dm.DestRangeStart + diffStart
	return Range{destinationID, (end - start)}, leftOver, 1
}

func overlap(aS int, aE int, bS int, bE int) (int, int) {
	if aS <= bE && aE >= bS {
		return max(aS, bS), min(aE, bE)
	}

	return -1, -1
}

func parseInput(input string, seedRange bool) Puzzle {
	p := Puzzle{}
	buffer := []string{}
	input = input + "\n"
	for k, line := range strings.Split(input, "\n") {
		if k == 0 {
			seeds := Seeds{}
			var err error
			if seedRange {
				seeds, err = parseRangeSeeds(line)
			} else {
				seeds, err = parseSeeds(line)
			}

			if err != nil {
				panic(err)
			}

			p.Seeds = seeds
			continue
		}

		if line == "" {
			if len(buffer) == 0 {
				continue
			}

			name, dm, err := parseMap(buffer)
			if err != nil {
				panic(err)
			}

			switch name {
			case "seed-to-soil":
				p.SeedToSoil = dm
			case "soil-to-fertilizer":
				p.SoilToFertilizer = dm
			case "fertilizer-to-water":
				p.FertilizerToWater = dm
			case "water-to-light":
				p.WaterToLight = dm
			case "light-to-temperature":
				p.LightToTemperature = dm
			case "temperature-to-humidity":
				p.TemperatureToHumidity = dm
			case "humidity-to-location":
				p.HumidityToLocation = dm
			}

			buffer = []string{}
			continue
		}

		buffer = append(buffer, line)
	}

	return p
}

func parseMap(rows []string) (string, []DestinationMap, error) {
	if len(rows) < 1 {
		return "", []DestinationMap{}, errors.Errorf("parseMap, empty?")
	}

	headerSplit := strings.Split(rows[0], " ")
	if len(headerSplit) < 2 {
		return "", []DestinationMap{}, errors.Errorf("parseMap, invalid header row")
	}

	dma := []DestinationMap{}
	for r, row := range rows[1:] {
		dm := DestinationMap{}
		for k, value := range strings.Split(row, " ") {
			nr, err := strconv.Atoi(value)
			if err != nil {
				return "", []DestinationMap{}, errors.Errorf("parseMap, unable to parse range values, map: %s, row: %d", headerSplit[0], r)
			}

			switch k {
			case 0:
				dm.DestRangeStart = nr
			case 1:
				dm.SourceRangeStart = nr
			case 2:
				dm.RangeLength = nr - 1
			}
		}

		dma = append(dma, dm)
	}

	return headerSplit[0], dma, nil
}

func parseRangeSeeds(input string) (Seeds, error) {
	rawSeeds := strings.Split(input, " ")[1:]
	s := Seeds{}
	seedLen := len(rawSeeds)
	for i := 0; i < seedLen; i += 2 {
		start, err := strconv.Atoi(rawSeeds[i])
		if err != nil {
			return Seeds{}, err
		}

		length, err := strconv.Atoi(rawSeeds[i+1])
		if err != nil {
			return Seeds{}, err
		}

		s = append(s, Seed{SeedID: start, SeedRangeLength: length - 1})
	}
	return s, nil
}

func parseSeeds(input string) (Seeds, error) {
	rawSeeds := strings.Split(input, " ")[1:]
	s := Seeds{}
	for _, v := range rawSeeds {
		i, err := strconv.Atoi(v)
		if err != nil {
			return Seeds{}, err
		}

		s = append(s, Seed{SeedID: i, SeedRangeLength: 0})
	}
	return s, nil
}
