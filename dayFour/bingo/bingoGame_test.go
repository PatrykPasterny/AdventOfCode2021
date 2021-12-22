package bingo_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayFour/bingo"
)

func TestGetBingoWinner(t *testing.T) {
	puzzleData := datareaders.GetBingoDataFromFile("bingoData_test.txt")
	correctValue := 4512

	if correctValue != bingo.GetBingoWinner(puzzleData) {
		t.Fatalf("The bingo game winning board sum of all unmarked numbers multiplied by the number that was just called when this board won should be equal to %d, but is %d.", correctValue, bingo.GetBingoWinner(puzzleData))
	}
}

func TestGetLastBingoWinner(t *testing.T) {
	puzzleData := datareaders.GetBingoDataFromFile("bingoData_test.txt")
	correctValue := 1924

	if correctValue != bingo.GetLastBingoWinner(puzzleData) {
		t.Fatalf("The bingo game last winning board sum of all unmarked numbers multiplied by the number that was just called when this board won should be equal to %d, but is %d.", correctValue, bingo.GetLastBingoWinner(puzzleData))
	}
}
