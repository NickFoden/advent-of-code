package adventOfCode

import (
	"fmt"
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

func mostPopularBit(d []string, i int64) string {

	bitCheck := 0

	for _, row := range d {
		if string(row[i]) == "0" {
			bitCheck += 1
		} else {
			bitCheck -= 1
		}

	}

	if bitCheck == 0 {
		return "1"
	}

	if bitCheck > 0 {
		return "0"
	}
	return "1"

}

func lifeSupport(data []string) int64 {
	rowLength := len(data[0])

	oxygenBinary := ""
	co2ScrubberBinary := ""

	oxygenSlice := data
	co2Slice := data

	// get oxygen slice
	for i := 0; i < rowLength; i++ {
		if len(oxygenSlice) <= 1 {
			break
		}
		var tempSlice []string
		highBit := mostPopularBit(oxygenSlice, int64(i))

		for _, val := range oxygenSlice {
			if string(val[i]) == highBit {
				tempSlice = append(tempSlice, val)
			}
		}
		oxygenSlice = tempSlice
		oxygenBinary = oxygenSlice[0]
	}

	// get co2 slice
	for i := 0; i < rowLength; i++ {
		if len(co2Slice) <= 1 {
			break
		}
		var tempSlice []string

		highBit := mostPopularBit(co2Slice, int64(i))

		for _, val := range co2Slice {
			if string(val[i]) != highBit {
				tempSlice = append(tempSlice, val)
			}
		}
		co2Slice = tempSlice
		co2ScrubberBinary = co2Slice[0]
	}

	oxygenGenerator, _ := strconv.ParseInt(oxygenBinary, 2, 64)
	co2Scrubber, _ := strconv.ParseInt(co2ScrubberBinary, 2, 64)

	fmt.Printf("oxygenGenerator=%v, co2Scrubber=%v \n", oxygenGenerator, co2Scrubber)

	return oxygenGenerator * co2Scrubber
}

func TestDay3(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/3a.txt")

	solveInput, _ := readStringLines("./inputs/3.txt")

	// PART 1
	if diff := cmp.Diff(int64(198), powerConsumption(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(int64(1131506), powerConsumption(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2
	if diff := cmp.Diff(int64(230), lifeSupport(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(int64(7863147), lifeSupport(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
