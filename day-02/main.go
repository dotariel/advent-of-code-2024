package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	safeReports := 0

	for _, report := range makeReports() {
		if safe(report) {
			safeReports++
		}
	}

	fmt.Printf("Part 1: %v\n", safeReports)
}

func part2() {
	safeReports := 0

	for _, report := range makeReports() {
		if safe(report) || canBeSafe(report) {
			safeReports++
		}
	}

	fmt.Printf("Part 2: %v\n", safeReports)
}

func makeReports() [][]int {
	reports := make([][]int, 0)

	bytes, _ := os.ReadFile("input.txt")

	for _, row := range strings.Split(string(bytes), "\n") {
		report := make([]int, 0)

		for _, f := range strings.Fields(row) {
			n, _ := strconv.Atoi(f)
			report = append(report, n)
		}

		reports = append(reports, report)
	}

	return reports
}

func safe(ns []int) bool {
	inc := true
	dec := true

	for i := range ns {
		if i == len(ns)-1 {
			break
		}

		inc = inc && ns[i] < ns[i+1]
		dec = dec && ns[i] > ns[i+1]
		diff := int(math.Abs(float64(ns[i] - ns[i+1])))
		safe := (diff > 0 && diff < 4)

		if !safe {
			return false
		}
	}

	return inc || dec
}

func canBeSafe(ns []int) bool {
	if safe(ns) {
		return true
	}

	for i := range ns {
		if safe(removeElement(ns, i)) {
			return true
		}
	}

	return false
}

func removeElement(ns []int, index int) []int {
	new := make([]int, len(ns)-1)
	copy(new[:index], ns[:index])
	copy(new[index:], ns[index+1:])

	return new
}
