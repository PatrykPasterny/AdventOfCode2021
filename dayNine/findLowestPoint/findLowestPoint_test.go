package findlowestpoint_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	findlowestpoint "github.com/PatrykPasterny/AdventOfCode2021/dayNine/findLowestPoint"
)

func TestGetSumOfRiskLevels(t *testing.T) {
	puzzleData, rowLength, colLength := datareaders.GetMatrixDataFromFile("pointsData_test.txt")
	correctValue := 15

	if correctValue != findlowestpoint.GetSumOfRiskLevels(puzzleData, rowLength, colLength) {
		t.Fatalf("The sum of risk levels of all low points on my heightmap should be equal to %d, but is %d.", correctValue, findlowestpoint.GetSumOfRiskLevels(puzzleData, rowLength, colLength))
	}
}

func TestGetThreeLargestBasins(t *testing.T) {
	puzzleData, rowLength, colLength := datareaders.GetMatrixDataFromFile("pointsData_test.txt")
	correctValue := 1134

	if correctValue != findlowestpoint.GetThreeLargestBasins(puzzleData, rowLength, colLength) {
		t.Fatalf("The sum of risk levels of all low points on my heightmap should be equal to %d, but is %d.", correctValue, findlowestpoint.GetThreeLargestBasins(puzzleData, rowLength, colLength))
	}
}
