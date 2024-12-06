package main

import (
	"fmt"
	"os"
	"sort"
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
	var validLines []string
	for _, line := range lineSplits {
		if line != "" {
			validLines = append(validLines, line)
		}
	}

	left := make([]uint64, len(validLines))
	right := make([]uint64, len(validLines))
	for i, line := range validLines {
		split := strings.Fields(line)
		leftVal, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			fmt.Println("Error parsing left number:", err)
			return
		}
		rightVal, err := strconv.ParseUint(split[1], 10, 64)
		if err != nil {
			fmt.Println("Error parsing right number:", err)
			return
		}
		left[i] = leftVal
		right[i] = rightVal
	}

	// Sort the lists
	sortUint64s(left)
	sortUint64s(right)

	// Calculate the total distance
	var part1 uint64
	var part2 uint64
	addedLeft := make([]uint64, len(left))
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance > left[i] { // This handles the case where right[i] is larger than left[i]
			distance = right[i] - left[i]
		}
		part1 += distance
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				addedLeft[i] += left[i]
			}
		}
		part2 += addedLeft[i]
	}

	// Output the result
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

// Custom sort function for uint64 slices
func sortUint64s(a []uint64) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}
