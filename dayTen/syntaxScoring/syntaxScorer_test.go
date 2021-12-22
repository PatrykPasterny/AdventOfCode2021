package syntaxscoring_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	syntaxscoring "github.com/PatrykPasterny/AdventOfCode2021/dayTen/syntaxScoring"
)

func TestScoreCorruptedLines(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("syntax_test.txt")
	firstCorrectValue := 26397
	secondCorrectValue := 288957

	firstValue, secondValue := syntaxscoring.ScoreCorruptedLines(puzzleData)

	if firstCorrectValue != firstValue {
		t.Fatalf("The incorrect score counted from given syntax data should be equal to %d, but is %d.", firstCorrectValue, firstValue)
	}

	if secondCorrectValue != secondValue {
		t.Fatalf("The incomplete score counted from given syntax data should be equal to %d, but is %d.", secondCorrectValue, secondValue)
	}
}
