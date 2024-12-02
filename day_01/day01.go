package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	part1(inputLines)
	part2(inputLines)
}

func part1(inputLines []string) {
	left, right := []int{}, []int{}

	for _, line := range inputLines {
		var x, y int
		fmt.Sscanf(line, "%d %d", &x, &y)
		left = append(left, x)
		right = append(right, y)
	}

	sort.Ints(left)
	sort.Ints(right)

	var diffs int

	for i := 0; i < len(left); i++ {
		diff := right[i] - left[i]

		if diff < 0 {
			diff = -diff
		}

		diffs += diff
	}

	fmt.Println("Part 1:", diffs)
}

func part2(inputLines []string) {
	left := []int{}
	rightMap := map[int]int{}

	for _, line := range inputLines {
		var x, y int
		fmt.Sscanf(line, "%d %d", &x, &y)
		left = append(left, x)
		rightMap[y]++
	}

	score := 0

	for _, x := range left {
		score += x * rightMap[x]
	}

	fmt.Println("Part 2:", score)
}
