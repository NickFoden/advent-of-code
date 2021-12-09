package adventOfCode

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func calculateRiskLevel(vals []int64) int64 {
	result := int64(0)

	for _, num := range vals {
		result += num + 1
	}

	return result
}

func friendsInLowPlaces(data []string) int64 {
	result := int64(0)

	var lowPoints []int64

	var caves [][]int64

	for _, row := range data {

		var rowVals []int64
		val := strings.Split(row, "")

		for _, num := range val {
			i, _ := strconv.ParseInt(num, 10, 64)
			rowVals = append(rowVals, i)
		}
		caves = append(caves, rowVals)
	}

	iBound := int64(len(caves))
	jBound := int64(len(data[0]))

	for i := int64(0); i < iBound; i++ {

		for j := int64(0); j < jBound; j++ {
			currentVal := caves[i][j]
			isLow := true

			// check left
			if j != 0 {
				if currentVal >= caves[i][j-1] {
					isLow = false
				}
			}
			// check above
			if i != 0 {
				if currentVal >= caves[i-1][j] {
					isLow = false
				}
			}

			// check right
			if j < jBound-1 {
				if currentVal >= caves[i][j+1] {
					isLow = false
				}
			}
			// check below
			if i < iBound-1 {
				if currentVal >= caves[i+1][j] {
					isLow = false
				}
			}

			if isLow {
				lowPoints = append(lowPoints, currentVal)
			}
		}
	}

	result = calculateRiskLevel(lowPoints)

	return result
}

func hasUncheckedCoords(basinCoords [][]int64) bool {

	for _, coords := range basinCoords {
		if coords[2] < 1 {
			return true
		}
	}

	return false
}

func nextCoordsToCheck(basinCoords [][]int64) []int64 {
	var result []int64
	for _, coords := range basinCoords {
		if coords[2] < 1 {
			result = coords
			break
		}
	}
	return result
}

func updatedBasinCoords(basinCoords [][]int64, coords []int64) [][]int64 {
	var result [][]int64
	for _, bc := range basinCoords {
		if coords[0] == bc[0] && coords[1] == bc[1] {
			c := make([]int64, 3)
			c[0] = bc[0]
			c[1] = bc[1]
			c[2] = 1
			result = append(result, c)
		} else {
			result = append(result, bc)
		}
	}
	return result
}

func notAddedYet(iCheck int64, jCheck int64, currentBasin [][]int64) bool {
	result := true

	for _, coords := range currentBasin {
		if coords[0] == iCheck && coords[1] == jCheck {
			result = false
		}

	}

	return result
}

func calculateBasins(i int64, j int64, caves [][]int64) [][]int64 {
	iBound := int64(len(caves))
	jBound := int64(len(caves[0]))

	c := make([]int64, 3)
	c[0] = i
	c[1] = j
	c[2] = 0

	var basinCoordinates [][]int64

	// add initial coordinates
	basinCoordinates = append(basinCoordinates, c)

	for hasUncheckedCoords(basinCoordinates) {

		coordsToCheck := nextCoordsToCheck(basinCoordinates)

		currentVal := caves[coordsToCheck[0]][coordsToCheck[1]]

		coordI := coordsToCheck[0]
		coordJ := coordsToCheck[1]

		if coordJ != 0 {
			// check left
			if currentVal <= caves[coordI][coordJ-1] && caves[coordI][coordJ-1] != 9 {
				if notAddedYet(coordI, coordJ-1, basinCoordinates) {
					c := make([]int64, 3)
					c[0] = coordI
					c[1] = coordJ - 1
					c[2] = 0
					basinCoordinates = append(basinCoordinates, c)
				}
			}
		}

		if coordI != 0 {
			// check above
			if currentVal <= caves[coordI-1][coordJ] && caves[coordI-1][coordJ] != 9 {
				if notAddedYet(coordI-1, coordJ, basinCoordinates) {
					c := make([]int64, 3)
					c[0] = coordI - 1
					c[1] = coordJ
					c[2] = 0
					basinCoordinates = append(basinCoordinates, c)
				}
			}
		}

		// check right
		if coordJ < jBound-1 {
			if currentVal <= caves[coordI][coordJ+1] && caves[coordI][coordJ+1] != 9 {
				if notAddedYet(coordI, coordJ+1, basinCoordinates) {
					c := make([]int64, 3)
					c[0] = coordI
					c[1] = coordJ + 1
					c[2] = 0
					basinCoordinates = append(basinCoordinates, c)
				}
			}
		}
		// check below
		if coordI < iBound-1 {
			if currentVal <= caves[coordI+1][coordJ] && caves[coordI+1][coordJ] != 9 {
				if notAddedYet(coordI+1, coordJ, basinCoordinates) {
					c := make([]int64, 3)
					c[0] = coordI + 1
					c[1] = coordJ
					c[2] = 0
					basinCoordinates = append(basinCoordinates, c)
				}
			}
		}

		basinCoordinates = updatedBasinCoords(basinCoordinates, coordsToCheck)
	}

	return basinCoordinates
}

func calculateFinalBasinsTotal(basins []int64) int64 {

	// order largest basins first
	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })

	// multiply the 3 largest basins
	return basins[0] * basins[1] * basins[2]
}

func basinsCheck(data []string) int64 {
	result := int64(0)
	var caves [][]int64
	var basinPoints []int64

	for _, row := range data {

		var rowVals []int64
		val := strings.Split(row, "")

		for _, num := range val {
			i, _ := strconv.ParseInt(num, 10, 64)
			rowVals = append(rowVals, i)
		}
		caves = append(caves, rowVals)
	}

	iBound := int64(len(caves))
	jBound := int64(len(data[0]))

	for i := int64(0); i < iBound; i++ {

		for j := int64(0); j < jBound; j++ {
			currentVal := caves[i][j]
			isLow := true

			// check left
			if j != 0 {
				if currentVal >= caves[i][j-1] {
					isLow = false
				}
			}
			// check above
			if i != 0 {
				if currentVal >= caves[i-1][j] {
					isLow = false
				}
			}

			// check right
			if j < jBound-1 {
				if currentVal >= caves[i][j+1] {
					isLow = false
				}
			}
			// check below
			if i < iBound-1 {
				if currentVal >= caves[i+1][j] {
					isLow = false
				}
			}

			if isLow {
				basinsTotal := calculateBasins(i, j, caves)
				basinPoints = append(basinPoints, int64(len(basinsTotal)))
			}
		}
	}

	result = calculateFinalBasinsTotal(basinPoints)

	return result
}

func TestDay9(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/9a.txt")
	solveInput, _ := readStringLines("./inputs/9.txt")

	// PART 1a
	if diff := cmp.Diff(int64(15), friendsInLowPlaces(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 1
	if diff := cmp.Diff(int64(491), friendsInLowPlaces(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2a
	if diff := cmp.Diff(int64(1134), basinsCheck(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2
	if diff := cmp.Diff(int64(1075536), basinsCheck(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
