package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var testData = `Time:      7  15   30
Distance:  9  40  200`

var data = `Time:        48     98     90     83
Distance:   390   1103   1112   1360`

func main() {
	rows := strings.Split(data, "\n")
	re := regexp.MustCompile(`(\d+)`)
	times := re.FindAllString(rows[0], 4)
	distances := re.FindAllString(rows[1], 4)

	sum := 1
	for idx, t := range times {
		time, _ := strconv.Atoi(t)
		distance, _ := strconv.Atoi(distances[idx])
		var count int
		for i := 0; i < time; i++ {
			d := i * (time - i)
			if d > distance {
				count++
			}
		}
		sum *= count
	}
	fmt.Printf("Part 1 Solution: %d\n", sum)

	// part2
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(rows[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(rows[1], ":")[1], " ", ""))

	var i, d int
	for ; i < time; i++ {
		d = i * (time - i)
		if d > distance {
			break
		}
	}

	fmt.Printf("Part 2 solution: %d\n", time-2*i+1)
}
