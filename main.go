package main

import (
	"bufio"
	"os"
	"strconv"
)

func read(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func sum(in []int) int {
	var s int
	for _, i := range in {
		s += i
	}
	return s
}

func main() {
	day4(read("input/day4.txt"))
}
