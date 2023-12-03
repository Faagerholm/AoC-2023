package main

import "testing"

func TestDay3(t *testing.T) {
	part1, part2 := day3(read("input/day3-example.txt"))

	if part1 != 4361 {
		t.FailNow()
	}
	if part2 != 467835 {
		t.FailNow()
	}
}
