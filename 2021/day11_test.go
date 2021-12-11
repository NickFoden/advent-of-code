package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func dumboOctopus(data []string, steps int64) int64 {
	flashes := int64(0)
	currentSteps := steps

	var matrix [][]int64

	for _, row := range data {
		vals := strings.Split(row, "")

		var newRow []int64

		for _, char := range vals {
			num, _ := strconv.ParseInt(char, 10, 64)
			newRow = append(newRow, num)
		}
		matrix = append(matrix, newRow)
	}

	iBound := int64(len(matrix))
	jBound := int64(len(matrix[0]))

	for currentSteps > 0 {

		flashChecks := make(map[int64]int64)

		for i := int64(0); i < iBound; i++ {
			for j := int64(0); j < jBound; j++ {
				val := matrix[i][j]
				if val == 9 {
					matrix[i][j] = 0
					flashChecks[i] = j

					//check above
					if i-1 > 0 {
						if matrix[i-1][j] == 9 {
							flashChecks[i-1] = j
							matrix[i-1][j] = 0
						} else {
							matrix[i-1][j] = matrix[i-1][j] + 1
						}
					}

					//check above right
					if i-1 > 0 && j+1 < jBound {
						if matrix[i-1][j+1] == 9 {
							flashChecks[i-1] = j + 1
							matrix[i-1][j+1] = 0
						} else {
							matrix[i-1][j+1] = matrix[i-1][j+1] + 1
						}
					}

					//check above left
					if i-1 > 0 && j-1 > 0 {
						if matrix[i-1][j-1] == 9 {
							flashChecks[i-1] = j - 1
							matrix[i-1][j-1] = 0
						} else {
							matrix[i-1][j-1] = matrix[i-1][j-1] + 1
						}
					}
					// check below
					if i+1 < iBound {
						if matrix[i+1][j] == 9 {
							flashChecks[i+1] = j
							matrix[i+1][j] = 0
						} else {
							matrix[i+1][j] = matrix[i+1][j] + 1
						}
					}

					// check right
					if j+1 < jBound {
						if matrix[i][j+1] == 9 {
							flashChecks[i] = j + 1
							matrix[i][j+1] = 0
						} else {
							matrix[i][j+1] = matrix[i][j+1] + 1
						}
					}

					// check left
					if j-1 > 0 {
						if matrix[i][j-1] == 9 {
							flashChecks[i] = j - 1
							matrix[i][j-1] = 0
						} else {
							matrix[i][j-1] = matrix[i][j-1] + 1
						}
					}

					//check below left
					if i+1 < iBound && j-1 > 0 {
						if matrix[i+1][j-1] == 9 {
							flashChecks[i+1] = j - 1
							matrix[i+1][j-1] = 0
						} else {
							matrix[i+1][j-1] = matrix[i+1][j-1] + 1
						}
					}

					//check below right
					if i+1 < iBound && j+1 < jBound {
						if matrix[i+1][j+1] == 9 {
							flashChecks[i+1] = j + 1
							matrix[i+1][j+1] = 0
						} else {
							matrix[i+1][j+1] = matrix[i+1][j+1] + 1
						}
					}
				} else {
					matrix[i][j] = val + 1
				}
			}
		}
		flashes += int64(len(flashChecks))
		currentSteps = currentSteps - 1
	}

	return flashes
}

func TestDay11(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/11a.txt")
	// solveInput, _ := readLines("./inputs/11.txt")

	// PART 1a
	if diff := cmp.Diff(int64(1656), dumboOctopus(sampleInput, int64(100))); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
