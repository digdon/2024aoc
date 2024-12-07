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

	part1(inputLines)
	part2(inputLines)
}

func part1(inputLines []string) {
	var part1Total int

	begin := time.Now().UnixMilli()

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

	fmt.Println("Part 1:", part1Total)
	fmt.Println(time.Since(time.UnixMilli(begin)))
}

func part2(inputLines []string) {
	var part2Total int

	begin := time.Now().UnixMilli()

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

	fmt.Println("Part 2:", part2Total)
	fmt.Println(time.Since(time.UnixMilli(begin)))
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
