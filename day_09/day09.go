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
	fmt.Println("Part 1:", partOne(input), time.Since(begin))
	begin = time.Now()
	fmt.Println("Part 2:", partTwo(input), time.Since(begin))
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

func partOne(input string) int {
	head, tail := generateBlocks(input)
	full, empty := tail, head

	for full != nil && empty != nil {
		// Find the next non-empty block, from the end
		for full != nil && full.id == -1 {
			full = full.prev
		}

		// Find the next empty block, with capacity, from the beginning
		for empty != nil && (empty.id != -1 || empty.size == 0) {
			empty = empty.next
		}

		if full == nil || empty == nil || full.start < empty.start {
			break
		}

		// Calculate the number of items in the non-empty block that can be moved
		moveCount := min(full.size, empty.size)

		// Process the empty block
		if moveCount == empty.size {
			// We're completely filling the empty block, so just update the id
			empty.id = full.id
			empty = empty.next
		} else {
			// We're only partially filling the empty block, so we need to split it
			newBlock := &Block{id: full.id, size: moveCount, start: empty.start}
			empty.prev.next = newBlock
			newBlock.prev = empty.prev
			newBlock.next = empty
			empty.prev = newBlock

			empty.start += moveCount
			empty.size -= moveCount
		}

		// Process the non-empty block
		if moveCount == full.size {
			// We've moved all of the items from the non-empty block, so we can drop it
			full = full.prev
			full.next = nil
			tail = full
		} else {
			// We only moved some items, so change the current size
			full.size -= moveCount
		}
	}

	return checksum(head)
}

func partTwo(input string) int {
	head, tail := generateBlocks(input)
	full, empty := tail, head

	for {
		// Find the next non-empty block that hasn't been moved
		for full != nil && (full.id == -1 || full.moved) {
			full = full.prev
		}

		if full == nil {
			// Nothing left to move
			break
		}

		// Try to find an empty block that can hold the current non-empty block, starting from the head
		for empty = head; empty != nil && (empty.id != -1 || empty.size < full.size); empty = empty.next {
		}

		if empty == nil || full.start < empty.start {
			// No empty block big enough for this file - go on to the next non-empty block
			full = full.prev
			continue
		}

		// Found a spot to move the non-empty block to
		if empty.size == full.size {
			// Blocks are the size size, so just swap contents
			empty.id, full.id = full.id, empty.id
			empty.moved = true
			full = full.prev
		} else {
			// We're only partially filling the empty block, so we need to split it
			newBlock := &Block{id: full.id, size: full.size, start: empty.start, moved: true}
			empty.prev.next = newBlock
			newBlock.prev = empty.prev
			newBlock.next = empty
			empty.prev = newBlock

			// Update the remaining empty block and mark the non-empty block as now empty
			empty.start += full.size
			empty.size -= full.size
			full.id = -1
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
