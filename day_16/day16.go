package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

func main() {
	grid := map[Point]bool{}
	var start, end Point

	scanner := bufio.NewScanner(os.Stdin)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, ch := range line {
			if ch == 'S' {
				start = Point{x, y}
				grid[Point{x, y}] = true
			} else if ch == 'E' {
				end = Point{x, y}
				grid[Point{x, y}] = true
			} else if ch == '.' {
				grid[Point{x, y}] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	begin := time.Now()
	cost, paths := dijkstra(grid, start, end)
	pointSet := map[Point]bool{}
	for _, path := range paths {
		for k := range path {
			pointSet[k] = true
		}
	}

	fmt.Println("Part 1:", cost)
	fmt.Println("Part 2:", len(pointSet))
	fmt.Println("Time:", time.Since(begin))
}

func dijkstra(grid map[Point]bool, start, end Point) (int, []map[Point]bool) {
	bestCost := math.MaxInt
	var bestPaths []map[Point]bool
	pointCosts := map[Key]int{}
	queue := []Step{{start, 0, EAST, map[Point]bool{start: true}}}

	for len(queue) > 0 && queue[0].cost <= bestCost {
		curr := queue[0]
		queue = queue[1:]

		if curr.point == end {
			if curr.cost < bestCost {
				bestCost = curr.cost
				bestPaths = []map[Point]bool{}
			}

			if curr.cost == bestCost {
				fmt.Println("Found end!")
				fmt.Println("Cost:", curr.cost)
				// fmt.Println(curr.path)
				bestPaths = append(bestPaths, curr.path)
			}

			continue
		}

		for _, dir := range DirectionList {
			nextPoint := Point{curr.point.x + Directions[dir][0], curr.point.y + Directions[dir][1]}

			if !grid[nextPoint] {
				continue
			}

			nextCost := curr.cost + (rotateCosts[curr.dir][dir] * 1000) + 1

			key := Key{nextPoint.x, nextPoint.y, int(dir)}

			if neighbourCost, ok := pointCosts[key]; !ok || nextCost <= neighbourCost {
				pointCosts[key] = nextCost
				nextPath := make(map[Point]bool, len(curr.path)+1)
				for k := range curr.path {
					nextPath[k] = true
				}
				nextPath[nextPoint] = true
				queue = append(queue, Step{nextPoint, nextCost, dir, nextPath})
			}
		}

		// Sort the queue by cost - by doing this, I make it a priority queue and I don't have to follow every active path
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})
	}

	return bestCost, bestPaths
}

type Dir int

const (
	NORTH Dir = iota
	EAST
	SOUTH
	WEST
)

var DirectionList = []Dir{NORTH, EAST, SOUTH, WEST}

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

type Key struct {
	x, y, dir int
}

type Step struct {
	point Point
	cost  int
	dir   Dir
	path  map[Point]bool
}

type Point struct {
	x, y int
}
