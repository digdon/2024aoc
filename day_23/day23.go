package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	connectionsMap := map[string]map[string]bool{}

	// Start by building a paired connection map
	for _, line := range inputLines {
		parts := strings.Split(line, "-")
		one := parts[0]
		two := parts[1]

		cmap, ok := connectionsMap[one]
		if !ok {
			cmap = map[string]bool{}
			connectionsMap[one] = cmap
		}
		cmap[two] = true

		cmap, ok = connectionsMap[two]
		if !ok {
			cmap = map[string]bool{}
			connectionsMap[two] = cmap
		}
		cmap[one] = true
	}

	// For each node, look at each neighbour and try to find a third node (based on the second)
	// that connects to the first. If a triangle is found, build a set
	triangles := map[string][]string{}

	for one, cmap := range connectionsMap {
		for two := range cmap {
			for three := range connectionsMap[two] {
				if connectionsMap[three][one] {
					nodes := []string{one, two, three}
					sort.Slice(nodes, func(i, j int) bool {
						return nodes[i] < nodes[j]
					})
					triangles[strings.Join(nodes, ",")] = nodes
				}
			}
		}
	}

	tCount := 0
	for _, nodes := range triangles {
		for _, node := range nodes {
			if node[0] == 't' {
				// fmt.Println(triangle)
				tCount++
				break
			}
		}
	}

	fmt.Printf("Part 1: %d (%s)\n", tCount, time.Since(begin))
}
