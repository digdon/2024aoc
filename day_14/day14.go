package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Part 2
	var positions map[Point]int
	iteration := -1

	for found, i := false, 0; !found && i < 10000; i++ {
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

	fmt.Println("Part 2:", iteration)

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

func part1(robots []Robot) {
	seconds := 100
	maxX, maxY := 101, 103
	halfX, halfY := maxX/2, maxY/2
	quadCounts := [4]int{}

	// Part 1
	for _, robot := range robots {
		newpx, newpy := (robot.px+(robot.vx*seconds))%maxX, (robot.py+(robot.vy*seconds))%maxY

		if newpx < 0 {
			newpx = maxX + newpx
		}

		if newpy < 0 {
			newpy = maxY + newpy
		}

		// Determine quadrant - skip anything on the centre lines
		if newpx != halfX && newpy != halfY {
			if newpx < halfX {
				if newpy < halfY {
					quadCounts[0]++
				} else {
					quadCounts[1]++
				}
			} else {
				if newpy < halfY {
					quadCounts[2]++
				} else {
					quadCounts[3]++
				}
			}
		}
	}

	safetyFactor := 1

	for i, count := range quadCounts {
		fmt.Printf("Quadrant %d: %d\n", i, count)
		safetyFactor *= count
	}

	fmt.Println("Safety factor:", safetyFactor)
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
