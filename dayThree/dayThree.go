package daythree

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayThree/diagnostic"
)

func DayThree(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)

	fmt.Printf("First part: Submarine power consumption equals %d.", diagnostic.CountPowerConsumption(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: Submarine power consumption equals %d.", diagnostic.CountLifeSupportRating(puzzleData))
}
