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

	result = int64(len(activeData))

	return result
}

func simulateLanternfishBetter(data []string, days int64) int64 {
	result := int64(0)
	currentFish := make([]int64, 9)

	for _, num := range strings.Split(data[0], ",") {
		val, _ := strconv.ParseInt(num, 10, 64)
		currentFish[val] = currentFish[val] + 1
	}

	for i := int64(0); i < days; i++ {
		newFish := currentFish[0]
		for i := 0; i < len(currentFish)-1; i++ {
			currentFish[i] = currentFish[i+1]
		}
		currentFish[8] = newFish
		currentFish[6] = currentFish[6] + newFish
	}

	for _, n := range currentFish {
		result = result + n
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
	if diff := cmp.Diff(int64(390011), simulateLanternfish(solveInput, 80)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2a

	// PART 2
	if diff := cmp.Diff(int64(1746710169834), simulateLanternfishBetter(solveInput, 256)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
