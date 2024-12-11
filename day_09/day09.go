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

	input := inputLines[0]
	// input := "2333133121414131402"
	// input := "12345"

	begin := time.Now()
	fmt.Println("Part 1:", defrag(input, false), time.Since(begin))
	begin = time.Now()
	fmt.Println("Part 2:", defrag(input, true), time.Since(begin))
}

func generateBlocks(input string) (*Block, *Block) {
	var head, tail *Block
	var totalBlocks int

	for empty, id, i := false, 0, 0; i < len(input); empty, i = !empty, i+1 {
		count, _ := strconv.Atoi(string(input[i]))
		var block *Block

		if empty {
			if count == 0 {
				continue
			}

			block = &Block{id: -1, size: count, start: totalBlocks}
		} else {
			block = &Block{id: id, size: count, start: totalBlocks}
			id++
		}

		if head == nil {
			head = block
		} else {
			tail.next = block
		}

		block.prev = tail
		tail = block

		totalBlocks += count
	}

	return head, tail
}

func defrag(input string, moveFull bool) int {
	head, tail := generateBlocks(input)
	full, empty := tail, head

	for {
		// Find the next non-empty block that hasn't been moved
		for ; full != nil && (full.id == -1 || full.moved); full = full.prev {
		}

		if full == nil {
			// Nothing left to move
			break
		}

		if moveFull {
			// Try to find an empty block that can hold the entirety of the current non-empty block, starting from the head
			for empty = head; empty != nil && (empty.id != -1 || empty.size < full.size); empty = empty.next {
			}
		} else {
			// Look for the next empty block that has any non-zero capacity, from the previous empty block
			for ; empty != nil && (empty.id != -1 || empty.size == 0); empty = empty.next {
			}
		}

		if empty == nil || full.start < empty.start {
			if moveFull {
				// No empty block big enough for this file - go on to the next non-empty block
				full = full.prev
				continue
			} else {
				// Can't move non-empty block contents further into the list
				break
			}
		}

		moveCount := min(full.size, empty.size)

		// Found a spot to move the non-empty block to
		if empty.size == full.size {
			// Blocks are the same size, so just swap contents
			empty.id, full.id = full.id, empty.id
			empty.moved = true
			full = full.prev
		} else {
			// We're only partially filling the empty block, so we need to split it
			newBlock := &Block{id: full.id, size: moveCount, start: empty.start, moved: true}
			empty.prev.next = newBlock
			newBlock.prev = empty.prev
			newBlock.next = empty
			empty.prev = newBlock

			// Update the remaining empty block
			empty.start += moveCount
			empty.size -= moveCount

			// Adjust the non-empty block, as necessary
			full.size -= moveCount

			if full.size == 0 {
				full.id = -1
			}
		}
	}

	return checksum(head)
}

func checksum(blocks *Block) int {
	var checksum int
	for block := blocks; block != nil; block = block.next {
		if block.id == -1 {
			continue
		}

		for i := 0; i < block.size; i++ {
			checksum += ((block.start + i) * block.id)
		}
	}

	return checksum
}

func debug(blocks *Block) {
	for count, block := 0, blocks; block != nil; count, block = count+1, block.next {
		if block.id == -1 {
			fmt.Printf("empty: start=%d, size=%d\n", block.start, block.size)
		} else {
			fmt.Printf("%-5d: start=%d, size=%d\n", block.id, block.start, block.size)
		}
	}
}

func display(blocks *Block) {
	for block := blocks; block != nil; block = block.next {
		var char string

		if block.id == -1 {
			char = "."
		} else {
			char = strconv.Itoa(block.id)
		}

		fmt.Printf("%s", strings.Repeat(char, block.size))
	}

	fmt.Println()
}

type Block struct {
	id         int
	moved      bool
	size       int
	start      int
	prev, next *Block
}
