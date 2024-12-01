package main

import (
	"log"
	"os"
)

func main() {
	data, error := os.ReadFile("test.txt")

	input := string(data)

	if error != nil {

		log.Fatal(error)
	}

	log.Println("Part 1:", part1(input))

	log.Println("Part 2:", part2(input))
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
