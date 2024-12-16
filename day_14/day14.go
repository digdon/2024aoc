package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	robots := []Robot{}

	for _, line := range inputLines {
		px, py, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, Robot{px, py, vx, vy})
	}

	maxX, maxY := 101, 103

	// Part 1
	quadCounts := calcQuadrants(calcRobotPositions(robots, maxX, maxY, 100), maxX, maxY)
	safetyFactor := 1

	for i, count := range quadCounts {
		fmt.Printf("Quadrant %d: %d\n", i, count)
		safetyFactor *= count
	}

	fmt.Println("Part 1:", safetyFactor)

	part2_column(robots, maxX, maxY)
	part2_unique(robots, maxX, maxY)
	part2_safety(robots, maxX, maxY)
}

func part2_column(robots []Robot, maxX, maxY int) {
	begin := time.Now()
	var positions map[Point]int
	iteration := -1

	for found, i := false, 0; !found && i < maxX*maxY; i++ {
		positions = calcRobotPositions(robots, maxX, maxY, i)

		// Look for any column that has a bunch of robots in it
		columnCounts := map[int]int{}

		for point := range positions {
			columnCounts[point.x]++
		}

		for x, count := range columnCounts {
			if count >= 30 {
				// Found a column with a bunch of robots in it, so now we check to see how many are consecutive
				maxConsecutive := 0
				consecutive := 0
				for y := 0; y < maxY; y++ {
					if positions[Point{x, y}] > 0 {
						consecutive++
					} else {
						if consecutive > maxConsecutive {
							maxConsecutive = consecutive
						}
						consecutive = 0
					}
				}

				if maxConsecutive >= 30 {
					found = true
					iteration = i
					break
				}
			}
		}
	}

	fmt.Println("Part 2 - column:", iteration, time.Since(begin))

	if iteration > -1 {
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				if positions[Point{x, y}] > 0 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func part2_unique(robots []Robot, maxX, maxY int) {
	begin := time.Now()
	var positions map[Point]int
	iteration := -1

	for i := 0; i < maxX*maxY; i++ {
		positions = calcRobotPositions(robots, maxX, maxY, i)

		// It turns out that all of the robots are in unique positions at the target iteration, so we can just check for that
		unique := true
		for _, count := range positions {
			if count > 1 {
				unique = false
				break
			}
		}

		if unique {
			iteration = i
			break
		}
	}

	fmt.Println("Part 2 - unique:", iteration, time.Since(begin))

	if iteration > -1 {
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				if positions[Point{x, y}] > 0 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func part2_safety(robots []Robot, maxX, maxY int) {
	begin := time.Now()
	lowestSF := math.MaxInt
	bestIteration := -1

	for i := 0; i < maxX*maxY; i++ {
		quadCounts := calcQuadrants(calcRobotPositions(robots, maxX, maxY, i), maxX, maxY)
		safetyFactor := 1

		for _, count := range quadCounts {
			if count > 0 {
				safetyFactor *= count
			}
		}

		if safetyFactor < lowestSF {
			lowestSF = safetyFactor
			bestIteration = i
		}
	}

	fmt.Println("Part 2 - safety:", bestIteration, time.Since(begin))
}

func calcRobotPositions(robots []Robot, maxX, maxY int, seconds int) map[Point]int {
	positions := map[Point]int{}

	for _, robot := range robots {
		newpx, newpy := (robot.px+(robot.vx*seconds))%maxX, (robot.py+(robot.vy*seconds))%maxY

		if newpx < 0 {
			newpx = maxX + newpx
		}

		if newpy < 0 {
			newpy = maxY + newpy
		}

		positions[Point{newpx, newpy}]++
	}

	return positions
}

func calcQuadrants(positions map[Point]int, maxX, maxY int) [4]int {
	halfX, halfY := maxX/2, maxY/2
	quadCounts := [4]int{}

	for point, count := range positions {
		// Determine quadrant - skip anything on the centre lines
		if point.x != halfX && point.y != halfY {
			if point.x < halfX {
				if point.y < halfY {
					quadCounts[0] += count
				} else {
					quadCounts[1] += count
				}
			} else {
				if point.y < halfY {
					quadCounts[2] += count
				} else {
					quadCounts[3] += count
				}
			}
		}
	}

	return quadCounts
}

type Robot struct {
	px, py, vx, vy int
}

type Point struct {
	x, y int
}
