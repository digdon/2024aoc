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

	regA, regB, regC := 0, 0, 0
	program := []int{}

	for _, line := range inputLines {
		if strings.HasPrefix(line, "Register A:") {
			fmt.Sscanf(line, "Register A: %d", &regA)
		} else if strings.HasPrefix(line, "Register B:") {
			fmt.Sscanf(line, "Register B: %d", &regB)
		} else if strings.HasPrefix(line, "Register C:") {
			fmt.Sscanf(line, "Register C: %d", &regC)
		} else if line == "" {
			continue
		} else {
			parts := strings.Split(strings.Split(line, " ")[1], ",")

			for _, part := range parts {
				value, _ := strconv.Atoi(part)
				program = append(program, value)
			}
		}
	}

	begin := time.Now()
	output := runProgram(program, regA, regB, regC)
	outputStrVals := []string{}
	for _, value := range output {
		outputStrVals = append(outputStrVals, strconv.Itoa(value))
	}
	outputString := strings.Join(outputStrVals, ",")

	fmt.Printf("Part 1: %s (%s)\n", outputString, time.Since(begin))

	// Part 2 stuff
	begin = time.Now()
	result := solve(program, 0, 0, regB, regC)
	fmt.Printf("Part 2: %d (%s)\n", result, time.Since(begin))
	output = runProgram(program, result, regB, regC)
	fmt.Println(output)
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, value := range a {
		if value != b[i] {
			return false
		}
	}

	return true
}

func solve(program []int, pos, regA, regB, regC int) int {
	output := runProgram(program, regA, regB, regC)

	if slicesEqual(output, program) {
		return regA
	}

	if pos == 0 || slicesEqual(output, program[len(program)-pos:]) {
		for ni := 0; ni < 8; ni++ {
			if na := solve(program, pos+1, 8*regA+ni, regB, regC); na > 0 {
				return na
			}
		}
	}

	return 0
}

func runProgram(program []int, regA, regB, regC int) []int {
	pc := 0
	output := []int{}

	// Run the program
	for pc < len(program) {
		instruction := program[pc]
		operand := program[pc+1]
		pc += 2

		switch instruction {
		case 0: // adv
			num := regA
			regA = num >> comboOperandValue(operand, regA, regB, regC)

		case 1: // bxl
			regB ^= operand

		case 2: // bst
			regB = comboOperandValue(operand, regA, regB, regC) % 8

		case 3: // jnz
			if regA != 0 {
				pc = operand
			}

		case 4: // bxc
			regB ^= regC

		case 5: // out
			output = append(output, comboOperandValue(operand, regA, regB, regC)%8)

		case 6: // bdv
			num := regA
			regB = num >> comboOperandValue(operand, regA, regB, regC)

		case 7: // cdv
			num := regA
			regC = num >> comboOperandValue(operand, regA, regB, regC)
		}
	}

	return output
}

func comboOperandValue(operand, regA, regB, regC int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	default:
		return -1
	}
}

var instMap = map[int]string{
	0: "adv",
	1: "bxl",
	2: "bst",
	3: "jnz",
	4: "bxc",
	5: "out",
	6: "bdv",
	7: "cdv",
}
