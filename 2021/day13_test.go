package adventOfCode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func checkDots(matrix [][]string) int64 {
	result := int64(0)
	for i := int64(0); i < int64(len(matrix)); i++ {
		for j := int64(0); j < int64(len(matrix[0])); j++ {
			if matrix[i][j] == "#" {
				result++
			}
		}
	}
	return result
}

func executePlan(matrix [][]string, plan string) [][]string {

	if strings.Contains(plan, "x") {
		reg, _ := regexp.Compile("[^0-9]+")
		processedString := reg.ReplaceAllString(plan, "")
		xValue, _ := strconv.ParseInt(processedString, 10, 64)

		fmt.Printf("Xvalue: %v \n", xValue)

		var matrix1 [][]string
		var matrix2 [][]string

		for _, row := range matrix {
			fmt.Printf("row: %v \n", row)
			endRow := int64(len(matrix[0]))
			temp1 := row[0:xValue]
			temp2 := row[xValue:endRow]

			for i, j := 0, len(temp2)-1; i < j; i, j = i+1, j-1 {
				temp2[i], temp2[j] = temp2[j], temp2[i]
			}
			matrix1 = append(matrix1, temp1)
			matrix2 = append(matrix2, temp2)
		}

		fmt.Printf("Matrix 1: %v \n", matrix1)
		fmt.Printf("Matrix 2: %v \n", matrix2)

		for i := int64(0); i < int64(len(matrix1[0])); i++ {
			for j := int64(0); j < int64(len(matrix1)); j++ {
				if matrix2[i][j] == "#" {
					matrix1[i][j] = "#"
				}
			}
		}

		return matrix1

	} else if strings.Contains(plan, "y") {
		reg, _ := regexp.Compile("[^0-9]+")
		processedString := reg.ReplaceAllString(plan, "")
		yValue, _ := strconv.ParseInt(processedString, 10, 64)

		endMatrix := len(matrix)

		matrix1 := matrix[0:yValue]
		matrix2 := matrix[yValue:endMatrix]

		for i, j := 0, len(matrix2)-1; i < j; i, j = i+1, j-1 {
			matrix2[i], matrix2[j] = matrix2[j], matrix2[i]
		}

		for i := int64(0); i < int64(len(matrix1[0])); i++ {
			for j := int64(0); j < int64(len(matrix1)); j++ {
				if matrix2[i][j] == "#" {
					matrix1[i][j] = "#"
				}

			}
		}
		return matrix1
	}

	return matrix
}

func transparentOrigami(data []string) int64 {
	result := int64(0)
	xBound := int64(0)
	yBound := int64(0)

	var instructions []string
	var coords [][]int64
	var matrix [][]string

	for _, row := range data {
		if strings.Contains(row, "fold") {
			instructions = append(instructions, row)
		} else if len(row) > 0 {
			c := make([]int64, 2)
			vals := strings.Split(row, ",")
			xVal, _ := strconv.ParseInt(vals[0], 10, 64)
			yVal, _ := strconv.ParseInt(vals[1], 10, 64)
			c[0] = xVal
			c[1] = yVal
			coords = append(coords, c)
			if xVal > xBound {
				xBound = xVal
			}
			if yVal > yBound {
				yBound = yVal
			}
		}
	}
	xBound++
	yBound++

	for i := int64(0); i < yBound; i++ {
		row := make([]string, xBound)
		for j := int64(0); j < xBound; j++ {
			row[j] = "."
		}
		matrix = append(matrix, row)
	}

	for _, cd := range coords {
		matrix[cd[1]][cd[0]] = "#"
	}

	for _, plan := range instructions {
		matrix = executePlan(matrix, plan)
	}

	result = checkDots(matrix)

	return result
}
func TestDay13(t *testing.T) {
	sampleInput, _ := readStringLines("./inputs/13a.txt")
	// solveInput, _ := readStringLines("./inputs/13.txt")

	// PART 1a
	if diff := cmp.Diff(int64(17), transparentOrigami(sampleInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

}
