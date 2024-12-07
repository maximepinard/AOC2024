package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileString := strings.ReplaceAll(string(fileContent), "\r", "")

	lineSplits := strings.Split(fileString, "\n")
	var part1 uint64
	var part2 uint64

	for _, line := range lineSplits {
		if line == "" {
			continue
		}

		split := strings.Fields(line)
		result := isValidSequence(split)
		if result == 1 {
			part1++
			part2++
		} else {
			// Check if removing any single level makes the report safe
			if isSafeWithDampener(split) {
				part2++
			}
		}
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func isValidSequence(split []string) int {
	mode := "no"
	prevVal := 0
	result := 1

	for j, s := range split {
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return -1
		}

		if j > 0 {
			if mode == "no" {
				if val > prevVal {
					mode = ">"
				} else if val < prevVal {
					mode = "<"
				} else {
					result--
				}
			} else if (mode == ">" && val <= prevVal) || (mode == "<" && val >= prevVal) {
				result--
			}

			if abs(val-prevVal) > 3 {
				result--
			}
		}

		prevVal = val
	}

	return result
}

func isSafeWithDampener(split []string) bool {
	for i := range split {
		// Remove the i-th element and check if the remaining sequence is valid
		remaining := append([]string{}, split[:i]...)
		remaining = append(remaining, split[i+1:]...)
		if isValidSequence(remaining) == 1 {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
