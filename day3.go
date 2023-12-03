package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func day3(input []string) (int, int) {
	re := regexp.MustCompile(`\d+`)

	padding := len(input[0])

	symbols := make(map[int]rune)
	numbers := make(map[int]int)
	for y, line := range input {
		for x, c := range line {
			if unicode.IsDigit(c) || c == '.' {
				continue
			}
			symbols[y*padding+x] = c
		}
	}

	hasSymbolsAdjecent := func(lineNr []int, y int) bool {
		for x := lineNr[0]; x < lineNr[1]; x++ {
			for _, idx := range []int{
				(y-1)*padding + x - 1,
				(y-1)*padding + x,
				(y-1)*padding + x + 1,
				y*padding + x - 1,
				y*padding + x + 1,
				(y+1)*padding + x - 1,
				(y+1)*padding + x,
				(y+1)*padding + x + 1,
			} {
				if _, ok := symbols[idx]; ok {
					return true
				}
			}
		}
		return false
	}
	var partNrs []int
	for y, line := range input {
		idx := re.FindAllStringIndex(line, len(line))
		for _, partNrIdx := range idx {
			if hasSymbolsAdjecent(partNrIdx, y) {
				for x := partNrIdx[0]; x < partNrIdx[1]; x++ {
					numbers[y*padding+x] = atoi(line[partNrIdx[0]:partNrIdx[1]])
				}
				partNrs = append(partNrs, atoi(line[partNrIdx[0]:partNrIdx[1]]))
			}
		}
	}
	fmt.Printf("Day3 part1 Solution: %d\n", sum(partNrs))

	// part2

	getAdjecentParts := func(idx int) []int {
		parts := make(map[int]struct{})
		for _, i := range []int{
			idx - padding - 1,
			idx - padding,
			idx - padding + 1,
			idx - 1,
			idx + 1,
			idx + padding - 1,
			idx + padding,
			idx + padding + 1,
		} {
			if nr, ok := numbers[i]; ok {
				parts[nr] = struct{}{}
			}
		}
		var unique []int
		for p := range parts {
			unique = append(unique, p)
		}
		return unique
	}
	var gears []int
	for idx, symbol := range symbols {
		if symbol == '*' {
			parts := getAdjecentParts(idx)
			if len(parts) == 2 {
				gears = append(gears, parts[0]*parts[1])
			}
		}
	}

	fmt.Printf("Day3 part2 Solution: %d\n", sum(gears))

	return sum(partNrs), sum(gears)
}
