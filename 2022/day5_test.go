package adventOfCode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setupStack(rw string, stk [][]string) [][]string {
	for idx, letter := range strings.Split(rw, "") {
		fmt.Printf("Index: %v, Letter: %v \n", idx, letter)
	}

	return stk

}

func supplyStacks(data []string) string {
	result := ""
	supplyStacks := [][]string{}

	for _, row := range data {
		if strings.Contains(row, "[") {
			supplyStacks = setupStack(row, supplyStacks)

			fmt.Printf("%v \n", row)
		}
	}

	return result
}

func TestDay5(t *testing.T) {

	sampleInput, _ := readStringLines("./inputs/5sample.txt")
	// puzzleInput, _ := readStringLines("./inputs/4.txt")

	// Part 1

	if diff := cmp.Diff(supplyStacks(sampleInput), "CMZ"); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// if diff := cmp.Diff(campCleanup(puzzleInput), int64(562)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// Part 2

	// if diff := cmp.Diff(campCleanupAnyOverlap(sampleInput), int64(4)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// if diff := cmp.Diff(campCleanupAnyOverlap(puzzleInput), int64(924)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }
}
