package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	var grid [][]rune
	var start, end Point

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for i, c := range line {
			row = append(row, c)

			if c == 'S' {
				start = Point{i, len(grid)}
			} else if c == 'E' {
				end = Point{i, len(grid)}
			}
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cost, paths := dijkstra(grid, start, end)
	fmt.Println("Part 1:", cost)
	for _, path := range paths {
		fmt.Println(path)
	}
}

func dijkstra(grid [][]rune, start, end Point) (int, [][]Point) {
	totalCost := math.MaxInt

	// Initialize the distance/cost map
	costMap := map[Point]int{start: 0}

	// Initialize the queue
	queue := []Step{{start, 0, EAST}}

	// Start the algorithm
	for len(queue) > 0 {
		// Get the current point
		curr := queue[0]
		queue = queue[1:]

		// Check if we have reached the end
		if curr.point == end {
			if curr.cost < totalCost {
				fmt.Println("New path!", curr.cost)
				totalCost = curr.cost
			} else if curr.cost == totalCost {
				fmt.Println("Matching path!", curr.cost)
			}

			continue
		}

		// Check neighbour elements
		for label, dir := range Directions {
			nx, ny := curr.point.x+dir[0], curr.point.y+dir[1]

			if grid[ny][nx] == '#' {
				continue
			}

			neighbour := Point{nx, ny}
			rotateCost := rotateCosts[curr.dir][label] * 1000
			altCost := curr.cost + rotateCost + 1
			neighbourCost, ok := costMap[neighbour]

			if !ok {
				neighbourCost = math.MaxInt
			}

			if altCost <= neighbourCost {
				costMap[neighbour] = altCost
				queue = append(queue, Step{neighbour, altCost, label})
			}
		}
	}

	var successfulPaths [][]Point

	return costMap[end], successfulPaths
}

type Dir int

const (
	NORTH Dir = iota
	EAST
	SOUTH
	WEST
)

var Directions = map[Dir][2]int{
	NORTH: {0, -1},
	EAST:  {1, 0},
	SOUTH: {0, 1},
	WEST:  {-1, 0},
}

var rotateCosts [][]int = [][]int{
	{0, 1, 2, 1},
	{1, 0, 1, 2},
	{2, 1, 0, 1},
	{1, 2, 1, 0},
}

type Step struct {
	point Point
	cost  int
	dir   Dir
}

type Point struct {
	x, y int
}
