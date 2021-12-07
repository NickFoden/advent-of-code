package adventOfCode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func crabFuel(data []string) int64 {
	result := int64(0)
	var crabs []int64

	total := int64(0)

	for _, num := range strings.Split(data[0], ",") {
		val, _ := strconv.ParseInt(num, 10, 64)
		total = total + val
		crabs = append(crabs, val)
	}

	fmt.Printf("total: %v \n", total)
	fmt.Printf("crabs: %v \n", len(crabs))

	fmt.Printf("mean: %v \n", (total / int64(len(crabs))))

	return result
}

func TestDay7(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/7a.txt")
	// solveInput, _ := readStringLines("./inputs/7.txt")

	// PART 1a
	if diff := cmp.Diff(int64(37), crabFuel(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
