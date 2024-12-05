package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	safeReports := 0

	bytes, _ := os.ReadFile("input.txt")

	for _, row := range strings.Split(string(bytes), "\n") {
		report := make([]int, 0)

		for _, f := range strings.Fields(row) {
			n, _ := strconv.Atoi(f)
			report = append(report, n)
		}

		if safe(report) {
			safeReports++
		}
	}

	fmt.Printf("Part 1: %v\n", safeReports)
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
