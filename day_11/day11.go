package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	// input := "125 17"
	input := "64599 31 674832 2659361 1 0 8867 321"

	countMap := map[int]int{}
	parts := strings.Split(input, " ")

	for _, part := range parts {
		value, _ := strconv.Atoi(part)
		countMap[value]++
	}

	// fmt.Println(countMap)
	var part1Count int
	begin := time.Now()

	for i := 1; i <= 75; i++ {
		newCountMap := map[int]int{}

		for key, count := range countMap {
			if key == 0 {
				newCountMap[1] += count
			} else {
				length := int(math.Log10(float64(key)) + 1)

				if length%2 == 0 {
					exp := int(math.Pow10(int(length / 2)))
					newCountMap[key/exp] += count
					newCountMap[key%exp] += count
				} else {
					newCountMap[key*2024] += count
				}
			}
		}

		countMap = newCountMap

		if i == 25 {
			for _, count := range countMap {
				part1Count += count
			}
		}
	}

	part2Count := 0
	for _, count := range countMap {
		part2Count += count
	}

	fmt.Println("Part 1:", part1Count)
	fmt.Println("Part 2:", part2Count)
	fmt.Println(time.Since(begin))
}
