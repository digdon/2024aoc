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

	left, right, rightMap := []int{}, []int{}, map[int]int{}

	for _, line := range inputLines {
		var x, y int
		fmt.Sscanf(line, "%d %d", &x, &y)
		left = append(left, x)
		right = append(right, y)
		rightMap[y]++
	}

	sort.Ints(left)
	sort.Ints(right)

	// Part 1
	var diffs int

	for i := 0; i < len(left); i++ {
		diff := right[i] - left[i]

		if diff < 0 {
			diff = -diff
		}

		diffs += diff
	}

	fmt.Println("Part 1:", diffs)

	// Part 2
	var score int

	for _, x := range left {
		score += x * rightMap[x]
	}

	fmt.Println("Part 2:", score)
}
