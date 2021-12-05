package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func determineHighNumber(nums []string) int64 {
	result := int64(0)
	for _, single := range nums {
		asNum, _ := strconv.ParseInt(single, 10, 64)
		if asNum > result {
			result = asNum
		}
	}
	return result
}
func highLow(num1 int64, num2 int64) (int64, int64) {
	if num1 >= num2 {
		return num1, num2
	}
	return num2, num1
}

func gradeMatrix(m [][]int64) int64 {
	result := int64(0)

	for _, single := range m {
		for _, num := range single {
			if num > 1 {
				result += 1
			}
		}
	}

	return result
}

func ventCheck(data []string) int64 {
	final := int64(0)
	var matrix [][]int64

	highNumber := int64(0)
	var allRows []string

	for _, row := range data {
		cleanedRow := strings.ReplaceAll(row, " ", "")
		splitRow := strings.Split(cleanedRow, "->")

		leftCoords := strings.Split(splitRow[0], ",")
		rightCoords := strings.Split(splitRow[1], ",")

		allRows = append(allRows, leftCoords...)
		allRows = append(allRows, rightCoords...)
	}

	highNumber = determineHighNumber(allRows)

	for i := int64(0); i < (highNumber + 1); i++ {
		matrixRow := make([]int64, (highNumber + 1))
		for i := int64(0); i < (highNumber + 1); i++ {
			matrixRow[i] = int64(0)
		}
		matrix = append(matrix, matrixRow)
	}

	for _, row := range data {
		cleanedRow := strings.ReplaceAll(row, " ", "")
		splitRow := strings.Split(cleanedRow, "->")

		leftCoords := strings.Split(splitRow[0], ",")
		rightCoords := strings.Split(splitRow[1], ",")

		xStart, _ := strconv.ParseInt(leftCoords[0], 10, 64)
		xFinish, _ := strconv.ParseInt(rightCoords[0], 10, 64)
		yStart, _ := strconv.ParseInt(leftCoords[1], 10, 64)
		yFinish, _ := strconv.ParseInt(rightCoords[1], 10, 64)

		for i := int64(0); i < (highNumber + 1); i++ {
			if yStart == yFinish && yStart == i {
				if xStart <= xFinish {
					for j := int64(0); j < (highNumber + 1); j++ {
						if j >= xStart && j <= xFinish {
							matrix[i][j] = matrix[i][j] + 1
						}
					}
				} else {
					for j := int64(0); j < (highNumber + 1); j++ {
						if j >= xFinish && j <= xStart {
							matrix[i][j] = matrix[i][j] + 1
						}
					}
				}
			} else if xStart == xFinish {
				if yStart <= yFinish {
					if yStart <= i && yFinish >= i {
						matrix[i][xStart] = matrix[i][xStart] + 1
					}
				} else {
					if yStart >= i && yFinish <= i {
						matrix[i][xStart] = matrix[i][xStart] + 1
					}
				}

			}
		}
	}

	final = gradeMatrix(matrix)

	return final
}

func ventCheckDiagonals(data []string) int64 {
	final := int64(0)
	var matrix [][]int64

	highNumber := int64(0)
	var allRows []string

	for _, row := range data {
		cleanedRow := strings.ReplaceAll(row, " ", "")
		splitRow := strings.Split(cleanedRow, "->")

		leftCoords := strings.Split(splitRow[0], ",")
		rightCoords := strings.Split(splitRow[1], ",")

		allRows = append(allRows, leftCoords...)
		allRows = append(allRows, rightCoords...)
	}

	highNumber = determineHighNumber(allRows)

	for i := int64(0); i < (highNumber + 1); i++ {
		matrixRow := make([]int64, (highNumber + 1))
		for i := int64(0); i < (highNumber + 1); i++ {
			matrixRow[i] = int64(0)
		}
		matrix = append(matrix, matrixRow)
	}

	for _, row := range data {
		cleanedRow := strings.ReplaceAll(row, " ", "")
		splitRow := strings.Split(cleanedRow, "->")

		leftCoords := strings.Split(splitRow[0], ",")
		rightCoords := strings.Split(splitRow[1], ",")

		xStart, _ := strconv.ParseInt(leftCoords[0], 10, 64)
		xFinish, _ := strconv.ParseInt(rightCoords[0], 10, 64)
		yStart, _ := strconv.ParseInt(leftCoords[1], 10, 64)
		yFinish, _ := strconv.ParseInt(rightCoords[1], 10, 64)

		for i := int64(0); i < (highNumber + 1); i++ {
			if yStart == yFinish && yStart == i {
				if xStart <= xFinish {
					for j := int64(0); j < (highNumber + 1); j++ {
						if j >= xStart && j <= xFinish {
							matrix[i][j] = matrix[i][j] + 1
						}
					}
				} else {
					for j := int64(0); j < (highNumber + 1); j++ {
						if j >= xFinish && j <= xStart {
							matrix[i][j] = matrix[i][j] + 1
						}
					}
				}
			} else if xStart == xFinish {
				if yStart <= yFinish {
					if yStart <= i && yFinish >= i {
						matrix[i][xStart] = matrix[i][xStart] + 1
					}
				} else {
					if yStart >= i && yFinish <= i {
						matrix[i][xStart] = matrix[i][xStart] + 1
					}
				}

			}
		}

		if yStart != yFinish || xStart != xFinish {
			if yStart < yFinish && xStart < xFinish {
				for i := int64(0); i <= (yFinish - yStart); i++ {
					yVal := yStart + i
					xVal := xStart + i
					matrix[yVal][xVal] = matrix[yVal][xVal] + 1
				}
			} else if yStart > yFinish && xStart < xFinish {
				for i := int64(0); i <= (yStart - yFinish); i++ {
					yVal := yStart - i
					xVal := xStart + i
					matrix[yVal][xVal] = matrix[yVal][xVal] + 1
				}

			} else if yStart > yFinish && xStart > xFinish {
				for i := int64(0); i <= (yStart - yFinish); i++ {
					yVal := yStart - i
					xVal := xStart - i
					matrix[yVal][xVal] = matrix[yVal][xVal] + 1
				}
			} else if yStart < yFinish && xStart > xFinish {
				for i := int64(0); i <= (yFinish - yStart); i++ {
					yVal := yStart + i
					xVal := xStart - i
					matrix[yVal][xVal] = matrix[yVal][xVal] + 1
				}
			}
		}

	}

	final = gradeMatrix(matrix)

	return final
}

func TestDay5(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/5a.txt")
	solveInput, _ := readStringLines("./inputs/5.txt")

	// PART 1a
	if diff := cmp.Diff(int64(5), ventCheck(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 1
	if diff := cmp.Diff(int64(6572), ventCheck(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2a
	if diff := cmp.Diff(int64(12), ventCheckDiagonals(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2
	if diff := cmp.Diff(int64(21466), ventCheckDiagonals(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
