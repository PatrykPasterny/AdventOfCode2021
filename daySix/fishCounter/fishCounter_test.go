package fishcounter_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	fishcounter "github.com/PatrykPasterny/AdventOfCode2021/daySix/fishCounter"
)

func TestGetLanternfishNumberAfterDays(t *testing.T) {
	puzzleData := datareaders.GetIntDataWithSeparatorFromFile("firstDayLanternfishes_test.txt", ",")
	correctValue18Days := 26
	correctValue80Days := 5934

	if correctValue18Days != fishcounter.GetLanternfishNumberAfterDays(puzzleData, 18) {
		t.Fatalf("The amount of lantern fishes after 18 days should be equal to %d, but is %d.", correctValue18Days, fishcounter.GetLanternfishNumberAfterDays(puzzleData, 18))
	}

	if correctValue80Days != fishcounter.GetLanternfishNumberAfterDays(puzzleData, 80) {
		t.Fatalf("The amount of lantern fishes after 80 days should be equal to %d, but is %d.", correctValue80Days, fishcounter.GetLanternfishNumberAfterDays(puzzleData, 80))
	}
}

func TestGetLanternfishNumberAfterDaysOptimal(t *testing.T) {
	puzzleData := datareaders.GetIntDataWithSeparatorFromFile("firstDayLanternfishes_test.txt", ",")
	correctValue18Days := 26
	correctValue80Days := 5934
	correctValue256Days := 26984457539

	if correctValue18Days != fishcounter.GetLanternfishNumberAfterDaysOptimal(puzzleData, 18) {
		t.Fatalf("The amount of lantern fishes after 18 days should be equal to %d, but is %d.", correctValue18Days, fishcounter.GetLanternfishNumberAfterDaysOptimal(puzzleData, 18))
	}

	if correctValue80Days != fishcounter.GetLanternfishNumberAfterDaysOptimal(puzzleData, 80) {
		t.Fatalf("The amount of lantern fishes after 80 days should be equal to %d, but is %d.", correctValue80Days, fishcounter.GetLanternfishNumberAfterDaysOptimal(puzzleData, 80))
	}

	if correctValue256Days != fishcounter.GetLanternfishNumberAfterDaysOptimal(puzzleData, 256) {
		t.Fatalf("The amount of lantern fishes after 256 days should be equal to %d, but is %d.", correctValue256Days, fishcounter.GetLanternfishNumberAfterDays(puzzleData, 256))
	}
}
