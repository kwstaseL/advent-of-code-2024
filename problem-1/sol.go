package main

import (
	"aoc2024/aoc_input"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

/*
	Alternatively, one could also use a heap here
*/
func part1(inputData string) {
	lines := strings.Split(strings.TrimSpace(inputData), "\n")

	var leftList, rightList []int

	for _, line := range lines {
		var left, right int
		_, err := fmt.Sscanf(line, "%d %d", &left, &right)
		if err != nil {
			log.Fatalf("Failed to parse line: %s", line)
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for len(leftList) > 0 && len(rightList) > 0 {
		leftMin := leftList[0]
		rightMin := rightList[0]

		leftList = leftList[1:]
		rightList = rightList[1:]

		totalDistance += int(math.Abs(float64(leftMin - rightMin)))
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
}

/*
	Bruteforcing is O(n*m) where n is the length of leftList and m is the length of rightList.
	Alternatively, we could sort the rightList and perform a binary search, in which case the complexity would be O(n*log(m))
*/

func part2(inputData string) {
	var leftList, rightList []int

	for _, line := range strings.Split(strings.TrimSpace(inputData), "\n") {
		var left, right int
		_, err := fmt.Sscanf(line, "%d %d", &left, &right)
		if err != nil {
			log.Fatalf("Failed to parse line: %s", line)
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	rightCount := make(map[int]int)
	for _, right := range rightList {
		rightCount[right]++
	}
	similarityScore := 0
	for _, left := range leftList {
		similarityScore += left * rightCount[left]
	}
	fmt.Printf("Similarity score: %d\n", similarityScore)

}

func main() {
	inputData, err := aoc_input.GetAocInput(1, 2024)
	if err != nil {
		log.Fatalf("Failed to get AoC input: %v", err)
	}
	part1(inputData)
	part2(inputData)
}
