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

type DestinationMap struct {
	DestRangeStart   int
	SourceRangeStart int
	RangeLength      int
}

type Seed struct {
	SeedID      int
	Soil        *int
	Fertilizer  *int
	Water       *int
	Light       *int
	Temperature *int
	Humidity    *int
	Location    *int
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
	return 0
}

func answerPart1(input string) int {
	p := parseInput(input)
	p.fillAllSeeds()

	answer := math.MaxInt
	for _, v := range p.Seeds {
		if v.Location != nil {
			answer = min(*v.Location, answer)
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

func (p *Puzzle) fillAllSeeds() {
	l := len(p.Seeds)
	for i := 0; i < l; i++ {
		p.Seeds[i].fillSeedValues(*p)
	}
}

func (s *Seed) fillSeedValues(p Puzzle) {
	s.Soil = resolveChain(&s.SeedID, p.SeedToSoil)
	s.Fertilizer = resolveChain(s.Soil, p.SoilToFertilizer)
	s.Water = resolveChain(s.Fertilizer, p.FertilizerToWater)
	s.Light = resolveChain(s.Water, p.WaterToLight)
	s.Temperature = resolveChain(s.Light, p.LightToTemperature)
	s.Humidity = resolveChain(s.Temperature, p.TemperatureToHumidity)
	s.Location = resolveChain(s.Humidity, p.HumidityToLocation)
}

func resolveChain(sourceID *int, dms []DestinationMap) *int {
	if sourceID == nil {
		return nil
	}

	for _, dm := range dms {
		sourceEnd := dm.SourceRangeStart + dm.RangeLength
		if *sourceID < dm.SourceRangeStart || *sourceID > sourceEnd {
			continue
		}

		offset := (dm.SourceRangeStart + dm.RangeLength) - *sourceID
		destinationID := (dm.DestRangeStart + dm.RangeLength) - offset

		return &destinationID
	}

	return sourceID
}

func parseInput(input string) Puzzle {
	p := Puzzle{}
	buffer := []string{}
	input = input + "\n"
	for k, line := range strings.Split(input, "\n") {
		if k == 0 { // strings.Contains("seeds:", line)
			seeds, err := parseSeeds(line)
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
				dm.RangeLength = nr
			}
		}

		dma = append(dma, dm)
	}

	return headerSplit[0], dma, nil
}

func parseSeeds(input string) (Seeds, error) {
	rawSeeds := strings.Split(input, " ")[1:]
	s := Seeds{}
	for _, v := range rawSeeds {
		i, err := strconv.Atoi(v)
		if err != nil {
			return Seeds{}, err
		}

		s = append(s, Seed{SeedID: i})
	}
	return s, nil
}
