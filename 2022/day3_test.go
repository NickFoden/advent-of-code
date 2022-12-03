package adventOfCode

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func determinePriority(item string) int64 {
	result := int64(0)
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for idx, v := range alphabet {
		if item == v {
			result = int64(idx) + 1
			break
		} else if item == strings.ToUpper(v) {
			result = int64(idx) + 1 + 26
			break
		}
	}
	return result
}

func ruckSackPriority(data []string) int64 {
	sum := int64(0)
	duplicateItems := []string{}

	for _, v := range data {
		totalLength := len(v)
		halfway := totalLength / 2

		firstCompartment := v[0:halfway]
		secondCompartment := strings.Split(v[halfway:totalLength], "")

		duplicates := []string{}

		for _, letter := range firstCompartment {
			if containsString(secondCompartment, string(letter)) && !containsString(duplicates, string(letter)) {
				duplicates = append(duplicates, string(letter))
				duplicateItems = append(duplicateItems, string(letter))
			}
		}

	}
	for _, v := range duplicateItems {
		priority := determinePriority(v)
		sum = sum + priority
	}

	return sum
}

func ruckSackBadges(data []string) int64 {

	sum := int64(0)
	badges := []string{}

	currentGroup := []string{}

	currentDuplicates := []string{}

	for i := 0; i < len(data); i++ {
		if i%3 == 2 {
			currentGroup = append(currentGroup, data[i])

			for _, letter := range currentGroup[0] {
				sLetter := string(letter)
				if (containsLetter(currentGroup[1], sLetter)) && (containsLetter(currentGroup[2], sLetter)) && !containsString(currentDuplicates, sLetter) {
					currentDuplicates = append(currentDuplicates, sLetter)
				}
			}

			if len(currentDuplicates) > 0 {
				badges = append(badges, currentDuplicates[0])
			}
			currentGroup = currentGroup[:0]
			currentDuplicates = currentDuplicates[:0]
		} else {
			currentGroup = append(currentGroup, data[i])
		}

	}

	for _, v := range badges {
		priority := determinePriority(v)
		sum = sum + priority
	}

	return sum

}

func TestDay3(t *testing.T) {

	sampleInput, _ := readStringLines("./inputs/3sample.txt")
	puzzleInput, _ := readStringLines("./inputs/3.txt")

	// Part 1

	if diff := cmp.Diff(ruckSackPriority(sampleInput), int64(157)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(ruckSackPriority(puzzleInput), int64(8252)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// Part 2

	if diff := cmp.Diff(ruckSackBadges(sampleInput), int64(70)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(ruckSackBadges(puzzleInput), int64(2828)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
