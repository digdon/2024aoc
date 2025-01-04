package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var numPad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var dirPad = [][]byte{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

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

	numPadMoves := generateNumPadMoves()
	dirPadMoves := generateDirPadMoves()

	// Part 1 stuff
	begin := time.Now()
	part1Score := 0

	for _, value := range inputLines {
		mult, _ := strconv.Atoi(value[:len(value)-1])
		part1Score += getCharCount(value, numPadMoves, dirPadMoves, map[Key]int{}, 2) * mult
	}

	fmt.Printf("Part 1: %d (%s)\n", part1Score, time.Since(begin))

	// Part 2 stuff
	begin = time.Now()
	part2Score := 0

	for _, value := range inputLines {
		mult, _ := strconv.Atoi(value[:len(value)-1])
		part2Score += getCharCount(value, numPadMoves, dirPadMoves, map[Key]int{}, 25) * mult
	}

	fmt.Printf("Part 2: %d (%s)\n", part2Score, time.Since(begin))
}

func getCharCount(code string, numPadMoves, dirPadMoves map[Move]string, cache map[Key]int, depth int) int {
	// Start by generating the moves required for the number pad
	sb := strings.Builder{}
	tempCode := "A" + code
	for i := 1; i < len(tempCode); i++ {
		sb.WriteString(numPadMoves[Move{tempCode[i-1], tempCode[i]}])
		sb.WriteString("A") // every move sequence ends with an 'A'
	}

	// Now generate the moves required for the direction pad
	return countDirChars(sb.String(), dirPadMoves, cache, depth)
}

func countDirChars(path string, dirPadMoves map[Move]string, cache map[Key]int, depth int) int {
	if depth == 0 {
		return len(path)
	}

	if path == "A" {
		return 1
	}

	key := Key{path, depth}
	if val, ok := cache[key]; ok {
		return val
	}

	total := 0

	moveSets := strings.Split(path, "A")
	// fmt.Println("Move sets:", moveSets)
	for move := 0; move < len(moveSets)-1; move++ {
		moves := moveSets[move]
		sb := strings.Builder{}
		for i := 0; i <= len(moves); i++ {
			var start, end byte
			if i == 0 {
				start = 'A'
			} else {
				start = moves[i-1]
			}

			if i == len(moves) {
				end = 'A'
			} else {
				end = moves[i]
			}

			sb.WriteString(dirPadMoves[Move{start, end}])
			sb.WriteString("A")
		}

		// fmt.Println(sb.String())
		total += countDirChars(sb.String(), dirPadMoves, cache, depth-1)
	}

	cache[key] = total

	return total
}

func generateNumPadMoves() map[Move]string {
	numPadPosMap := generatePositionMap(numPad)
	moves := map[Move]string{}

	for startChar, startPos := range numPadPosMap {
		for endChar, endPos := range numPadPosMap {
			if startChar != ' ' && endChar != ' ' {
				moves[Move{startChar, endChar}] = generateMoves(startPos, endPos, true)
			}
		}
	}

	return moves
}

func generateDirPadMoves() map[Move]string {
	dirPadPosMap := generatePositionMap(dirPad)
	moves := map[Move]string{}

	for startChar, startPos := range dirPadPosMap {
		for endChar, endPos := range dirPadPosMap {
			if startChar != ' ' && endChar != ' ' {
				moves[Move{startChar, endChar}] = generateMoves(startPos, endPos, false)
			}
		}
	}

	return moves
}

func generatePositionMap(pad [][]byte) map[byte]Point {
	posMap := map[byte]Point{}
	for y := 0; y < len(pad); y++ {
		for x := 0; x < len(pad[y]); x++ {
			posMap[pad[y][x]] = Point{x, y}
		}
	}
	return posMap
}

func generateMoves(start, end Point, numPadMove bool) string {
	diffX, diffY := end.x-start.x, end.y-start.y

	// Handle pad-specific special cases first
	if numPadMove {
		// num pad special cases
		if start.y == len(numPad)-1 && end.x == 0 {
			// Starting on the bottom row and moving to the far left, so go up first to avoid the empty space
			return getVerticalMoves(diffY) + getHorizontalMoves(diffX)
		} else if start.x == 0 && end.y == len(numPad)-1 {
			// Starting on the far left and moving to the bottom row, so go right first to avoid the empty space
			return getHorizontalMoves(diffX) + getVerticalMoves(diffY)
		}
	} else {
		// dir pad special cases
		if start.y == 0 && end.x == 0 {
			// Starting on the top row and moving to the far left, so go down first to avoid the empty space
			return getVerticalMoves(diffY) + getHorizontalMoves(diffX)
		} else if start.x == 0 && end.y == 0 {
			// Starting on the far left and moving to the top row, so go right first to avoid the empty space
			return getHorizontalMoves(diffX) + getVerticalMoves(diffY)
		}
	}

	if (diffY < 0 || diffY > 0) && diffX < 0 {
		// up/down and left
		return getHorizontalMoves(diffX) + getVerticalMoves(diffY)
	} else if (diffY < 0 || diffY > 0) && diffX > 0 {
		// up/down and right
		return getVerticalMoves(diffY) + getHorizontalMoves(diffX)
	} else {
		// Whatever is left (eg: diffY/diffX == 0) - doesn't matter
		return getVerticalMoves(diffY) + getHorizontalMoves(diffX)
	}
}

func getVerticalMoves(diffY int) string {
	if diffY > 0 {
		// Moving down
		return strings.Repeat("v", diffY)
	} else {
		// Moving up
		return strings.Repeat("^", -diffY)
	}
}

func getHorizontalMoves(diffX int) string {
	if diffX > 0 {
		// Moving right
		return strings.Repeat(">", diffX)
	} else {
		// Moving left
		return strings.Repeat("<", -diffX)
	}
}

type Key struct {
	code  string
	depth int
}

type Move struct {
	start, end byte
}

type Point struct {
	x, y int
}
