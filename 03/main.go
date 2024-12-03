package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, error := os.ReadFile("03/input.txt")

	input := string(data)

	if error != nil {

		log.Fatal(error)
	}

	log.Println("Part 1:", part1(input))

	log.Println("Part 2:", part2(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {

			sum += convNumsAndMultiply(match[1], match[2])
		}
	}

	return sum
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	sum := 0
	enabled := true
	for _, line := range lines {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch {
			case match[0] == "do()":
				enabled = true
			case match[0] == "don't()":
				enabled = false
			case enabled && match[1] != "" && match[2] != "":
				sum += convNumsAndMultiply(match[1], match[2])
			}
		}
	}

	return sum
}

func convNumsAndMultiply(a, b string) int {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	return x * y
}
