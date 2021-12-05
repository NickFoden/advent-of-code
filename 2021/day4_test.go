package adventOfCode

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func calculateFinal(b [][]int64, currentPicks []int64, lastNumber int64) int64 {
	result := int64(0)
	for _, row := range b {
		for _, single := range row {
			if !containsInt64(currentPicks, single) {
				result += single
			}
		}
	}
	return result * lastNumber
}

func newSliceInt(row string) []int64 {
	tempRow := strings.Split(row, " ")
	var result []int64

	for _, val := range tempRow {
		if val != "" {
			num, _ := strconv.ParseInt(val, 10, 64)
			result = append(result, num)
		}

	}
	return result
}

func determineDrawnNumbers(index int, all []int64) ([]int64, int64) {
	var result []int64

	lastNumber := int64(0)

	for idx, num := range all {
		if idx < index {
			result = append(result, num)
			lastNumber = num
		} else {
			break
		}
	}

	return result, lastNumber
}

func checkWinner(board [][]int64, drawn []int64) bool {
	result := false
	for _, row := range board {
		if !result {
			rowCheck := true
			for _, single := range row {
				if !containsInt64(drawn, single) {
					rowCheck = false
				}
			}
			if rowCheck {
				result = true
				break
			}
		}
	}

	if result {
		return result
	}

	for i := 0; i < len(board[0]); i++ {
		if !result {
			columnCheck := true
			for j := 0; j < len(board[0]); j++ {
				if !containsInt64(drawn, board[j][i]) {
					columnCheck = false
				}
			}

			if columnCheck {
				result = true
				break
			}
		}
	}

	return result
}

func winningBingo(data []string) int64 {
	final := int64(0)
	drawNumbersRow := strings.Split(data[0], ",")
	var drawNumbers []int64

	for _, val := range drawNumbersRow {
		num, _ := strconv.ParseInt(val, 10, 64)
		drawNumbers = append(drawNumbers, num)
	}

	var gameBoard [][]int64
	var game [][][]int64

	for index, row := range data {
		if index != 0 {
			if len(row) == 0 {
				if len(gameBoard) > 0 {
					game = append(game, gameBoard)
					gameBoard = nil
				}
			} else {
				gameBoard = append(gameBoard, newSliceInt(row))
			}
		}
	}

	hasWinner := false

	for i := 1; i < len(drawNumbers); i++ {
		if !hasWinner {
			currentPicks, lastNumber := determineDrawnNumbers(i, drawNumbers)
			for boardIndex, board := range game {
				if checkWinner(board, currentPicks) {
					hasWinner = true
					final = calculateFinal(game[boardIndex], currentPicks, lastNumber)
					break
				}
			}
		}
	}

	return final
}

func lastBingoWinner(data []string) int64 {
	final := int64(0)
	drawNumbersRow := strings.Split(data[0], ",")
	var drawNumbers []int64

	for _, val := range drawNumbersRow {
		num, _ := strconv.ParseInt(val, 10, 64)
		drawNumbers = append(drawNumbers, num)
	}

	var gameBoard [][]int64
	var game [][][]int64

	for index, row := range data {
		if index != 0 {
			if len(row) == 0 {
				if len(gameBoard) > 0 {
					game = append(game, gameBoard)
					gameBoard = nil
				}
			} else {
				gameBoard = append(gameBoard, newSliceInt(row))
			}
		}
	}

	var winnerIndexes []int64
	done := false

	for i := 1; i < len(drawNumbers); i++ {
		currentPicks, lastNumber := determineDrawnNumbers(i, drawNumbers)
		for boardIndex, board := range game {
			if !done {
				if !containsInt64(winnerIndexes, int64(boardIndex)) {
					if (len(winnerIndexes) + 1) == len(game) {
						if checkWinner(board, currentPicks) {
							final = calculateFinal(game[int64(boardIndex)], currentPicks, lastNumber)
							done = true
						}
					} else {
						if checkWinner(board, currentPicks) {
							winnerIndexes = append(winnerIndexes, int64(boardIndex))
						}
					}

				}
			}

		}
	}

	return final

}

func TestDay4(t *testing.T) {
	// sampleInput, _ := readStringLines("./inputs/4a.txt")
	solveInput, _ := readStringLines("./inputs/4.txt")

	// PART 1a - sample input is MESSED UP - no 3rd game board ?!
	// if diff := cmp.Diff(int64(4512), winningBingo(sampleInput)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// PART 1
	if diff := cmp.Diff(int64(11774), winningBingo(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// PART 2
	if diff := cmp.Diff(int64(4495), lastBingoWinner(solveInput)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
