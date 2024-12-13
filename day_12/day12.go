package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var grid [][]rune
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	regionMap := map[rune][]Region{}
	visited := map[Point]bool{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			currentChar := grid[y][x]

			if visited[Point{x, y}] {
				// We've already processed this point
				continue
			}

			// "Flood fill" from this point to find all points in this region
			xQueue, yQueue := []int{x}, []int{y}
			area, perimeter, corners := 0, 0, 0

			for len(xQueue) > 0 {
				x, y := xQueue[0], yQueue[0]
				xQueue, yQueue = xQueue[1:], yQueue[1:]

				if visited[Point{x, y}] {
					// Already been here
					continue
				}

				visited[Point{x, y}] = true
				area++

				// Is this point a corner of a region?
				corners += cornerCount(grid, x, y)

				for _, dir := range Directions {
					nextX, nextY := x+dir[0], y+dir[1]

					if nextY < 0 || nextY >= len(grid) || nextX < 0 || nextX >= len(grid[nextY]) {
						perimeter++
						continue
					}

					nextChar := grid[nextY][nextX]

					if nextChar == currentChar {
						xQueue, yQueue = append(xQueue, nextX), append(yQueue, nextY)
					} else if nextChar != currentChar {
						perimeter++
					}
				}
			}

			regionMap[rune(currentChar)] = append(regionMap[rune(currentChar)], Region{area: area, perimeter: perimeter, corners: corners})
		}
	}

	var part1Price, part2Price int
	for _, regionList := range regionMap {
		for _, region := range regionList {
			part1Price += region.area * region.perimeter
			part2Price += region.area * region.corners
		}
	}

	// fmt.Println(regionMap)

	fmt.Println("Part 1:", part1Price)
	fmt.Println("Part 2:", part2Price)
}

func cornerCount(grid [][]rune, x, y int) int {
	var count int

	currentChar := grid[y][x]

	// Check up and left for convex corner - up and left are both different from the current character
	if (x-1 < 0 || grid[y][x-1] != currentChar) && (y-1 < 0 || grid[y-1][x] != currentChar) {
		count++
	}

	// Check up and right for convex corner - up and right are both different from the current character
	if (x+1 >= len(grid[y]) || grid[y][x+1] != currentChar) && (y-1 < 0 || grid[y-1][x] != currentChar) {
		count++
	}

	// Check down and left for convex corner - down and left are both different from the current character
	if (x-1 < 0 || grid[y][x-1] != currentChar) && (y+1 >= len(grid) || grid[y+1][x] != currentChar) {
		count++
	}

	// Check down and right for convex corner - down and right are both different from the current character
	if (x+1 >= len(grid[y]) || grid[y][x+1] != currentChar) && (y+1 >= len(grid) || grid[y+1][x] != currentChar) {
		count++
	}

	// Check up and left for concave corner - up and left are both the same as the current character, and up/left is different
	if x-1 >= 0 && grid[y][x-1] == currentChar && y-1 >= 0 && grid[y-1][x] == currentChar && grid[y-1][x-1] != currentChar {
		count++
	}

	// Check up and right for concave corner - up and right are both the same as the current character, and up/right is different
	if x+1 < len(grid[y]) && grid[y][x+1] == currentChar && y-1 >= 0 && grid[y-1][x] == currentChar && grid[y-1][x+1] != currentChar {
		count++
	}

	// Check down and left for concave corner - down and left are both the same as the current character, and down/left is different
	if x-1 >= 0 && grid[y][x-1] == currentChar && y+1 < len(grid) && grid[y+1][x] == currentChar && grid[y+1][x-1] != currentChar {
		count++
	}

	// Check down and right for concave corner - down and right are both the same as the current character, and down/right is different
	if x+1 < len(grid[y]) && grid[y][x+1] == currentChar && y+1 < len(grid) && grid[y+1][x] == currentChar && grid[y+1][x+1] != currentChar {
		count++
	}

	return count
}

var Directions = [4][2]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type Region struct {
	area, perimeter, corners int
}

type Point struct {
	x, y int
}
