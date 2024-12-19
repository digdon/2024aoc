package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

	output := runProgram(program, regA, regB, regC)
	outputStrVals := []string{}
	for _, value := range output {
		outputStrVals = append(outputStrVals, strconv.Itoa(value))
	}
	outputString := strings.Join(outputStrVals, ",")

	fmt.Println("Part 1:", outputString)
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
			div := int(math.Pow(float64(2), float64(comboOperandValue(operand, regA, regB, regC))))
			regA = num / div

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
			div := int(math.Pow(float64(2), float64(comboOperandValue(operand, regA, regB, regC))))
			regB = num / div

		case 7: // cdv
			num := regA
			div := int(math.Pow(float64(2), float64(comboOperandValue(operand, regA, regB, regC))))
			regC = num / div
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
