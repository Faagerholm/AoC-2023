package main

import (
	"fmt"
	"math"
	"strings"
)

type mapper struct {
	in, out, inc []int
}

func (m mapper) convert(in int) int {
	for i := range m.inc {
		if in >= m.in[i] && in <= m.in[i]+m.inc[i] {
			return in + m.out[i] - m.in[i]
		}
	}
	return in
}

func day5(input []string) (int, int) {
	seeds := atoia(strings.Split(input[0], ": ")[1])
	part1 := day5solver(seeds, input)

	var seedRange []int
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			seedRange = append(seedRange, j)
		}
	}
	part2 := day5solver(seedRange, input)
	return part1, part2
}

func day5solver(seeds []int, input []string) int {
	i := 3

	toMap := func(m *mapper) {
		for ; ; i++ {
			if i >= len(input) {
				return
			}
			converter := atoia(input[i])
			if len(converter) != 3 {
				i += 2
				return
			}
			m.in = append(m.in, converter[1])
			m.out = append(m.out, converter[0])
			m.inc = append(m.inc, converter[2])
		}
	}

	var seedToSoil mapper
	var soilToFert mapper
	var fertToWater mapper
	var waterToLight mapper
	var lightToTemp mapper
	var tempToHum mapper
	var humToLoc mapper
	for i < len(input) {
		toMap(&seedToSoil)
		toMap(&soilToFert)
		toMap(&fertToWater)
		toMap(&waterToLight)
		toMap(&lightToTemp)
		toMap(&tempToHum)
		toMap(&humToLoc)
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		seed = seedToSoil.convert(seed)
		seed = soilToFert.convert(seed)
		seed = fertToWater.convert(seed)
		seed = waterToLight.convert(seed)
		seed = lightToTemp.convert(seed)
		seed = tempToHum.convert(seed)
		seed = humToLoc.convert(seed)
		if seed < lowest {
			lowest = seed
		}
	}
	fmt.Printf("lowest location: %d\n", lowest)
	return lowest
}
