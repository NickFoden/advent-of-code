package adventOfCode

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func checkOpenIndex(options []string, val string) (int64, bool) {

	for index, next := range options {
		if val == next {
			return int64(index), true
		}
	}
	return -1, false
}

func checkCloseIndex(options []string, val string) (int64, bool) {

	for index, next := range options {
		if val == next {
			return int64(index), true
		}
	}
	return -1, false
}

func scoreCharFails(chars []string) int64 {
	result := int64(0)
	rubric := map[string]int64{")": 3, "]": 57, "}": 1197, ">": 25137}

	for _, c := range chars {
		result += rubric[c]
	}
	return result
}

func syntaxScoring(data []string) int64 {
	result := int64(0)

	open := []string{"(", "[", "{", "<"}
	close := []string{")", "]", "}", ">"}

	var charFails []string

	for _, row := range data {
		var charFail string
		var memory []int64
		vals := strings.Split(row, "")

		for charIndex, char := range vals {
			if len(charFail) > 0 {
				break
			}

			if charIndex == len(vals)-1 {
				break
			}

			lastMemoryIndex := int64(0)
			if len(memory) > 0 {
				lastMemoryIndex = memory[len(memory)-1]
			}

			indexOfOpen, foundOpen := checkOpenIndex(open, char)
			indexOfClosed, foundClosed := checkCloseIndex(close, char)

			if foundOpen {
				memory = append(memory, indexOfOpen)
			} else if len(memory) > 0 {
				if foundClosed && indexOfClosed == lastMemoryIndex {
					memory = memory[:len(memory)-1]
				} else {
					charFail = char
				}
			} else {
				break
			}
		}
		if len(charFail) > 0 {
			charFails = append(charFails, charFail)
		}
	}

	result = scoreCharFails(charFails)
	return result
}

func TestDay10(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/10a.txt")
	solveInput, _ := readStringLines("./inputs/10.txt")

	// PART 1a
	if diff := cmp.Diff(int64(26397), syntaxScoring(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 1
	if diff := cmp.Diff(int64(271245), syntaxScoring(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
