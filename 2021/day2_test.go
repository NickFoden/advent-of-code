package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func calculateDistance(data []string) int64 {
	horizontal := int64(0)
	depth := int64(0)
	for _, v := range data {

		line := strings.Fields(v)
		direction := line[0]
		val, _ := strconv.ParseInt(line[1], 10, 64)

		switch direction {
		case "forward":
			horizontal = horizontal + val
		case "down":
			depth = depth + val
		case "up":
			depth = depth - val
		}
	}
	return horizontal * depth
}

func TestDay2(t *testing.T) {
	// Test the first example
	sampleInput := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	if diff := cmp.Diff(calculateDistance(sampleInput), int64(150)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	b, _ := readStringLines("./inputs/2.txt")

	// PART 1
	if diff := cmp.Diff(calculateDistance(b), int64(1561344)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
