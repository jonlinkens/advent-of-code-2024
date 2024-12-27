package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, error := os.ReadFile("05/input.txt")

	input := string(data)

	if error != nil {
		log.Fatal(error)
	}

	log.Println("Part 1:", part1(input))
	log.Println("Part 2:", part2(input))
}

type Rule struct {
	a int
	b int
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")

	rulesString := parts[0]
	rules := getRules(rulesString)

	updatesString := parts[1]
	updates := getUpdates(updatesString)

	sum := 0
	for _, update := range updates {
		if isValid(update, rules) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")

	rulesString := parts[0]
	rules := getRules(rulesString)

	updatesString := parts[1]
	updates := getUpdates(updatesString)

	sum := 0

	for _, update := range updates {
		if !isValid(update, rules) {
			sortedUpdate := reorder(update, rules)
			middleIndex := len(sortedUpdate) / 2
			sum += sortedUpdate[middleIndex]
		}
	}

	return sum
}

func getRules(rawString string) []Rule {
	rules := []Rule{}

	for _, rule := range strings.Split(rawString, "\n") {
		parts := strings.Split(rule, "|")
		ruleA, _ := strconv.Atoi(parts[0])
		ruleB, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{ruleA, ruleB})
	}

	return rules
}

func getUpdates(rawString string) [][]int {
	var updates [][]int

	for _, update := range strings.Split(rawString, "\n") {
		var u []int
		for _, _num := range strings.Split(update, ",") {
			num, _ := strconv.Atoi(_num)
			u = append(u, int(num))
		}
		updates = append(updates, u)
	}

	return updates
}

func compare(a, b int, rules []Rule) int {
	for _, rule := range rules {
		if a == rule.a && b == rule.b {
			return -1
		} else if b == rule.a && a == rule.b {
			return 1
		}
	}
	return 0
}

func reorder(update []int, rules []Rule) []int {
	sorted := make([]int, len(update))
	copy(sorted, update)

	sort.Slice(sorted, func(i, j int) bool {
		return compare(sorted[i], sorted[j], rules) < 0
	})

	return sorted
}

func isValid(update []int, rules []Rule) bool {
	for i := 0; i < len(update)-1; i++ {
		if compare(update[i], update[i+1], rules) > 0 {
			return false
		}
	}
	return true
}
