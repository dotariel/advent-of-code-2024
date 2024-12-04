package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	as := make([]int, 0)
	bs := make([]int, 0)

	bytes, _ := os.ReadFile("input.txt")

	for _, row := range strings.Split(string(bytes), "\n") {
		fields := strings.Fields(row)
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])

		as = append(as, a)
		bs = append(bs, b)
	}

	fmt.Printf("Part 1: %v\n", Diff(as, bs))
	fmt.Printf("Part 2: %v\n", Similarity(as, bs))
}

func Diff(a, b []int) int {
	slices.Sort(a)
	slices.Sort(b)

	diff := 0

	for i := range a {
		diff += int(math.Abs(float64(a[i]) - float64(b[i])))
	}

	return diff
}

func Similarity(a, b []int) int {
	similarity := 0

	for _, val := range a {
		similarity += val * findCardinality(b, val)
	}

	return similarity
}

func findCardinality(ns []int, n int) int {
	count := 0

	for _, val := range ns {
		if val == n {
			count++
		}
	}

	return count
}
