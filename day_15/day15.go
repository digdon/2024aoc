package main

import (
	"bufio"
	"log"
	"os"
	"whatever/part1"
	"whatever/part2"
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
