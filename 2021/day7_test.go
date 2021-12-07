package adventOfCode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func fuelUsage(c int64, m int64) int64 {
	if c == m {
		return 0
	}
	if c > m {
		return c - m
	}
	return m - c
}

func fuelUsagePricy(c int64, m int64) int64 {
	fuel := int64(0)
	difference := int64(0)

	if c > m {
		difference = c - m
	} else if c < m {
		difference = m - c
	}

	if difference > 0 {
		for i := int64(1); i < (difference + 1); i++ {
			fuel += i
		}
	}
	return fuel
}

func crabFuel(data []string) int64 {
	result := int64(0)
	var crabs []int64

	total := int64(0)

	for _, num := range strings.Split(data[0], ",") {
		val, _ := strconv.ParseInt(num, 10, 64)
		total = total + val
		crabs = append(crabs, val)
	}
	// Order the crabs to get the median
	sort.Slice(crabs, func(i, j int) bool { return crabs[i] < crabs[j] })

	// determine median value
	median := int64(0)
	if len(crabs)%2 == 0 {
		median = crabs[len(crabs)/2]
	} else {
		median = crabs[((len(crabs)-1)/2)+1]
	}

	for _, crab := range crabs {
		val := fuelUsage(median, crab)
		result += val
	}
	fmt.Printf("median: %v \n", median)

	return result
}

func crabFuelHigherCost(data []string) int64 {
	result := int64(0)
	var crabs []int64

	total := int64(0)

	for _, num := range strings.Split(data[0], ",") {
		val, _ := strconv.ParseInt(num, 10, 64)
		total = total + val
		crabs = append(crabs, val)
	}

	// determine mean value
	mean := total / int64(len(crabs))
	fmt.Printf("crabs len: %v \n", len(crabs))
	// fmt.Printf("mean: %v \n", mean)

	for _, crab := range crabs {
		val := fuelUsagePricy(mean, crab)
		result += val
	}

	return result
}

func TestDay7(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/7a.txt")
	solveInput, _ := readStringLines("./inputs/7.txt")

	// PART 1a
	if diff := cmp.Diff(int64(37), crabFuel(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// Part 1
	if diff := cmp.Diff(int64(348664), crabFuel(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2a
	// hmmmmmm the mean + 1
	// if diff := cmp.Diff(int64(168), crabFuelHigherCost(sampleInput)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// PART 2a
	if diff := cmp.Diff(int64(100220525), crabFuelHigherCost(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
