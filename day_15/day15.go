package main

import (
	"2024aoc/part1"
	"2024aoc/part2"
	"bufio"
	"log"
	"os"
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

	part1.Solve(inputLines)
	part2.Solve(inputLines)
}
