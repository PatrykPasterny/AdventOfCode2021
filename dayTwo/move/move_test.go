package move_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayTwo/move"
)

func TestMoveSumbarine(t *testing.T) {
	testData := datareaders.GetDataFromFile("move_test.txt")
	correctValue := 150

	if correctValue != move.MoveSubmarine(testData) {
		t.Fatalf("The correct value should be %d, but is %d. ", correctValue, move.MoveSubmarine(testData))
	}
}

func TestMoveSubmarineIncludingAim(t *testing.T) {
	testData := datareaders.GetDataFromFile("move_test.txt")
	correctValue := 900

	if correctValue != move.MoveSubmarineIncludingAim(testData) {
		t.Fatalf("The correct value should be %d, but is %d.", correctValue, move.MoveSubmarineIncludingAim(testData))
	}
}
