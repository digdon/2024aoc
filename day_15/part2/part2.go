package part2

import (
	"fmt"
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

		row := []byte{}
		for pos, c := range line {
			if c == '#' {
				row = append(row, '#', '#')
			} else if c == 'O' {
				row = append(row, '[', ']')
			} else {
				row = append(row, '.', '.')

				if c == '@' {
					rx, ry = pos*2, len(grid)
				}
			}
		}
		grid = append(grid, row)
	}

	for _, line := range inputLines[len(grid)+1:] {
		moves = append(moves, line)
	}

	// Start moving the robot
	for _, moveline := range moves {
		for _, move := range moveline {
			dir := directions[move]
			nextX, nextY := rx+dir[0], ry+dir[1]
			ch := grid[nextY][nextX]

			if ch == '#' {
				// Hit an obstacle, nothing to do
				continue
			} else if ch == '[' || ch == ']' {
				// Found a box - try to move it
				if boxCanMove(grid, nextX, nextY, move) {
					// Box can be moved, so move it
					moveBox(grid, nextX, nextY, move)

					// Move the robot
					rx, ry = nextX, nextY
				}
			} else {
				// Empty space, move the robot
				rx, ry = nextX, nextY
			}

			// displayGrid(grid, rx, ry, move)
		}
	}

	var part2Sum int
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if grid[y][x] == '[' {
				// fmt.Printf("Box at %d,%d\n", x, y)
				part2Sum += (y * 100) + x
			}
		}
	}

	fmt.Printf("Part 2: %d (%s)\n", part2Sum, time.Since(begin))
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

func boxCanMove(grid [][]byte, bx, by int, move rune) bool {
	if grid[by][bx] == '#' {
		// Box is against an obstacle and cannot be moved
		return false
	}

	if grid[by][bx] == '.' {
		return true
	}

	if move == '>' {
		return boxCanMove(grid, bx+2, by, move)
	} else if move == '<' {
		return boxCanMove(grid, bx-2, by, move)
	} else {
		var addY, leftX, rightX int

		if move == '^' {
			addY = -1
		} else if move == 'v' {
			addY = 1
		}

		if grid[by][bx] == '[' {
			leftX, rightX = bx, bx+1
		} else {
			leftX, rightX = bx-1, bx
		}

		left := boxCanMove(grid, leftX, by+addY, move)
		right := boxCanMove(grid, rightX, by+addY, move)

		if !(left && right) {
			return false
		}
	}

	return true
}

func moveBox(grid [][]byte, bx, by int, move rune) {
	if move == '>' {
		if grid[by][bx+2] == '[' {
			// Another box is in the way, move it first
			moveBox(grid, bx+2, by, move)
		}

		grid[by][bx+1] = '['
		grid[by][bx+2] = ']'
		grid[by][bx] = '.'
	} else if move == '<' {
		if grid[by][bx-2] == ']' {
			// Another box is in the way, move it first
			moveBox(grid, bx-2, by, move)
		}

		grid[by][bx-2] = '['
		grid[by][bx-1] = ']'
		grid[by][bx] = '.'
	} else {
		var tempBy, leftBx, rightBx int

		if move == '^' {
			tempBy = by - 1
		} else {
			tempBy = by + 1
		}

		if grid[by][bx] == '[' {
			leftBx, rightBx = bx, bx+1
		} else {
			leftBx, rightBx = bx-1, bx
		}

		if grid[tempBy][leftBx] == '[' || grid[tempBy][leftBx] == ']' {
			moveBox(grid, leftBx, tempBy, move)
		}

		if grid[tempBy][rightBx] == '[' || grid[tempBy][rightBx] == ']' {
			moveBox(grid, rightBx, tempBy, move)
		}

		grid[tempBy][leftBx] = '['
		grid[tempBy][rightBx] = ']'
		grid[by][leftBx] = '.'
		grid[by][rightBx] = '.'
	}
}

var directions = map[rune][2]int{
	'^': {0, -1}, // up
	'v': {0, 1},  // down
	'<': {-1, 0}, // left
	'>': {1, 0},  // right
}
