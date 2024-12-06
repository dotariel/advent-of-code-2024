package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents, _ := os.ReadFile(("input.txt"))
	input := string(contents)

	part1(input)
	part2(input)
}

func part1(input string) {
	fmt.Printf("Part 1: %v\n", GetSum(input))
}

func part2(input string) {
	fmt.Printf("Part 2: %v\n", GetSumWithConditions(input))
}

func GetSum(input string) int {
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)

	sum := 0

	for _, match := range r.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		sum += mul(x, y)
	}

	return sum
}

func GetSumWithConditions(input string) int {
	const (
		DO   = "do()"
		DONT = "don't()"
	)

	sum := 0

	for {
		dont := strings.Index(input, DONT)
		if dont == -1 {
			sum += GetSum(input)
			break
		}

		sum += GetSum(input[0:dont])
		input = input[dont+len(DONT):]

		do := strings.Index(input, DO)
		if do == -1 {
			break
		}

		input = input[do+len(DO):]
	}

	return sum
}

func mul(x, y int) int {
	return x * y
}
