package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type octopus struct {
	val     int64
	flashed bool
}

func resetFlashers(matrix [][]octopus) [][]octopus {
	for i := int64(0); i < 10; i++ {
		for j := int64(0); j < 10; j++ {
			currentOctopus := matrix[i][j]
			if currentOctopus.flashed == true {
				matrix[i][j].val = 0
			}
			matrix[i][j].flashed = false
		}
	}

	return matrix
}

func updateMatrixByFlasher(matrix [][]octopus, flashCount int64) ([][]octopus, int64) {

	runningFlashCount := flashCount

	for i := int64(0); i < 10; i++ {
		for j := int64(0); j < 10; j++ {
			currentOctopus := matrix[i][j]

			// update to be flashed and bump neighbors by 1
			if currentOctopus.val > 9 && currentOctopus.flashed == false {
				// start over if a previously checked coordinate is bumped to above 9
				restart := false
				//  above
				if i-1 >= 0 {
					if matrix[i-1][j].val+1 > 9 && matrix[i-1][j].flashed == false {
						restart = true
					}
					matrix[i-1][j].val = matrix[i-1][j].val + 1
				}
				// above right
				if i-1 >= 0 && j+1 < 10 {
					if matrix[i-1][j+1].val+1 > 9 && matrix[i-1][j+1].flashed == false {
						restart = true
					}
					matrix[i-1][j+1].val = matrix[i-1][j+1].val + 1
				}
				//  right
				if j+1 < 10 {
					matrix[i][j+1].val = matrix[i][j+1].val + 1
				}
				// below right
				if i+1 < 10 && j+1 < 10 {
					matrix[i+1][j+1].val = matrix[i+1][j+1].val + 1
				}
				// below
				if i+1 < 10 {
					matrix[i+1][j].val = matrix[i+1][j].val + 1
				}
				// below left
				if i+1 < 10 && j-1 >= 0 {
					matrix[i+1][j-1].val = matrix[i+1][j-1].val + 1
				}
				// left
				if j-1 >= 0 {
					if matrix[i][j-1].val+1 > 9 && matrix[i][j-1].flashed == false {
						restart = true
					}
					matrix[i][j-1].val = matrix[i][j-1].val + 1
				}
				// above left
				if i-1 >= 0 && j-1 >= 0 {
					if matrix[i-1][j-1].val+1 > 9 && matrix[i-1][j-1].flashed == false {
						restart = true
					}
					matrix[i-1][j-1].val = matrix[i-1][j-1].val + 1
				}
				// update to be flashed
				matrix[i][j].flashed = true
				runningFlashCount++

				if restart {
					matrix, runningFlashCount = updateMatrixByFlasher(matrix, runningFlashCount)
					break
				}
			}

		}
	}

	return matrix, runningFlashCount
}

func checkAllFlashMatrix(matrix [][]octopus) bool {
	allFlash := true

	for i := int64(0); i < 10; i++ {
		for j := int64(0); j < 10; j++ {
			currentOctopus := matrix[i][j]
			if currentOctopus.flashed == false {
				allFlash = false
			}
		}
	}
	return allFlash
}

func dumboOctopus(data []string, steps int64) int64 {
	flashes := int64(0)
	currentSteps := steps
	var matrix [][]octopus

	for _, row := range data {
		vals := strings.Split(row, "")

		var newRow []octopus

		for _, char := range vals {
			num, _ := strconv.ParseInt(char, 10, 64)
			var newOctopus octopus
			newOctopus.flashed = false
			newOctopus.val = num
			newRow = append(newRow, newOctopus)
		}
		matrix = append(matrix, newRow)
	}

	for currentSteps > 0 {

		hasFlasher := false
		// First increase each octopus level by 1
		for i := int64(0); i < 10; i++ {
			for j := int64(0); j < 10; j++ {
				if matrix[i][j].val == 9 {
					hasFlasher = true
				}
				matrix[i][j].val = matrix[i][j].val + 1
			}
		}

		if hasFlasher {
			matrix, flashes = updateMatrixByFlasher(matrix, flashes)
		}

		matrix = resetFlashers(matrix)

		currentSteps--
	}

	return flashes
}

func dumboOctopusAllFlash(data []string) int64 {
	flashes := int64(0)
	allFlashed := false
	var matrix [][]octopus
	result := int64(0)

	for _, row := range data {
		vals := strings.Split(row, "")

		var newRow []octopus

		for _, char := range vals {
			num, _ := strconv.ParseInt(char, 10, 64)
			var newOctopus octopus
			newOctopus.flashed = false
			newOctopus.val = num
			newRow = append(newRow, newOctopus)
		}
		matrix = append(matrix, newRow)
	}

	for !allFlashed {

		hasFlasher := false
		// First increase each octopus level by 1
		for i := int64(0); i < 10; i++ {
			for j := int64(0); j < 10; j++ {
				if matrix[i][j].val == 9 {
					hasFlasher = true
				}
				matrix[i][j].val = matrix[i][j].val + 1
			}
		}

		if hasFlasher {
			matrix, flashes = updateMatrixByFlasher(matrix, flashes)
		}

		didAllFlash := checkAllFlashMatrix(matrix)

		if didAllFlash {
			allFlashed = true
		}

		matrix = resetFlashers(matrix)

		result++

	}

	return result
}
func TestDay11(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/11a.txt")
	solveInput, _ := readStringLines("./inputs/11.txt")

	// PART 1a
	if diff := cmp.Diff(int64(1656), dumboOctopus(sampleInput, int64(100))); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 1
	if diff := cmp.Diff(int64(1694), dumboOctopus(solveInput, int64(100))); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2
	if diff := cmp.Diff(int64(346), dumboOctopusAllFlash(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
