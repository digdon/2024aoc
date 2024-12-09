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

		root := Node{value: values[0], operation: ADD, result: values[0], parent: nil, children: nil}

		if process(targetValue, root, values[1:], includeConcat) {
			runningTotal += targetValue
		}
	}

	return runningTotal
}

func process(targetValue int, node Node, values []int, includeConcat bool) bool {
	if node.result == targetValue {
		// displayResult(&node, targetValue)
		return true
	} else if node.result > targetValue {
		return false
	}

	if len(values) == 0 {
		return false
	}

	node.children = append(node.children, Node{value: values[0], operation: ADD, result: node.result + values[0], parent: &node, children: nil})
	node.children = append(node.children, Node{value: values[0], operation: MUL, result: node.result * values[0], parent: &node, children: nil})

	if includeConcat {
		concatValue, _ := strconv.Atoi(strconv.Itoa(node.result) + strconv.Itoa(values[0]))
		node.children = append(node.children, Node{value: values[0], operation: CONCAT, result: concatValue, parent: &node, children: nil})
	}

	for _, child := range node.children {
		if process(targetValue, child, values[1:], includeConcat) {
			return true
		}
	}

	return false
}

func displayResult(node *Node, result int) {
	sb := ""
	for node != nil {
		sb = strconv.Itoa(node.value) + sb

		if node.parent != nil {
			switch node.operation {
			case ADD:
				sb = " + " + sb
			case MUL:
				sb = " * " + sb
			case CONCAT:
				sb = " || " + sb
			}
		}

		node = node.parent
	}

	fmt.Println(result, ": ", sb)
}

type Node struct {
	value     int
	operation int
	result    int
	parent    *Node
	children  []Node
}

const (
	ADD = iota
	MUL
	CONCAT
)
