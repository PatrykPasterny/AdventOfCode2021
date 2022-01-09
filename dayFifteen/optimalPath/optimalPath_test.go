package optimalpath_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	optimalpath "github.com/PatrykPasterny/AdventOfCode2021/dayFifteen/optimalPath"
)

func TestGetRiskLevelOfShortestPath(t *testing.T) {
	puzzleData, _, _ := datareaders.GetMatrixDataFromFile("cave_test.txt")
	correctValue := 40

	value := optimalpath.GetRiskLevelOfShortestPath(puzzleData)

	if correctValue != value {
		t.Fatalf("First part: The lowest risk level for given cave should be equal to %d, but is equal to %d.", correctValue, value)
	}
}

func TestGetRiskLevelOfShortestPathForExtendedMatrix(t *testing.T) {
	puzzleData, _, _ := datareaders.GetMatrixDataFromFile("cave.txt")
	correctValue := 2916

	value := optimalpath.GetRiskLevelOfShortestPathForExtendedMatrix(puzzleData)

	if correctValue != value {
		t.Fatalf("First part: The lowest risk level for extended given cave should be equal to %d, but is equal to %d.", correctValue, value)
	}
}
