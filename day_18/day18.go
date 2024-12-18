package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

	maxX, maxY := 70, 70
	firstChunk := 1024
	obstaclesMap := map[Point]bool{}

	for i := 0; i < firstChunk; i++ {
		var x, y int
		fmt.Sscanf(inputLines[i], "%d,%d", &x, &y)
		obstaclesMap[Point{x, y}] = true
	}

	begin := time.Now()
	path := bfs(maxX, maxY, obstaclesMap)
	pathMap := map[Point]bool{}
	for _, p := range path {
		pathMap[p] = true
	}

	fmt.Println("Part 1:", len(path)-1, time.Since(begin))

	begin = time.Now()
	for i := firstChunk; i < len(inputLines); i++ {
		var x, y int
		fmt.Sscanf(inputLines[i], "%d,%d", &x, &y)
		oPoint := Point{x, y}
		obstaclesMap[oPoint] = true

		if !pathMap[oPoint] {
			// New obstacle won't affect the path, so we can skip processing it
			continue
		}

		path := bfs(maxX, maxY, obstaclesMap)

		if path == nil {
			fmt.Printf("Part 2: %d,%d %s\n", oPoint.x, oPoint.y, time.Since(begin))
			break
		}

		newPathMap := map[Point]bool{}
		for _, p := range path {
			newPathMap[p] = true
		}
		pathMap = newPathMap
	}
}

func bfs(maxX, maxY int, obstaclesMap map[Point]bool) []Point {
	start, end := Point{0, 0}, Point{maxX, maxY}

	parentMap := map[Point]Point{}
	visited := map[Point]bool{start: true}
	queue := []Step{{start, 0}}
	// minSteps := 0
	pathFound := false

	for len(queue) > 0 {
		currStep := queue[0]
		queue = queue[1:]

		if currStep.point == end {
			// minSteps = currStep.steps
			pathFound = true
			break
		}

		for _, dir := range Directions {
			nextX, nextY := currStep.point.x+dir[0], currStep.point.y+dir[1]

			if nextX < 0 || nextX > maxX || nextY < 0 || nextY > maxY {
				// Outside the grid
				continue
			}

			nextPoint := Point{nextX, nextY}

			if obstaclesMap[nextPoint] || visited[nextPoint] {
				// Obstacle or already visited
				continue
			}

			visited[nextPoint] = true
			parentMap[nextPoint] = currStep.point
			queue = append(queue, Step{nextPoint, currStep.steps + 1})
		}
	}

	if !pathFound {
		return nil
	}

	// Reconstruct the path
	path := []Point{end}
	for curr := end; curr != start; {
		curr = parentMap[curr]
		path = append([]Point{curr}, path...)
	}

	return path
}

var Directions = [][]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type Step struct {
	point Point
	steps int
}

type Point struct {
	x, y int
}
