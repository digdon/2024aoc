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

	begin := time.Now()

	lockMode := false
	locks, keys := [][]int{}, [][]int{}

	for i := 0; i < len(inputLines); {
		// Count number of rows for this item
		itemRows := 0
		for ; itemRows+i < len(inputLines) && inputLines[i+itemRows] != ""; itemRows++ {
		}

		if inputLines[i] == "#####" {
			lockMode = true
		} else {
			lockMode = false
		}

		pinCounts := []int{}

		for col := 0; col < len(inputLines[i]); col++ {
			count := 0
			for row := 1; row < itemRows-1; row++ {
				if inputLines[i+row][col] == '#' {
					count++
				}
			}
			pinCounts = append(pinCounts, count)
		}

		if lockMode {
			locks = append(locks, pinCounts)
		} else {
			keys = append(keys, pinCounts)
		}

		i += itemRows + 1
	}

	lockKeyCount := 0

	for _, lock := range locks {
		for _, key := range keys {
			match := true
			for i := range lock {
				if (key[i] + lock[i]) > 5 {
					match = false
					break
				}
			}

			if match {
				// fmt.Printf("Match! lock %v and key %v\n", lock, key)
				lockKeyCount++
			}
		}
	}

	fmt.Printf("Part 1: %d (%s)\n", lockKeyCount, time.Since(begin))
}
