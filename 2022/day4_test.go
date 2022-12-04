package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func campCleanup(data []string) int64 {
	result := int64(0)

	for _, day := range data {
		pairs := strings.Split(day, ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")

		low1, _ := strconv.Atoi(first[0])
		high1, _ := strconv.Atoi(first[1])

		low2, _ := strconv.Atoi(second[0])
		high2, _ := strconv.Atoi(second[1])

		if low1 <= low2 &&
			high1 >= low2 &&
			low1 <= high2 &&
			high1 >= high2 {
			result = result + 1
		} else if low2 <= low1 &&
			high2 >= low1 &&
			low2 <= high1 &&
			high2 >= high1 {
			result = result + 1
		}

	}

	return result
}

func TestDay4(t *testing.T) {

	sampleInput, _ := readStringLines("./inputs/4sample.txt")
	puzzleInput, _ := readStringLines("./inputs/4.txt")

	// Part 1

	if diff := cmp.Diff(campCleanup(sampleInput), int64(2)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(campCleanup(puzzleInput), int64(8252)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// Part 2

	// if diff := cmp.Diff(ruckSackBadges(sampleInput), int64(70)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// if diff := cmp.Diff(ruckSackBadges(puzzleInput), int64(2828)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }
}
