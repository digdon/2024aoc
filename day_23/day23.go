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

	// Part 1
	fmt.Printf("Part 1: %d (%s)\n", part1(connectionsMap), time.Since(begin))

	// Part 2
	begin = time.Now()
	fmt.Printf("Part 2: %s (%s)\n", strings.Join(part2(connectionsMap), ","), time.Since(begin))
}

func part1(connectionsMap map[string]map[string]bool) int {
	// For each node, look at each neighbour and try to find a third node (based on the second)
	// that connects to the first. If a triangle is found, build a set
	triangles := map[string][]string{}

	for one, cmap := range connectionsMap {
		// We're only interested in triangles where at least one node starts with 't'.
		// Adding this check reduces execution time by about 20x
		if one[0] != 't' {
			continue
		}

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

	return tCount
}

func part2(connectionsMap map[string]map[string]bool) []string {
	// Prep B-K maps
	r := map[string]bool{}
	p := map[string]bool{}
	x := map[string]bool{}

	for node := range connectionsMap {
		p[node] = true
	}

	var cliques [][]string
	bronKerbosch(connectionsMap, r, p, x, &cliques)

	var largestClique []string
	for _, clique := range cliques {
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}

	sort.Strings(largestClique)

	return largestClique
}

func bronKerbosch(connectionsMap map[string]map[string]bool, r, p, x map[string]bool, cliques *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		clique := make([]string, 0, len(r))
		for node := range r {
			clique = append(clique, node)
		}
		*cliques = append(*cliques, clique)
		return
	}

	for v := range p {
		newR := copySet(r)
		newR[v] = true
		newP := intersect(p, connectionsMap[v])
		newX := intersect(x, connectionsMap[v])

		bronKerbosch(connectionsMap, newR, newP, newX, cliques)

		delete(p, v)
		x[v] = true
	}
}

func copySet(s map[string]bool) map[string]bool {
	c := make(map[string]bool, len(s))
	for k, v := range s {
		c[k] = v
	}
	return c
}

func intersect(a, b map[string]bool) map[string]bool {
	c := map[string]bool{}
	for k := range a {
		if b[k] {
			c[k] = true
		}
	}
	return c
}
