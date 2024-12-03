package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	// inputLines = []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
	// inputLines = []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	re, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	var part1Sum, part2Sum int
	doIt := true

	for _, line := range inputLines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, i := range matches {
			if i[0] == "do()" {
				doIt = true
			} else if i[0] == "don't()" {
				doIt = false
			} else {
				first, _ := strconv.Atoi(i[1])
				second, _ := strconv.Atoi(i[2])
				value := first * second
				part1Sum += value
				if doIt {
					part2Sum += value
				}
			}
		}
	}

	fmt.Println("Part 1:", part1Sum)
	fmt.Println("Part 2:", part2Sum)
}
