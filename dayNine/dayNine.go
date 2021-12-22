package daynine

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	findlowestpoint "github.com/PatrykPasterny/AdventOfCode2021/dayNine/findLowestPoint"
)

func DayNine(filename string) {
	puzzleData, rowLength, colLength := datareaders.GetMatrixDataFromFile(filename)

	fmt.Printf("First part: Sum of the risk levels of all low points on submarine's heightmap is equal to %d.", findlowestpoint.GetSumOfRiskLevels(puzzleData, rowLength, colLength))
	fmt.Println()
	fmt.Printf("Second part: Sum of the three low points deepest basins is equal to %d.", findlowestpoint.GetThreeLargestBasins(puzzleData, rowLength, colLength))
}
