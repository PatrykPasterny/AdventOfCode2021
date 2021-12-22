package scanners_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayFive/scanners"
)

func TestGetNumberOfOverlappingLines(t *testing.T) {
	data := datareaders.GetDataFromFile("hydrothermalVents_test.txt")
	correctValue := 5

	if correctValue != scanners.GetNumberOfOverlappingLinesWithoutDiagonals(data) {
		t.Fatalf("The number of overlapping hydrothermal vent lines without diagonal should be equal to %d, but is %d.", correctValue, scanners.GetNumberOfOverlappingLinesWithoutDiagonals(data))
	}
}

func TestGetNumberOfOverlappingLinesWithDiagonals(t *testing.T) {
	data := datareaders.GetDataFromFile("hydrothermalVents_test.txt")
	correctValue := 12

	if correctValue != scanners.GetNumberOfOverlappingLinesWithDiagonals(data) {
		t.Fatalf("The number of overlapping hydrothermal vent lines with diagonals should be equal to %d, but is %d.", correctValue, scanners.GetNumberOfOverlappingLinesWithDiagonals(data))
	}
}
