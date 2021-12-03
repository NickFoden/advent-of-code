package adventOfCode

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func powerConsumption(data []string) int64 {
	rowLength := len(data[0])
	bits := make([]int64, rowLength)

	for _, row := range data {
		for index, bit := range row {
			if string(bit) == "0" {
				bits[index] = bits[index] + 1
			} else {
				bits[index] = bits[index] - 1
			}
		}
	}

	// fmt.Printf("Bits = %v", bits)

	gammaBinary := ""
	epsilonBinary := ""

	for _, b := range bits {
		if b > 0 {
			gammaBinary += "0"
			epsilonBinary += "1"
		} else {
			gammaBinary += "1"
			epsilonBinary += "0"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaBinary, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonBinary, 2, 64)

	return gammaRate * epsilonRate
}

func TestDay3(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/3a.txt")

	if diff := cmp.Diff(int64(198), powerConsumption(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	solveInput, _ := readStringLines("./inputs/3.txt")

	// PART 1
	if diff := cmp.Diff(int64(1131506), powerConsumption(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
	// PART 2
	// if diff := cmp.Diff(calculateDistanceWithAIm(b), int64(1848454425)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }
}
