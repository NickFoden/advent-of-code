package adventOfCode

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const ROCK = "Rock"
const PAPER = "Paper"
const SCISSORS = "Scissors"

const WIN = "Win"
const DRAW = "Draw"
const LOST = "Lose"

func calculateMove(g string) string {
	switch g {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	}
	return ""
}

func calculatePointsForChoice(p string) int64 {
	switch p {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	}
	return 0
}

func calculateMatch(p1 string, p2 string) int64 {
	score := int64(0)
	pointsForChoice := calculatePointsForChoice(p2)
	score = score + pointsForChoice

	// p2 loses
	if p1 == ROCK && p2 == SCISSORS || p1 == PAPER && p2 == ROCK || p1 == SCISSORS && p2 == PAPER {
		// 0 points for a loss
	} else if p1 == p2 {
		// 3 points for a draw
		score = score + 3
	} else {
		// 6 points for a win
		score = score + 6

	}
	return score
}

func calculateGameScore(g string) int64 {
	match := strings.Split(g, " ")

	player1 := calculateMove(match[0])
	player2 := calculateMove(match[1])

	score := calculateMatch(player1, player2)

	return score
}

func rockPapersStrategy(data []string) int64 {
	score := int64(0)

	for _, game := range data {
		gameScore := calculateGameScore(game)
		score = score + gameScore
	}

	return score
}

func determineOutcome(o string) string {
	switch o {
	case "X":
		return LOST
	case "Y":
		return DRAW
	case "Z":
		return WIN
	}
	return ""
}

func determinePlayer2(p1 string, out string) string {

	if out == DRAW {
		return p1
	}

	if p1 == ROCK {
		if out == WIN {
			return PAPER
		} else {
			return SCISSORS
		}
	}
	if p1 == PAPER {
		if out == WIN {
			return SCISSORS
		} else {
			return ROCK
		}
	}
	if p1 == SCISSORS {
		if out == WIN {
			return ROCK
		} else {
			return PAPER
		}
	}
	return ""
}

func rockPapersRequiredStrategy(data []string) int64 {
	score := int64(0)

	for _, game := range data {
		match := strings.Split(game, " ")
		player1 := calculateMove(match[0])
		outcome := determineOutcome(match[1])
		player2 := determinePlayer2(player1, outcome)
		gameScore := calculateMatch(player1, player2)
		score = score + gameScore
	}

	return score
}

func TestDay2(t *testing.T) {

	sampleInput, _ := readStringLines("./inputs/2sample.txt")
	puzzleInput, _ := readStringLines("./inputs/2.txt")

	// Part 1

	if diff := cmp.Diff(rockPapersStrategy(sampleInput), int64(15)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(rockPapersStrategy(puzzleInput), int64(10816)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	// Part 2

	if diff := cmp.Diff(rockPapersRequiredStrategy(sampleInput), int64(12)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(rockPapersRequiredStrategy(puzzleInput), int64(11657)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
}
