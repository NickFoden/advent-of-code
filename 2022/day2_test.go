package adventOfCode

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const ROCK = "Rock"
const PAPER = "Paper"
const SCISSORS = "Scissors"

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

func TestDay2(t *testing.T) {

	sampleInput, _ := readStringLines("./inputs/2sample.txt")
	puzzleInput, _ := readStringLines("./inputs/2.txt")

	// Part 1

	if diff := cmp.Diff(rockPapersStrategy(sampleInput), int64(15)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(rockPapersStrategy(puzzleInput), int64(15)); diff != "" {
		t.Errorf("Value mismatch (-want +got):\n%s", diff)
	}
	// if diff := cmp.Diff(mostCalories(puzzleInput), int64(70116)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// Part 2

	// if diff := cmp.Diff(topThreeMostCalories(sampleInput), int64(45000)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }

	// if diff := cmp.Diff(topThreeMostCalories(puzzleInput), int64(206582)); diff != "" {
	// 	t.Errorf("Value mismatch (-want +got):\n%s", diff)
	// }
}
