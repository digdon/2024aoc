package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	// Process the input
	wireInit := true
	wires := map[string]bool{}
	zSet := map[string]bool{}
	outputOpMap := map[string]Operation{}

	for _, line := range inputLines {
		if line == "" {
			// Empty line signals the end of the wire initialization and start of operations
			wireInit = false
			continue
		}

		if wireInit {
			// Wire initialization, found at the beginning of the input
			pos := strings.Index(line, ":")
			wire := line[:pos]
			value, _ := strconv.Atoi(line[pos+2:])
			wires[wire] = value == 1

			if wire[0] == 'z' {
				zSet[wire] = true
			}
		} else {
			// Operation definition
			var one, two, output, op string
			fmt.Sscanf(line, "%s %s %s -> %s", &one, &op, &two, &output)
			outputOpMap[output] = Operation{op, one, two, output}

			if output[0] == 'z' {
				zSet[output] = true
			}
		}
	}

	// Part 1 stuff

	// Find all of the 'z' wire values
	zList := []string{}
	maxZ := ""
	for z := range zSet {
		if z > maxZ {
			maxZ = z
		}
		zList = append(zList, z)
	}
	sort.Strings(zList)
	// sort.Sort(sort.Reverse(sort.StringSlice(zList)))

	result, bitString := calculateOutput(zList, wires, outputOpMap)
	fmt.Println(bitString)
	fmt.Printf("Part 1: %d (%s)\n", result, time.Since(begin))

	// Part 2 stuff
	begin = time.Now()
	failedGates := []string{}
	for _, gate := range outputOpMap {
		one, two, output := gate.inputOne, gate.inputTwo, gate.output

		// Rule #1 - if output is 'z', operation must be XOR, unless it's the last 'z' bit
		if output[0] == 'z' && output != maxZ && gate.op != "XOR" {
			failedGates = append(failedGates, output)
			continue
		}

		// Rule #2 - if output is not 'z' and the inputs are NOT 'x' and 'y', operation must NOT be XOR
		if output[0] != 'z' && one[0] != 'x' && one[0] != 'y' && two[0] != 'x' && two[0] != 'y' && gate.op == "XOR" {
			failedGates = append(failedGates, output)
			continue
		}

		// Rule #3 - if XOR with 'x and 'y' inputs, there must be an XOR gate with this output as the input
		if gate.op == "XOR" && (one[0] == 'x' || one[0] == 'y') && (two[0] == 'x' || two[0] == 'y') &&
			one != "x00" && two != "x00" && one != "y00" && two != "y00" {
			found := false
			for _, otherGate := range outputOpMap {
				if (otherGate.inputOne == output || otherGate.inputTwo == output) && otherGate.op == "XOR" {
					found = true
					break
				}
			}

			if !found {
				failedGates = append(failedGates, output)
				continue
			}
		}

		// Rule #4 - if AND with 'x' and 'y' inputs, there must be an OR gate with this output as the input
		if gate.op == "AND" && (one[0] == 'x' || one[0] == 'y') && (two[0] == 'x' || two[0] == 'y') &&
			one != "x00" && two != "x00" && one != "y00" && two != "y00" {
			found := false
			for _, otherGate := range outputOpMap {
				if (otherGate.inputOne == output || otherGate.inputTwo == output) && otherGate.op == "OR" {
					found = true
					break
				}
			}

			if !found {
				failedGates = append(failedGates, output)
				continue
			}
		}
	}

	sort.Strings(failedGates)
	fmt.Printf("Part 2: %s (%s)\n", strings.Join(failedGates, ","), time.Since(begin))
}

func calculateOutput(zList []string, wires map[string]bool, outputOpMap map[string]Operation) (int, string) {
	sort.Strings(zList)
	bits := getZvalues(zList, wires, outputOpMap)
	result := 0
	sb := strings.Builder{}

	for i := len(bits) - 1; i >= 0; i-- {
		result <<= 1
		if bits[i] {
			result++
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}
	bitString := sb.String()

	return result, bitString
}

func getZvalues(zList []string, wires map[string]bool, outputOpMap map[string]Operation) []bool {
	bitList := []bool{}

	for _, z := range zList {
		bitList = append(bitList, getWireValue(z, wires, outputOpMap))
	}

	return bitList
}

func getWireValue(wire string, wires map[string]bool, outputOpMap map[string]Operation) bool {
	if value, found := wires[wire]; found {
		return value
	}

	// No value yet - need to calculate it
	operation := outputOpMap[wire]
	inputOne := getWireValue(operation.inputOne, wires, outputOpMap)
	inputTwo := getWireValue(operation.inputTwo, wires, outputOpMap)

	// Perform the oprtation
	value := false
	switch operation.op {
	case "AND":
		value = inputOne && inputTwo
	case "OR":
		value = inputOne || inputTwo
	case "XOR":
		value = inputOne != inputTwo
	}

	return value
}

type Operation struct {
	op                 string
	inputOne, inputTwo string
	output             string
}
