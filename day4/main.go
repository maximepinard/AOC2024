package main

import (
	"fmt"
	"os"
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
	var part1 int
	var part2 uint64

	maxLength := 0
	for _, line := range lineSplits {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	matrix := make([][]rune, len(lineSplits))
	for i := range matrix {
		matrix[i] = make([]rune, maxLength)
	}

	for j, line := range lineSplits {
		for i, char := range line {
			matrix[j][i] = char
		}
	}

	part1 = countXMAS(matrix)
	part2 = uint64(countXMAS2(matrix))

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func countXMAS(matrix [][]rune) int {
	rows := len(matrix)
	cols := len(matrix[0])
	count := 0

	// Directions: right, left, down, up, down-right, down-left, up-right, up-left
	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if matrix[j][i] == 'X' {
				for _, dir := range directions {
					if checkDirection(matrix, j, i, dir[0], dir[1]) {
						count++
					}
				}
			}
		}
	}

	return count
}

func countXMAS2(matrix [][]rune) int {
	rows := len(matrix)
	cols := len(matrix[0])
	count := 0

	// Directions: right, left, down, up, down-right, down-left, up-right, up-left
	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if matrix[j][i] == 'A' {
				foundDirections := make(map[int]bool)
				for dirIndex, dir := range directions {
					if checkDirection2(matrix, j, i, dir[0], dir[1]) {
						foundDirections[dirIndex] = true
					}
				}
				if len(foundDirections) >= 2 {
					for dirIndex := range foundDirections {
						oppositeDirIndex := (dirIndex + 4) % 8
						if foundDirections[oppositeDirIndex] {
							count++
							break
						}
					}
				}
			}
		}
	}

	return count
}

func checkDirection(matrix [][]rune, row, col, dr, dc int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	// Check if the sequence "M", "A", "S" exists in the given direction
	if row+3*dr < 0 || row+3*dr >= rows || col+3*dc < 0 || col+3*dc >= cols {
		return false
	}

	return matrix[row+dr][col+dc] == 'M' &&
		matrix[row+2*dr][col+2*dc] == 'A' &&
		matrix[row+3*dr][col+3*dc] == 'S'
}

func checkDirection2(matrix [][]rune, row, col, dr, dc int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	// Check if the sequence "M", "A", "S" exists in the given direction
	if row-dr < 0 || row-dr >= rows || col-dc < 0 || col-dc >= cols ||
		row+dr < 0 || row+dr >= rows || col+dc < 0 || col+dc >= cols {
		return false
	}

	return matrix[row-dr][col-dc] == 'M' &&
		matrix[row+dr][col+dc] == 'S'
}
