package part1

import (
	"fmt"
	"strings"
	"time"
)

func Solve(inputLines []string) {
	begin := time.Now()

	// Build the grid, find the robot start location, and build the movement list
	grid := [][]byte{}
	moves := []string{}
	rx, ry := -1, -1

	for _, line := range inputLines {
		if line == "" {
			break
		}

		grid = append(grid, []byte(line))

		if pos := strings.Index(line, "@"); pos != -1 {
			rx, ry = pos, len(grid)-1
		}
	}

	for _, line := range inputLines[len(grid)+1:] {
		moves = append(moves, line)
	}

	// fmt.Println(rx, ry)
	// fmt.Println(moves)
	grid[ry][rx] = '.'

	// Start moving the robot
	for _, moveline := range moves {
		for _, move := range moveline {
			dir := directions[move]
			nextX, nextY := rx+dir[0], ry+dir[1]

			if grid[nextY][nextX] == '#' {
				// Hit an obstacle, nothing to do
				continue
			} else if grid[nextY][nextX] == 'O' {
				// Found a box - try to move it
				if moveBox(grid, nextX, nextY, dir) {
					// Box was successfully moved, so move the robot
					rx, ry = nextX, nextY
				}
			} else {
				// Empty space, move the robot
				rx, ry = nextX, nextY
			}

			// displayGrid(grid, rx, ry, rune(moveline[0]))
		}
	}

	var part1Sum int
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == 'O' {
				// fmt.Printf("Box at %d,%d\n", x, y)
				part1Sum += (y * 100) + x
			}
		}
	}

	fmt.Printf("Part 1: %d (%s)\n", part1Sum, time.Since(begin))
}

func displayGrid(grid [][]byte, rx, ry int, move rune) {
	fmt.Printf("Move %c:\n", move)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if x == rx && y == ry {
				fmt.Printf("@")
			} else {
				fmt.Printf("%c", grid[y][x])
			}
		}

		fmt.Println()
	}
}

func moveBox(grid [][]byte, bx, by int, dir [2]int) bool {
	nextX, nextY := bx+dir[0], by+dir[1]
	moved := false

	if grid[nextY][nextX] == '#' {
		// Hit an obstacle and cannot move the box
		return false
	}

	if grid[nextY][nextX] == 'O' {
		// The box is trying to move into another box - try to move that next box
		moved = moveBox(grid, nextX, nextY, dir)

		if !moved {
			// The next box could not be moved, so this box also cannot be moved
			return false
		}
	}

	// Move the box into the empty space
	grid[nextY][nextX] = 'O'
	grid[by][bx] = '.'

	return true
}

var directions = map[rune][2]int{
	'^': {0, -1}, // up
	'v': {0, 1},  // down
	'<': {-1, 0}, // left
	'>': {1, 0},  // right
}
