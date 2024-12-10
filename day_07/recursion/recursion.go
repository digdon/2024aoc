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
	fmt.Println("Part 1:", evaluateEquations(inputLines, false), time.Since(begin))
	begin = time.Now()
	fmt.Println("Part 2:", evaluateEquations(inputLines, true), time.Since(begin))
}

func evaluateEquations(inputLines []string, includeConcat bool) int {
	var runningTotal int

	for _, line := range inputLines {
		colonIdx := strings.Index(line, ": ")
		targetValue, _ := strconv.Atoi(line[:colonIdx])
		parts := strings.Split(line[colonIdx+2:], " ")
		values := []int{}

		for _, part := range parts {
			value, _ := strconv.Atoi(part)
			values = append(values, value)
		}

		if process(targetValue, values[0], values[1:], includeConcat) {
			runningTotal += targetValue
		}
	}

	return runningTotal
}

func process(targetValue int, runningTotal int, values []int, includeConcat bool) bool {
	if runningTotal == targetValue && len(values) == 0 {
		return true
	} else if runningTotal > targetValue {
		return false
	}

	if len(values) == 0 {
		return false
	}

	if process(targetValue, runningTotal+values[0], values[1:], includeConcat) {
		return true
	}

	if process(targetValue, runningTotal*values[0], values[1:], includeConcat) {
		return true
	}

	if includeConcat {
		concatValue, _ := strconv.Atoi(strconv.Itoa(runningTotal) + strconv.Itoa(values[0]))
		if process(targetValue, concatValue, values[1:], includeConcat) {
			return true
		}
	}

	return false
}
