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
	part1Results := part1(inputLines)
	part1Time := time.Since(begin)

	begin = time.Now()
	part2Results := part2(inputLines)
	part2Time := time.Since(begin)

	fmt.Printf("Part 1: %d (in %s)\n", part1Results, part1Time)
	fmt.Printf("Part 2: %d (in %s)\n", part2Results, part2Time)
}

func part1(inputLines []string) int {
	var part1Total int

	for _, line := range inputLines {
		colonIdx := strings.Index(line, ": ")
		targetValue, _ := strconv.Atoi(line[:colonIdx])
		parts := strings.Split(line[colonIdx+2:], " ")
		values := []int{}

		for _, part := range parts {
			value, _ := strconv.Atoi(part)
			values = append(values, value)
		}

		for addMul := 0; addMul < (1 << (len(values) - 1)); addMul++ {
			result := values[0]

			for i := 1; i < len(values); i++ {
				if addMul&(1<<(i-1)) != 0 {
					result += values[i]
				} else {
					result *= values[i]
				}
			}

			if result == targetValue {
				part1Total += targetValue
				break
			}
		}
	}

	return part1Total
}

func part2(inputLines []string) int {
	var part2Total int

	for _, line := range inputLines {
		colonIdx := strings.Index(line, ": ")
		targetValue, _ := strconv.Atoi(line[:colonIdx])
		parts := strings.Split(line[colonIdx+2:], " ")
		values := []int{}

		for _, part := range parts {
			value, _ := strconv.Atoi(part)
			values = append(values, value)
		}

		cache := map[string]bool{}

		for matched, concat := false, 0; !matched && concat < (1<<(len(parts)-1)); concat++ {
			for addMul := 0; addMul < (1 << (len(values) - 1)); addMul++ {
				var sb strings.Builder
				for i := 0; i < len(values)-1; i++ {
					if concat&(1<<i) != 0 {
						sb.WriteString("|")
					} else if addMul&(1<<i) != 0 {
						sb.WriteString("+")
					} else {
						sb.WriteString("*")
					}
				}

				key := sb.String()
				if cache[key] {
					continue
				}

				result := values[0]

				for i := 1; i < len(values); i++ {
					if concat&(1<<(i-1)) != 0 {
						result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(values[i]))
					} else if addMul&(1<<(i-1)) != 0 {
						result += values[i]
					} else {
						result *= values[i]
					}
				}

				if result == targetValue {
					display(targetValue, values, addMul, concat, result)
					part2Total += targetValue
					matched = true
					break
				}

				cache[key] = true
			}
		}
	}

	return part2Total
}

func display(targetValue int, values []int, addMul, concat, result int) {
	fmt.Print("Target: ", targetValue, " -> ")

	for i, value := range values {
		if i > 0 {
			if concat&(1<<(i-1)) != 0 {
				print(" || ")
			} else if addMul&(1<<(i-1)) != 0 {
				print(" + ")
			} else {
				print(" * ")
			}
		}

		print(value)
	}

	fmt.Println(" = ", result)

}
