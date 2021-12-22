package daythirteen

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	foldmanual "github.com/PatrykPasterny/AdventOfCode2021/dayThirteen/foldManual"
)

func DayThirteen(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)

	dots, _ := foldmanual.FoldManualNTimes(puzzleData, 1)
	fmt.Printf("First part: The number of dots after one fold is equal to %d.", dots)
	fmt.Println()
	_, manual := foldmanual.FoldManualNTimes(puzzleData, 12)
	fmt.Println("Second part: The manual after all folds looks like that:")
	foldmanual.PrintManual(manual)
}
