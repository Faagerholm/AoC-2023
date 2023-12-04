package main

import (
	"fmt"
	"strings"
)

func day4(input []string) (int, int) {
	cards := make([][][]int, len(input))
	control := make([][]int, len(input))
	parse := func(s string) []int {
		nrs := strings.Split(s, " ")
		var numbers []int
		for _, n := range nrs {
			numbers = append(numbers, atoi(n))
		}
		return numbers
	}

	points := func(winning, nrs []int) int {
		var sum int
		for _, w := range winning {
			for _, n := range nrs {
				if w == n && w != 0 {
					if sum == 0 {
						sum = 1
					} else {
						sum *= 2
					}
				}
			}
		}
		return sum
	}

	// part1 - Count total points
	var tot int
	for idx, line := range input {
		scratch := strings.Split(line, " | ")
		card := parse(strings.Trim(scratch[1], " "))
		winning := parse(strings.Split(strings.Trim(scratch[0], " "), ":")[1])
		tot += points(card, winning)

		cards[idx] = append(cards[idx], card)
		control[idx] = winning
	}

	// part2 - Count total number of scratch cards
	var count int
	for idx := range input {
		for _, c := range cards[idx] {
			p := points(control[idx], c)
			next := 1
			for p >= 1 {
				cards[idx+next] = append(cards[idx+next], cards[idx+next][0])
				next++
				p /= 2
			}
			count++
		}
	}
	fmt.Printf("Day4 Part1 Solution: %d\n", tot)
	fmt.Printf("Day4 Part2 Solution: %d\n", count)
	return tot, count
}
