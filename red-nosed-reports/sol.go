package main

import (
	"aoc_input/aoc_input"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func checkReport(report []int) bool {
	if len(report) < 2 {
		return true
	}

	increasing := report[1] > report[0]

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		// Validate the difference based on the current direction.
		if (increasing && (diff < 1 || diff > 3)) || (!increasing && (diff < -3 || diff > -1)) {
			return false
		}

		// Handle transitions between increasing and decreasing.
		if increasing && diff < 0 {
			increasing = false
		} else if !increasing && diff > 0 {
			return false
		}
	}

	return true
}


func countSafeReports(reports string) int {
	safeCount := 0
	lines := strings.Split(reports, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		reportParts := strings.Split(line, " ")
		report := make([]int, len(reportParts))
		isValid := true

		for i, part := range reportParts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Printf("Warning: failed to parse number %s: %v", part, err)
				isValid = false
				break
			}
			report[i] = num
		}

		if !isValid {
			continue // Skip reports with parsing errors.
		}

		if checkReport(report) {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	reports, err := aoc_input.GetAocInput(2, 2024)
	if err != nil {
		log.Fatalf("Failed to get AoC input: %v", err)
	}
	safeReports := countSafeReports(reports)
	fmt.Println("Number of safe reports:", safeReports)
}
