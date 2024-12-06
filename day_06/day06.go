package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputLines [][]byte
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputLines = append(inputLines, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Start by finding guard
	var startX, startY, startDir int

	for y, guardFound := 0, false; y < len(inputLines) && !guardFound; y++ {
		for x := 0; x < len(inputLines[y]); x++ {
			char := inputLines[y][x]
			if char != '.' && char != '#' {
				startX, startY = x, y
				startDir = charToDirMap[char]
				guardFound = true
				break
			}
		}
	}

	// Part 1
	visited := part1Locations(inputLines, startX, startY, startDir)
	fmt.Println("Part 1:", len(visited))

	// Part 2
	delete(visited, Point{startX, startY}) // We remove this because we won't try adding an obstacle at the starting point
	path := []Point{}
	for point := range visited {
		path = append(path, point)
	}
	count := part2(inputLines, startX, startY, startDir, path)
	fmt.Println("Part 2:", count)
}

func part1Locations(grid [][]byte, x, y, dir int) map[Point]bool {
	visited := map[Point]bool{}

	for {
		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
			break
		}

		visited[Point{x, y}] = true
		tempX, tempY := x+Directions[dir][0], y+Directions[dir][1]
		if tempY >= 0 && tempY < len(grid) && tempX >= 0 && tempX < len(grid[tempY]) && grid[tempY][tempX] == '#' {
			// Obstacle - turn right
			dir = (dir + 1) % 4
		} else {
			// Move forward
			x, y = tempX, tempY
		}
	}

	return visited
}

func part2(grid [][]byte, startX, startY, startDir int, possible []Point) int {
	var count int

	// For every point the guard originally reached, add an obstacle and see if we can get the guard to form a loop.
	// We already know that the guard can only reach these particular points if there are no obstacles, so that means
	// we only need to try putting obstacles along this path.
	for _, point := range possible {
		// Set an obstacle at the next point
		grid[point.y][point.x] = '#'

		// Reset everything
		visited := map[Path]bool{}
		loopFound := false
		x, y, dir := startX, startY, startDir
		checkVisited := false

		// Move the guard
		for {
			if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
				// Fell off the edge, so no loop found
				break
			}

			position := Path{Point{x, y}, dir}

			// Check for looping. However, we only need to start checking after we've reached the obstacle we placed
			if checkVisited && visited[position] {
				// We've reached this spot, in this direction, already once, so we must be in a loop
				loopFound = true
				break
			}

			visited[position] = true
			tempX, tempY := x+Directions[dir][0], y+Directions[dir][1]
			if tempY >= 0 && tempY < len(grid) && tempX >= 0 && tempX < len(grid[tempY]) && grid[tempY][tempX] == '#' {
				// Obstacle - turn right
				dir = (dir + 1) % 4

				if !checkVisited && tempX == point.x && tempY == point.y {
					// We've reached the obstacle we placed, so now we can enable the check for looping
					checkVisited = true
				}
			} else {
				// Move forward
				x, y = tempX, tempY
			}
		}

		if loopFound {
			// fmt.Println("loop created by obstacle at", point)
			count++
		}

		// Set the current obstacle point back to empty
		grid[point.y][point.x] = '.'
	}

	return count
}

type Path struct {
	Point
	dir int
}

type Point struct {
	x, y int
}

var charToDirMap = map[byte]int{
	'^': UP,
	'>': RIGHT,
	'v': DOWN,
	'<': LEFT,
}

const (
	UP int = 0
	RIGHT
	DOWN
	LEFT
)

var Directions = [][]int{
	{0, -1}, // UP
	{1, 0},  // RIGHT
	{0, 1},  // DOWN
	{-1, 0}, // LEFT
}
