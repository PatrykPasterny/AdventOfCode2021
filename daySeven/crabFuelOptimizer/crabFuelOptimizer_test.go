package crabfueloptimizer_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	crabfueloptimizer "github.com/PatrykPasterny/AdventOfCode2021/daySeven/crabFuelOptimizer"
)

func TestGetOptimizeAmountOfFuelToAllignCrabs(t *testing.T) {
	puzzleData := datareaders.GetIntDataWithSeparatorFromFile("crabPositions_test.txt", ",")
	correctValue := 37

	if correctValue != crabfueloptimizer.GetOptimizeAmountOfFuelToAllignCrabs(puzzleData) {
		t.Fatalf("The crab allignment value that uses the less amount of fuel should be %d, but is %d.", correctValue, crabfueloptimizer.GetOptimizeAmountOfFuelToAllignCrabs(puzzleData))
	}
}

func TestGetOptimizeAmountOfFuelToAllignCrabsConsideringRaisingFuelBurn(t *testing.T) {
	puzzleData := datareaders.GetIntDataWithSeparatorFromFile("crabPositions_test.txt", ",")
	correctValue := 168

	if correctValue != crabfueloptimizer.GetOptimizeAmountOfFuelToAllignCrabsConsideringRaisingFuelBurn(puzzleData) {
		t.Fatalf("The crab allignment value that uses the less amount of fuel should be %d, but is %d.", correctValue, crabfueloptimizer.GetOptimizeAmountOfFuelToAllignCrabsConsideringRaisingFuelBurn(puzzleData))
	}
}
