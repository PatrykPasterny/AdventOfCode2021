package octopusnavigation_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	octopusnavigation "github.com/PatrykPasterny/AdventOfCode2021/dayEleven/octopusNavigation"
)

func TestCountOctopusFlashesAfterNSteps(t *testing.T) {
	puzzleData, rowLength, colLength := datareaders.GetMatrixDataFromFile("initialOctopusState_test.txt")
	correctValue := 1656
	value := octopusnavigation.CountOctopusFlashesAfterNSteps(puzzleData, 100, rowLength, colLength)

	if correctValue != value {
		t.Fatalf("The number of octopuses that flashed during 100 steps should be equal to %d, but is %d.", correctValue, value)
	}
}

func TestGetStepWhenAllOctopusesFlash(t *testing.T) {
	puzzleData, rowLength, colLength := datareaders.GetMatrixDataFromFile("initialOctopusState_test.txt")
	correctValue := 195
	value := octopusnavigation.GetStepWhenAllOctopusesFlash(puzzleData, rowLength, colLength)

	if correctValue != value {
		t.Fatalf("The step during which all octopuses have flashed should be equal to %d, but is %d.", correctValue, value)
	}
}
