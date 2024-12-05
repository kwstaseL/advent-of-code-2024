package main

import (
	"aoc_input/aoc_input"
	"fmt"
	"regexp"
	"strconv"
)

/*
	Alternatively, do an exhaustive search for all possible multiplications and sum them up.
*/
func sumMultiplications(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range matches {
		if len(match) == 3 {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			sum += x * y
		}
	}

	return sum
}

func main() {
	input, _ := aoc_input.GetAocInput(3, 2024)
	result := sumMultiplications(input)
	fmt.Println("Total sum of valid multiplications:", result)
}
