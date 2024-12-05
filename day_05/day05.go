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

	// Parse the input for the table of page pre-requisites and the list of page updates
	prereqs := map[int]map[int]bool{}
	updates := [][]int{}
	evalUpdates := false

	for _, line := range inputLines {
		if line == "" {
			// Prereqs come first, followed by a newline, then the updates
			// We've come to the newline, so switch to processing updates
			evalUpdates = true
			continue
		}

		if evalUpdates {
			list := []int{}
			for _, item := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(item)
				list = append(list, value)
			}
			updates = append(updates, list)
		} else {
			var first, second int
			fmt.Sscanf(line, "%d|%d", &first, &second)

			if _, found := prereqs[second]; !found {
				prereqs[second] = map[int]bool{}
			}

			prereqs[second][first] = true
		}
	}

	var part1Total, part2Total int

	for _, update := range updates {
		if inOrder(update, prereqs) {
			part1Total += update[len(update)/2]
		} else {
			update = reorder(update, prereqs)
			part2Total += update[len(update)/2]
		}
	}

	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)
}

func inOrder(update []int, prereqs map[int]map[int]bool) bool {
	// Build the set of pages being updated
	contains := map[int]bool{}
	for _, page := range update {
		contains[page] = true
	}

	// Process the page updates, in order, ensuring that if a page has any prereqs, they have been seen
	seen := map[int]bool{}
	valid := true

	for _, page := range update {
		if pagePrereqs, found := prereqs[page]; found {
			// Page has prereqs - prereqs only matter if both the page and prereq(s) are in the same update
			// Check each prereq to see if it is included in the update and, if it is, that it has been seen
			for prereq := range pagePrereqs {
				if contains[prereq] && !seen[prereq] {
					// Page has a prereq contained in the update, but it hasn't been seen yet
					valid = false
					break
				}
			}
		}

		if !valid {
			break
		}

		seen[page] = true
	}

	return valid
}

func reorder(update []int, prereqs map[int]map[int]bool) []int {
	// Build the set of pages being updated
	contains := map[int]bool{}
	for _, page := range update {
		contains[page] = true
	}

	// Topological sort
	sortedList := make([]int, 0, len(update))
	seen := map[int]bool{}

	var visitAll func(items []int)
	visitAll = func(items []int) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				if prereq, found := prereqs[item]; found {
					// The current page has prereqs - process them first
					temp := make([]int, 0, len(prereq))
					for x := range prereq {
						// Only include prereqs that are contained in the update
						if contains[x] {
							temp = append(temp, x)
						}
					}
					visitAll(temp)
				}

				sortedList = append(sortedList, item)
			}
		}
	}

	visitAll(update)

	return sortedList
}
