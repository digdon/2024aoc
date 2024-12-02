package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputLines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var part1Count, part2Count int

	for _, report := range inputLines {
		values := []int{}

		for _, item := range strings.Split(report, " ") {
			if value, err := strconv.Atoi(item); err == nil {
				values = append(values, value)
			}
		}

		safe := isSafe(values)

		if safe {
			part1Count++
			part2Count++
		} else {
			// Part 2 stuff - remove one value at a time and check if it's safe
			for i := 0; i < len(values); i++ {
				newValues := append([]int{}, values[:i]...)
				newValues = append(newValues, values[i+1:]...)

				if isSafe(newValues) {
					part2Count++
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", part1Count)
	fmt.Println("Part 2:", part2Count)
}

func isSafe(values []int) bool {
	pv, inc := values[0], values[1] >= values[0]

	for i := 1; i < len(values); i++ {
		diff := values[i] - pv

		if (inc && (diff < 1 || diff > 3)) || (!inc && (diff > -1 || diff < -3)) {
			return false
		}

		pv = values[i]
	}

	return true
}
