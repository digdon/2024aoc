package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	freqMap := map[byte][]Point{}

	// Find all of the antenna locations and group by frequency
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			char := grid[y][x]

			if char == '#' {
				grid[y][x] = '.'
			}

			if char != '.' && char != '#' {
				freqMap[char] = append(freqMap[char], Point{x, y})
			}
		}
	}

	// Now process each group of antennas. Find the (manhattan) distance between each pair of antennas
	// and then create a set of antinodes for each pair of antennas
	part1NodeMap := map[Point]int{}
	part2NodeMap := map[Point]int{}

	for _, points := range freqMap {
		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				diffX, diffY := points[i].x-points[j].x, points[i].y-points[j].y

				// Part 1 stuff - we only need 1 antinode in each direction
				addAntinodes(part1NodeMap, points[i], diffX, diffY, len(grid[0]), len(grid), true)
				addAntinodes(part1NodeMap, points[j], -diffX, -diffY, len(grid[0]), len(grid), true)

				// Part 2 stuff - multiple antinodes in each direction
				addAntinodes(part2NodeMap, points[i], diffX, diffY, len(grid[0]), len(grid), false)
				addAntinodes(part2NodeMap, points[j], -diffX, -diffY, len(grid[0]), len(grid), false)

				// You may be wondering about this.... part two means that antinodes expand repeatedly
				// in all directions. The two lines above cover the points less than point i and greater than point j,
				// but i and j themselves would also be antinodes, so we need to add them as well
				part2NodeMap[points[i]]++
				part2NodeMap[points[j]]++
			}
		}
	}

	// fmt.Println("--------")
	// for row := 0; row < len(grid); row++ {
	// 	for col := 0; col < len(grid[row]); col++ {
	// 		char := grid[row][col]

	// 		if part2NodeMap[Point{col, row}] > 0 {
	// 			char = '#'
	// 		}

	// 		fmt.Printf("%c", char)
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("Part 1:", len(part1NodeMap))
	fmt.Println("Part 2:", len(part2NodeMap))
}

func addAntinodes(nodeMap map[Point]int, point Point, diffX, diffY, maxX, maxY int, limit bool) {
	for x, y := point.x+diffX, point.y+diffY; x >= 0 && x < maxX && y >= 0 && y < maxY; x, y = x+diffX, y+diffY {
		nodeMap[Point{x, y}]++

		if limit {
			break
		}
	}
}

type Point struct {
	x, y int
}
