package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

	begin := time.Now()

	path := bfs(grid, start, end)
	pathMap := map[Point]int{}

	for step, point := range path {
		pathMap[point] = step
	}

	fmt.Printf("Part 1: %d (%s)\n", part1(grid, path, pathMap), time.Since(begin))
	begin = time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(path), time.Since(begin))
}

func part1(grid [][]rune, path []Point, pathMap map[Point]int) int {
	cheatCount := 0

	for step, point := range path {
		for _, dir := range Directions {
			oneX, oneY := point.x+dir.x, point.y+dir.y
			twoX, twoY := point.x+2*dir.x, point.y+2*dir.y

			if twoY < 0 || twoY >= len(grid) || twoX < 0 || twoX >= len(grid[twoY]) {
				// end of cheat falls off the grid, so skip this direction
				continue
			}

			if grid[oneY][oneX] == '#' && (grid[twoY][twoX] == '.' || grid[twoY][twoX] == 'E') {
				if nextStep, ok := pathMap[Point{twoX, twoY}]; ok {
					// Found a workable cheat
					saved := (nextStep - step) - 2

					if saved >= 100 {
						cheatCount++
					}
				}
			}
		}
	}

	return cheatCount
}

func part2(path []Point) int {
	cheatCount := 0

	for i, point := range path {
		for j := i + 1; j < len(path); j++ {
			// calculate manhattan distance
			diffX := path[j].x - point.x
			if diffX < 0 {
				diffX = -diffX
			}
			diffY := path[j].y - point.y
			if diffY < 0 {
				diffY = -diffY
			}
			dist := diffX + diffY

			if dist > 20 {
				// Too far away, skip
				continue
			}

			saved := (j - i) - dist

			if saved >= 100 {
				cheatCount++
			}
		}
	}

	return cheatCount
}

func bfs(grid [][]rune, start, end Point) []Point {
	parentMap := map[Point]Point{}
	queue := []Point{start}
	visited := map[Point]bool{start: true}
	path := []Point{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			// Found the end, build the path
			for curr != start {
				path = append([]Point{curr}, path...)
				curr = parentMap[curr]
			}

			path = append([]Point{start}, path...)
			break
		}

		for _, dir := range Directions {
			next := Point{curr.x + dir.x, curr.y + dir.y}

			if next.x < 0 || next.x >= len(grid[0]) || next.y < 0 || next.y >= len(grid) {
				continue
			}

			if grid[next.y][next.x] == '#' || visited[next] {
				continue
			}

			parentMap[next] = curr
			visited[next] = true
			queue = append(queue, next)
		}
	}

	return path
}

var Directions = []Point{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type Point struct {
	x, y int
}
