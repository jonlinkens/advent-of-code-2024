package main

import (
	"log"
	"os"
	"strings"
)

const WORD = "XMAS"

func main() {
	data, err := os.ReadFile("04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	log.Println("Part 1:", part1(input))
	log.Println("Part 2:", part2(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	count := 0

	for y, line := range lines {
		if len(line) == 0 {
			continue
		}

		for x, letter := range line {
			if byte(letter) == WORD[0] {
				if checkEast(lines, x, y) {
					count++
				}
				if checkWest(lines, x, y) {
					count++
				}
				if checkNorth(lines, x, y) {
					count++
				}
				if checkSouth(lines, x, y) {
					count++
				}
				if checkNorthEast(lines, x, y) {
					count++
				}
				if checkNorthWest(lines, x, y) {
					count++
				}
				if checkSouthEast(lines, x, y) {
					count++
				}
				if checkSouthWest(lines, x, y) {
					count++
				}
			}
		}
	}

	return count
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	count := 0

	for y, line := range lines {
		if len(line) == 0 || y == 0 || y >= len(lines)-1 {
			continue
		}
		prevLine := lines[y-1]
		nextLine := lines[y+1]

		for x, letter := range line {
			if x == 0 || x >= len(line)-1 ||
				x >= len(prevLine)-1 || x >= len(nextLine)-1 {
				continue
			}
			if letter == 'A' && checkXPattern(prevLine, nextLine, x) {
				count++
			}
		}
	}
	return count
}

func checkXPattern(prevLine, nextLine string, x int) bool {
	tlbr := (prevLine[x-1] == 'M' && nextLine[x+1] == 'S') ||
		(prevLine[x-1] == 'S' && nextLine[x+1] == 'M')

	trbl := (prevLine[x+1] == 'M' && nextLine[x-1] == 'S') ||
		(prevLine[x+1] == 'S' && nextLine[x-1] == 'M')

	return tlbr && trbl
}

func checkEast(lines []string, x int, y int) bool {
	if x+len(WORD) > len(lines[y]) {
		return false
	}

	for i := range WORD {
		if lines[y][x+i] != WORD[i] {
			return false
		}
	}

	return true
}

func checkWest(lines []string, x int, y int) bool {
	if x-len(WORD)+1 < 0 {
		return false
	}

	for i := range WORD {
		if lines[y][x-i] != WORD[i] {
			return false
		}
	}

	return true
}

func checkNorth(lines []string, x int, y int) bool {
	if y-len(WORD)+1 < 0 {
		return false
	}

	for i := range WORD {
		if len(lines[y-i]) <= x || lines[y-i][x] != WORD[i] {
			return false
		}
	}

	return true
}

func checkSouth(lines []string, x int, y int) bool {
	if y+len(WORD) > len(lines) {
		return false
	}

	for i := range WORD {
		if len(lines[y+i]) <= x || lines[y+i][x] != WORD[i] {
			return false
		}
	}

	return true
}

func checkNorthEast(lines []string, x int, y int) bool {
	offset := len(WORD) - 1
	if y-offset < 0 || x+offset >= len(lines[y]) {
		return false
	}

	for i := range WORD {
		if len(lines[y-i]) <= x+i || lines[y-i][x+i] != WORD[i] {
			return false
		}
	}

	return true
}

func checkNorthWest(lines []string, x int, y int) bool {
	offset := len(WORD) - 1
	if y-offset < 0 || x-offset < 0 {
		return false
	}

	for i := range WORD {
		if len(lines[y-i]) <= x-i || lines[y-i][x-i] != WORD[i] {
			return false
		}
	}

	return true
}

func checkSouthEast(lines []string, x int, y int) bool {
	offset := len(WORD) - 1
	if y+offset >= len(lines) || x+offset >= len(lines[y]) {
		return false
	}

	for i := range WORD {
		if len(lines[y+i]) <= x+i || lines[y+i][x+i] != WORD[i] {
			return false
		}
	}

	return true
}

func checkSouthWest(lines []string, x int, y int) bool {
	offset := len(WORD) - 1
	if y+offset >= len(lines) || x-offset < 0 {
		return false
	}

	for i := range WORD {
		if len(lines[y+i]) <= x-i || lines[y+i][x-i] != WORD[i] {
			return false
		}
	}

	return true
}
