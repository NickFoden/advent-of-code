package adventOfCode

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func checkIncrements(data []int64) int64 {
	result := int64(0)
	iterator := int64(0)

	for index, v := range data {
		if index != 0 && v > int64(iterator) {
			result++
		}
		iterator = v
	}

	return result
}

func TestDay1(t *testing.T) {
	// Test the first example
	sampleInput := []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if diff := cmp.Diff(checkIncrements(sampleInput), int64(7)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	b, _ := readLines("./inputs/1.txt")

	if diff := cmp.Diff(checkIncrements(b), int64(1722)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
