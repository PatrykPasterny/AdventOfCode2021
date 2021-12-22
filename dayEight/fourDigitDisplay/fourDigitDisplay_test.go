package fourdigitdisplay_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	fourdigitdisplay "github.com/PatrykPasterny/AdventOfCode2021/dayEight/fourDigitDisplay"
)

func TestCountEasyDigits(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("signal_test.txt")
	correctValue := 26
	if correctValue != fourdigitdisplay.CountEasyDigits(puzzleData) {
		t.Fatalf("The amount of easy digits in the given signal should be equal to %d, but is %d.", correctValue, fourdigitdisplay.CountEasyDigits(puzzleData))
	}
}

func TestGetAllSummedFourDigits(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("signal_test.txt")
	correctValue := 61229
	if correctValue != fourdigitdisplay.GetAllSummedFourDigits(puzzleData) {
		t.Fatalf("The sum of all four digits in the given signal should be equal to %d, but is %d.", correctValue, fourdigitdisplay.GetAllSummedFourDigits(puzzleData))
	}
}
