package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	part1 := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	t.Log("With part 1 example")
	solution1 := day1(part1)

	if solution1 != 142 {
		t.FailNow()
	}

	part2 := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	t.Log("With part 2 example")
	solution2 := day1(part2)
	if solution2 != 65 {
		t.FailNow()
	}
}
