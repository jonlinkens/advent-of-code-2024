package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jonlinkens/advent-of-code-2024/util"
)

func main() {
	data, error := os.ReadFile("01/input.txt")

	input := string(data)

	if error != nil {

		log.Fatal(error)
	}

	log.Println("Part 1:", part1(input))
	log.Println("Part 2:", part2(input))
}

func part1(input string) int {

	firstList, secondList := getTwoLists(input)

	sum := 0

	for range firstList {
		firstMin := slices.Min(firstList)
		secondMin := slices.Min(secondList)

		sum += util.AbsoluteDiff(firstMin, secondMin)

		firstList = util.RemoveFromSlice(firstList, firstMin)
		secondList = util.RemoveFromSlice(secondList, secondMin)

	}

	return sum
}

func part2(input string) int {

	firstList, secondList := getTwoLists(input)

	occurrences := make(map[int]int)

	for _, element := range secondList {
		occurrences[element] += 1
	}

	score := 0
	for _, element := range firstList {
		count := occurrences[element]

		score += element * count

	}

	return score
}

func getTwoLists(input string) ([]int, []int) {

	lines := strings.Split(input, "\n")

	var firstList []int
	var secondList []int

	for _, element := range lines {
		if len(element) > 0 {

			ids := strings.Split(element, "   ")

			first, _ := strconv.Atoi(ids[0])
			second, _ := strconv.Atoi(ids[1])

			firstList = append(firstList, first)
			secondList = append(secondList, second)
		}
	}

	if len(firstList) != len(secondList) {
		log.Fatal("Lists not equal length")
	}

	return firstList, secondList
}
