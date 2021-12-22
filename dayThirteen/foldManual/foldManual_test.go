package foldmanual_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	foldmanual "github.com/PatrykPasterny/AdventOfCode2021/dayThirteen/foldManual"
)

func TestFoldManualOneTimes(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("manual_test.txt")
	correctDataAfterOneFold := 17
	valueAfterOneFold, _ := foldmanual.FoldManualNTimes(puzzleData, 1)

	if correctDataAfterOneFold != valueAfterOneFold {
		t.Fatalf("The number of dots after one fold should be equal to %d, but is equal to %d.", correctDataAfterOneFold, valueAfterOneFold)
	}
}
