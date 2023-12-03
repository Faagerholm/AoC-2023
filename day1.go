package main

import (
	"fmt"
	"unicode"
)

func day1(input []string) int {
	var sum int
	for _, line := range input {
		var digits []int
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, atoi(string(char)))
			}
		}
		if len(digits) > 0 {
			sum += digits[0]*10 + digits[len(digits)-1]
		}
	}
	fmt.Printf("Day1 - part1 solution: %d\n", sum)

	sum = 0
	alfaNumberic := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range input {
		var digits []int
		for pos, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, atoi(string(char)))
			}
			for i, d := range alfaNumberic {
				if len(d) <= len(line)-pos && line[pos:pos+len(d)] == d {
					digits = append(digits, i+1)
				}
			}
		}
		if len(digits) > 1 {
			sum += digits[0]*10 + digits[len(digits)-1]
		}
	}
	fmt.Printf("Day1 - part2 solution: %d\n", sum)

	return sum
}
