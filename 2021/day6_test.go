package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func updatedTime(current int64) (int64, bool) {
	switch current {
	case 1:
		return 0, false
	case 2:
		return 1, false
	case 3:
		return 2, false
	case 4:
		return 3, false
	case 5:
		return 4, false
	case 6:
		return 5, false
	case 7:
		return 6, false
	case 8:
		return 7, false
	default:
		// 0 case
		return 6, true
	}
}

func simulateLanternfish(data []string, days int64) int64 {
	result := int64(0)
	var activeData []int64

	firstRow := strings.Split(data[0], ",")
	for _, single := range firstRow {
		num, _ := strconv.ParseInt(single, 10, 64)
		activeData = append(activeData, num)
	}

	for i := int64(0); i < days; i++ {
		for index, single := range activeData {
			newVal, newFish := updatedTime(single)

			activeData[index] = newVal

			if newFish {
				activeData = append(activeData, int64(8))
			}

		}

	}

	for _, char := range activeData {
		result++
	}

	return result
}

func TestDay6(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/6a.txt")
	solveInput, _ := readStringLines("./inputs/6.txt")

	// PART 1a
	if diff := cmp.Diff(int64(5934), simulateLanternfish(sampleInput, 80)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 1
	// if diff := cmp.Diff(int64(390011), simulateLanternfish(solveInput, 80)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// PART 2a

	// PART 2
	if diff := cmp.Diff(int64(390011), simulateLanternfish(solveInput, 256)); diff != "" {

		// 1248036 too low
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
