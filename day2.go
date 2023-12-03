package main

import (
	"fmt"
	"strings"
)

func day2(lines []string) (int, int) {
	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var sumOfLegalGames int
	var sumOfHighest int
	for id, line := range lines {
		games := strings.Split(line, ":")[1]
		legalGame := true
		highest := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range strings.Split(games, ";") {
			for _, gem := range strings.Split(game, ",") {
				p := strings.Split(strings.TrimSpace(gem), " ")
				c := atoi(p[0])
				if max[p[1]] < c {
					legalGame = false
				}
				if highest[p[1]] < c {
					highest[p[1]] = c
				}
			}
		}
		if legalGame {
			sumOfLegalGames += id + 1
		}
		sumOfHighest += highest["red"] * highest["blue"] * highest["green"]
	}
	fmt.Printf("Day2 - part1 Solution: %d\n", sumOfLegalGames)
	fmt.Printf("Day2 - part2 Solution: %d\n", sumOfHighest)

	return sumOfLegalGames, sumOfHighest
}
