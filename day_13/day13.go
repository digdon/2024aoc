package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	systems := []System{}
	var a, b XYvalues

	for _, line := range inputLines {
		var x, y int
		if strings.HasPrefix(line, "Button A:") {
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			a = XYvalues{x, y}
		} else if strings.HasPrefix(line, "Button B:") {
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			b = XYvalues{x, y}
		} else if strings.HasPrefix(line, "Prize:") {
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			systems = append(systems, System{Equation{a.x, b.x, x}, Equation{a.y, b.y, y}})
		} else {
			continue
		}
	}

	var part1Tokens, part2Tokens int

	for _, system := range systems {
		// fmt.Println(system, solveSystem(system))
		part1Tokens += solveSystem(system)

		system.one.result += 10000000000000
		system.two.result += 10000000000000
		part2Tokens += solveSystem(system)
	}

	fmt.Println("Part 1:", part1Tokens)
	fmt.Println("Part 2:", part2Tokens)
	fmt.Println(time.Since(begin))
}

func solveSystem(system System) int {
	// Start by solving for 'a'

	lcmB := lcm(system.one.b, system.two.b)

	oneMult, twoMult := lcmB/system.one.b, lcmB/system.two.b
	oneA, oneResult := system.one.a*oneMult, system.one.result*oneMult
	twoA, twoResult := system.two.a*twoMult, system.two.result*twoMult
	a := (oneResult - twoResult) / (oneA - twoA)

	if (oneResult-twoResult)%(oneA-twoA) != 0 {
		return 0
	}

	if a < 0 {
		a = -a
	}

	// Plug 'a into equation one to solve for 'b'
	b := (system.one.result - system.one.a*a) / system.one.b

	if (system.one.result-system.one.a*a)%system.one.b != 0 {
		return 0
	}

	if b < 0 {
		b = -b
	}

	return a*3 + b
}

func lcm(x, y int) int {
	if y == 0 {
		return 0
	}

	return (x * y) / gcd(x, y)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

type System struct {
	one, two Equation
}

type Equation struct {
	a, b, result int
}

type XYvalues struct {
	x, y int
}
