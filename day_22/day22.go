package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	secretSum := 0
	deltaPriceMap := map[[4]int]int{}

	begin := time.Now()

	for _, line := range inputLines {
		secret, _ := strconv.Atoi(line)
		changes := []int{}
		deltaSeenMap := map[[4]int]bool{}
		prevPrice := secret % 10

		for i := 0; i < 2000; i++ {
			secret = iterateSecret(secret)

			price := secret % 10
			changes = append(changes, price-prevPrice)

			if len(changes) >= 4 {
				if len(changes) > 4 {
					changes = changes[1:]
				}

				deltas := [4]int{changes[0], changes[1], changes[2], changes[3]}

				// Only add the price if we haven't seen this delta before (this is because we're only to use the first
				// occurrence of a delta)
				if !deltaSeenMap[deltas] {
					deltaPriceMap[deltas] += price
					deltaSeenMap[deltas] = true
				}
			}

			prevPrice = price
		}

		secretSum += secret
	}

	fmt.Println("Part 1:", secretSum)

	bestPrice := 0
	for _, price := range deltaPriceMap {
		if price > bestPrice {
			bestPrice = price
		}
	}

	fmt.Println("Part 2:", bestPrice)
	fmt.Println("Time:", time.Since(begin))
}

func iterateSecret(secret int) int {
	// Step 1
	secret = ((secret * 64) ^ secret) % 16777216

	// Step 2
	secret = ((secret / 32) ^ secret) % 16777216

	// Step 3
	secret = ((secret * 2048) ^ secret) % 16777216

	return secret
}
