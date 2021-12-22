package dayeleven

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	octopusnavigation "github.com/PatrykPasterny/AdventOfCode2021/dayEleven/octopusNavigation"
)

func DayEleven(filename string) {
	puzzleDataFirst, rowLength, colLength := datareaders.GetMatrixDataFromFile(filename)
	fmt.Printf("First part: The number of octopuses that flashed during 100 steps is equal to %d.", octopusnavigation.CountOctopusFlashesAfterNSteps(puzzleDataFirst, 100, rowLength, colLength))
	fmt.Println()

	puzzleDataSecond, rowLength, colLength := datareaders.GetMatrixDataFromFile(filename)
	fmt.Printf("Second part: The step during which all octopuses have flashed is equal to %d.", octopusnavigation.GetStepWhenAllOctopusesFlash(puzzleDataSecond, rowLength, colLength))
}
