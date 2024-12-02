package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jonlinkens/advent-of-code-2024/util"
)

func main() {
	data, error := os.ReadFile("02/input.txt")

	input := string(data)

	if error != nil {

		log.Fatal(error)
	}

	log.Println("Part 1:", part1(input))

	log.Println("Part 2:", part2(input))
}

func part1(input string) int {
	lines := getLinesOfInts(input)

	safeReports := 0

	for _, line := range lines {
		if isReportSafe(line) {
			safeReports++
		}

	}

	return safeReports
}

func part2(input string) int {
	lines := getLinesOfInts(input)

	safeReports := 0

	for _, line := range lines {
		if isReportSafeWithDampening(line) {
			safeReports++
		}

	}

	return safeReports
}

func isReportSafeWithDampening(report []int) bool {
	if isReportSafe(report) {
		return true
	}

	for i := range report {
		newSlice := util.RemoveFromSliceOrdered(report, i)
		if isReportSafe(newSlice) {
			return true
		}
	}
	return false
}

func isReportSafe(report []int) bool {
	if len(report) == 0 {
		return false
	}

	isIncreasing := report[0] < report[1]

	for i := 0; i < len(report)-1; i++ {
		if (isIncreasing && report[i] > report[i+1]) || (!isIncreasing && report[i] < report[i+1]) {
			return false
		}

		diff := util.AbsoluteDiff(report[i], report[i+1])
		if diff > 3 || diff < 1 {
			return false
		}
	}

	return true
}

func getLinesOfInts(input string) [][]int {
	lines := strings.Split(input, "\n")
	result := make([][]int, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		ints := make([]int, len(fields))
		for j, field := range fields {
			num, _ := strconv.Atoi(field)
			ints[j] = num
		}
		result[i] = ints
	}
	return result
}
