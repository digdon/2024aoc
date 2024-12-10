package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var grid [][]byte
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Find all of the trailheads
	trailheads := []Point{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				trailheads = append(trailheads, Point{x, y})
			}
		}
	}

	begin := time.Now()
	part1Sum := 0

	for _, trailhead := range trailheads {
		part1Sum += trailheadScore(grid, trailhead, false)
		// fmt.Println("Trailhead", trailhead, "score:", score)
	}

	fmt.Println("Part 1:", part1Sum, time.Since(begin))

	begin = time.Now()
	part2Sum := 0

	for _, trailhead := range trailheads {
		part2Sum += trailheadScore(grid, trailhead, true)
		// fmt.Println("Trailhead", trailhead, "score:", score)
	}

	fmt.Println("Part 2:", part2Sum, time.Since(begin))
}

func trailheadScore(grid [][]byte, trailhead Point, allTrails bool) int {
	var score int
	queue := []Point{trailhead}
	visited := map[Point]bool{}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		// When allTrails is false, points can only be visited a single time. This means we find a single path to each '9', if possible.
		// When it's true, points can be visited multiple times, allowing us to find many different paths to each '9'.
		if !allTrails {
			if _, ok := visited[point]; ok {
				continue
			}
		}

		visited[point] = true

		if grid[point.y][point.x] == '9' {
			score++
			continue
		}

		for _, dir := range Directions {
			nextX, nextY := point.x+dir[0], point.y+dir[1]

			if nextY < 0 || nextY >= len(grid) || nextX < 0 || nextX >= len(grid[nextY]) || grid[nextY][nextX] == '.' {
				continue
			}

			if grid[nextY][nextX]-grid[point.y][point.x] != 1 {
				continue
			}

			queue = append(queue, Point{nextX, nextY})
		}
	}

	return score
}

var Directions [][]int = [][]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type Point struct {
	x, y int
}
