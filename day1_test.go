package adventOfCode

import (
	"bufio"
	"os"
	"strconv"
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

func readLines(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		lines = append(lines, val)
	}
	return lines, scanner.Err()
}

func TestDay1(t *testing.T) {
	// Test the first example
	sampleInput := []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if diff := cmp.Diff(checkIncrements(sampleInput), int64(7)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	b, _ := readLines("day1input.txt")

	if diff := cmp.Diff(checkIncrements(b), int64(1722)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
