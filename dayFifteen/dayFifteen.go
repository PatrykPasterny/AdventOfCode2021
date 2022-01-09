package dayfifteen

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	optimalpath "github.com/PatrykPasterny/AdventOfCode2021/dayFifteen/optimalPath"
)

func DayFifteen(filename string) {
	puzzleData, _, _ := datareaders.GetMatrixDataFromFile(filename)

	value := optimalpath.GetRiskLevelOfShortestPath(puzzleData)
	secValue := optimalpath.GetRiskLevelOfShortestPathForExtendedMatrix(puzzleData)
	fmt.Printf("First part: The lowest risk level for given cave is equal to %d.", value)
	fmt.Println()
	fmt.Printf("Second part: The lowest risk level for extended given cave is equal to %d.", secValue)

}
