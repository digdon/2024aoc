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
	fmt.Println("bfs done in", time.Since(begin))

	begin = time.Now()
	part1, part2 := calculateCheats(path)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	fmt.Println("Time:", time.Since(begin))
}

func calculateCheats(path []Point) (int, int) {
	part1Count, part2Count := 0, 0

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
			saved := (j - i) - dist

			if dist <= 2 && saved >= 100 {
				part1Count++
			}

			if dist <= 20 && saved >= 100 {
				part2Count++
			}
		}
	}

	return part1Count, part2Count
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
