package adventOfCode

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func mostCalories(data []int64) int64 {
	max := int64(0)
	currentElf := int64(0)

	for _, v := range data {
		if v == 0 {
			if currentElf > max {
				max = currentElf
			}
			currentElf = 0
		} else {
			currentElf = currentElf + v
		}
	}

	return max
}

func topThreeMostCalories(data []int64) int64 {
	max := int64(0)
	topThree := [3]int64{0, 0, 0}
	allvals := []int64{}
	currentElf := int64(0)
	total := len(data)

	for i := 0; i < total+1; i++ {
		if i == total {
			allvals = append(allvals, currentElf)
		} else {
			v := data[i]
			if v == 0 {
				if currentElf > 0 {
					allvals = append(allvals, currentElf)
				}
				currentElf = 0
			} else {
				currentElf = currentElf + v
			}
		}
	}

	sort.Slice(allvals, func(i, j int) bool { return allvals[i] > allvals[j] })

	for i := 0; i < 3; i++ {
		topThree[i] = allvals[i]
	}

	for _, e := range topThree {
		max = max + e
	}

	return max
}

func TestDay1(t *testing.T) {

	sampleInput, _ := readLines("./inputs/1sample.txt")
	puzzleInput, _ := readLines("./inputs/1.txt")

	// Part 1

	if diff := cmp.Diff(mostCalories(sampleInput), int64(24000)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(mostCalories(puzzleInput), int64(70116)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// Part 2

	if diff := cmp.Diff(topThreeMostCalories(sampleInput), int64(45000)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(topThreeMostCalories(puzzleInput), int64(206582)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
