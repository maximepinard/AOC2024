package main

import (
	"fmt"
	"os"
	"regexp"
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

	do := 1
	for _, line := range lineSplits {
		pattern := `mul\((\d+),(\d+)\)`
		re := regexp.MustCompile(pattern)

		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				num1, _ := strconv.ParseUint(match[1], 10, 64)
				num2, _ := strconv.ParseUint(match[2], 10, 64)
				part1 += num1 * num2
			}
		}

		substrings := strings.Split(line, "don't()")
		for i, substring := range substrings {
			if i > 0 {
				do = 0
			}
			substringsDo := strings.Split(substring, "do()")
			for j, substringDo := range substringsDo {
				// fmt.Println(substringDo)
				// fmt.Println("do()")
				if j > 0 {
					do = 1
				}
				if do == 1 {
					matches = re.FindAllStringSubmatch(substringDo, -1)
					for _, match := range matches {
						if len(match) == 3 {
							num1, _ := strconv.ParseUint(match[1], 10, 64)
							num2, _ := strconv.ParseUint(match[2], 10, 64)
							part2 += num1 * num2
							// fmt.Println(match[0])
						}
					}
				}
			}
			// fmt.Println("don't()")
		}
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}
