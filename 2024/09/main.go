package main

import (
	"fmt"

	"github.com/nathanheffley/advent-of-code/input"
)

func main() {
	line := input.ReadInputFileToLines("input.txt")[0]

	nums := make([]int, 0)
	for _, r := range line {
		nums = append(nums, int(r-'0'))
	}

	blocks := make([]int, 0)
	for i, n := range nums {
		if i%2 == 0 {
			for x := 0; x < n; x++ {
				blocks = append(blocks, i/2)
			}
		} else {
			for x := 0; x < n; x++ {
				blocks = append(blocks, -1)
			}
		}
	}

	wholeBlocks := make([]int, len(blocks))
	copy(wholeBlocks, blocks)

	// Sort with fragmentation.
	fragmentedBlocks := make([]int, len(blocks))
	copy(fragmentedBlocks, blocks)
	filledTo := 0
	for i := len(blocks) - 1; i >= filledTo; i-- {
		if blocks[i] == -1 {
			continue
		}
		for j := 0; j < i; j++ {
			if blocks[j] == -1 {
				blocks[j] = blocks[i]
				blocks[i] = -1
				filledTo = j
				break
			}
		}
	}
	fragmentedChecksum := 0
	for i, num := range blocks {
		if num != -1 {
			fragmentedChecksum += i * num
		}
	}
	fmt.Printf("Fragmented checksum: %d\n", fragmentedChecksum)

	// Sort with no fragmentation.
	movedFiles := make(map[int]bool, 0)
	movedFiles[-1] = true
	for i := len(wholeBlocks) - 1; i >= 0; i-- {
		if _, ok := movedFiles[wholeBlocks[i]]; ok {
			continue
		}
		movedFiles[wholeBlocks[i]] = true
		blockStart := i
		for ii := i - 1; ii >= 0; ii-- {
			if wholeBlocks[ii] == wholeBlocks[i] {
				blockStart = ii
			} else {
				break
			}
		}

		for j := 0; j < blockStart; j++ {
			if wholeBlocks[j] == -1 {
				currentEmptyBlockLength := 0
				for jj := j; jj < blockStart; jj++ {
					if wholeBlocks[jj] == -1 {
						currentEmptyBlockLength++
						if currentEmptyBlockLength == i-blockStart+1 {
							break
						}
					} else {
						currentEmptyBlockLength = 0
						break
					}
				}
				if currentEmptyBlockLength == i-blockStart+1 {
					// Move the block.
					offset := 0
					for originalIndex := blockStart; originalIndex <= i; originalIndex++ {
						wholeBlocks[j+offset] = wholeBlocks[originalIndex]
						wholeBlocks[originalIndex] = -1
						offset++
					}
					break
				}
			}
		}

		i = blockStart
	}
	wholeChecksum := 0
	for i, num := range wholeBlocks {
		if num != -1 {
			wholeChecksum += i * num
		}
	}
	fmt.Printf("Whole checksum: %d\n", wholeChecksum)
}
