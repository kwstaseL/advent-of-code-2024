package main

import (
	"aoc_input/aoc_input"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Part 1: Sum all valid mul(X,Y) operations
func part1(input string) int {
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

// Part 2: Sum only the enabled mul(X,Y) operations based on do() and don't()
func part2(input string) int {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	enabled := true

	for _, match := range matches {
		instruction := strings.TrimSpace(match[0])

		if instruction == "do()" {
			enabled = true
		} else if instruction == "don't()" {
				enabled = false
		} else {
			if enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
			}
		}
	}

	return sum
}

func main() {
	input, err := aoc_input.GetAocInput(3, 2024)
	if err != nil {
		fmt.Printf("Error fetching input: %v\n", err)
		return
	}
	part1Result := part1(input)
	fmt.Println("Part 1 - Total sum of valid multiplications:", part1Result)

	part2Result := part2(input)
	fmt.Println("Part 2 - Total sum of valid enabled multiplications:", part2Result)
}
