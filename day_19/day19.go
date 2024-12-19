package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

	patterns = strings.Split(inputLines[0], ", ")

	part1Count, part2Count := 0, 0
	begin := time.Now()
	for i := 2; i < len(inputLines); i++ {
		matchCount := countMatches(inputLines[i])

		if matchCount > 0 {
			part1Count++
		}

		part2Count += matchCount
	}

	fmt.Println("Part 1:", part1Count)
	fmt.Println("Part 2:", part2Count)
	fmt.Println(time.Since(begin))
}

var patterns []string
var cache = map[string]int{}

func countMatches(design string) int {
	if count, ok := cache[design]; ok {
		return count
	}

	if design == "" {
		return 1
	}

	count := 0

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			count += countMatches(design[len(pattern):])
		}
	}

	cache[design] = count

	return count
}
