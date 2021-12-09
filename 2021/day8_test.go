package adventOfCode

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func checkStringEasyValues(s string) bool {
	switch len(s) {
	case 2, 3, 4, 7:
		return true
	default:
		return false
	}
}

func sevenSegmentCheck(data []string) int64 {
	result := int64(0)
	var outputs []string

	for _, row := range data {
		val := strings.Split(row, "|")
		outputs = append(outputs, val[1])
	}

	for _, sequence := range outputs {

		singleRow := strings.Split(sequence, " ")

		for _, group := range singleRow {
			isEasyDigit := checkStringEasyValues(group)

			if isEasyDigit {
				result++
			}
		}

	}

	return result
}

func checkConfiguration(s string) string {

	single := strings.Split(s, " ")
	fmt.Printf("single %v \n", single)

	return s

	// acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab

}

func sevenSegmentCheckSignals(data []string) int64 {
	result := int64(0)
	displayKey := make(map[int]string)

	displayKey[0] = "abcefg"
	displayKey[1] = "cf"
	displayKey[2] = "acdeg"
	displayKey[3] = "acdfg"
	displayKey[4] = "bcdf"
	displayKey[5] = "abdfg"
	displayKey[6] = "abdefg"
	displayKey[7] = "acf"
	displayKey[8] = "abcdefg"
	displayKey[9] = "abcdfg"

	// s := make([]string, 1)
	// s[0] = val[1]

	for _, row := range data {
		val := strings.Split(row, "|")

		legend := checkConfiguration(val[0])
		if len(legend) > 100 {
			fmt.Printf("legend: %v \n", legend)
		}
	}

	return result
}

func TestDay8(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/8a.txt")
	solveInput, _ := readStringLines("./inputs/8.txt")

	// PART 1a
	if diff := cmp.Diff(int64(26), sevenSegmentCheck(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 1
	if diff := cmp.Diff(int64(272), sevenSegmentCheck(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2a
	// if diff := cmp.Diff(int64(61229), sevenSegmentCheckSignals(sampleInput)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

}
