package main

import (
	"bufio"
	"fmt"
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

	partOne(inputLines)
	partTwo(inputLines)
}

func partOne(inputLines []string) {
	var count int

	for y := 0; y < len(inputLines); y++ {
		for x := 0; x < len(inputLines[y]); x++ {
			if inputLines[y][x] == TARGET[0] {
				// Found a possible starting point, so now look for full string in every direction
				for _, dir := range DIRECTIONS {
					found := true
					for i, newX, newY := 1, x+dir[0], y+dir[1]; i < len(TARGET); i, newX, newY = i+1, newX+dir[0], newY+dir[1] {
						if newX < 0 || newX >= len(inputLines[y]) || newY < 0 || newY >= len(inputLines) || inputLines[newY][newX] != TARGET[i] {
							found = false
							break
						}
					}

					if found {
						count++
					}
				}
			}
		}
	}

	fmt.Println("Part 1:", count)
}

var TARGET = "XMAS"

var DIRECTIONS = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
}

var count int

func partTwo(inputLines []string) {
	// In order for MAS to be in an 'X' shape, we know we cannot be along the edge, so we only need
	// to check the inner parts of the grid for the 'A' character
	for y := 1; y < len(inputLines)-1; y++ {
		for x := 1; x < len(inputLines[y])-1; x++ {
			if inputLines[y][x] == 'A' {
				// Found an 'A', so build strings in the two main diagonal directions (upper left to bottom right and upper right to bottom left)
				first := string(inputLines[y-1][x-1]) + "A" + string(inputLines[y+1][x+1])
				second := string(inputLines[y-1][x+1]) + "A" + string(inputLines[y+1][x-1])

				// Each diagonal string can only correctly have one of two values - SAM or MAS. If this is true, we've got a hit
				if (first == "MAS" || first == "SAM") && (second == "MAS" || second == "SAM") {
					count++
				}
			}
		}
	}

	fmt.Println("Part 2:", count)
}
