package dayfive

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayFive/scanners"
)

func DayFive(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)

	fmt.Printf("First part: The number of overlapping hydrothermal vent lines without diagonals is equal to %d.", scanners.GetNumberOfOverlappingLinesWithoutDiagonals(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: The number of overlapping hydrothermal vent lines with diagonals is equal to %d.", scanners.GetNumberOfOverlappingLinesWithDiagonals(puzzleData))
}
