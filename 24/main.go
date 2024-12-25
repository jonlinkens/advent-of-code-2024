package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, error := os.ReadFile("24/input.txt")

	input := string(data)

	if error != nil {
		log.Fatal(error)
	}

	log.Println("Part 1:", part1(input))

}

type Point struct {
	i, j int
}

type Set map[Point]bool

func part1(input string) int {

	patternStrings := strings.Split(input, "\n\n")
	patterns := make([]Set, 0, len(patternStrings))

	for _, patternString := range patternStrings {
		lines := strings.Split(patternString, "\n")
		pattern := make(Set)

		for i, line := range lines {
			for j, char := range line {
				if char == '#' {
					pattern[Point{i, j}] = true
				}
			}
		}
		patterns = append(patterns, pattern)
	}

	count := 0
	for i := 0; i < len(patterns); i++ {
		for j := i + 1; j < len(patterns); j++ {
			if !patterns[i].intersection(patterns[j]) {
				count++
			}
		}
	}

	return count
}

func (s Set) intersection(other Set) bool {
	for point := range s {
		if other[point] {
			return true
		}
	}
	return false
}
