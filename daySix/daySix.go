package daysix

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	fishcounter "github.com/PatrykPasterny/AdventOfCode2021/daySix/fishCounter"
)

func DaySix(filename string) {
	puzzleData := datareaders.GetIntDataWithSeparatorFromFile(filename, ",")

	fmt.Printf("First part: The number of lanternfishes after 80 days is equal to %d.", fishcounter.GetLanternfishNumberAfterDays(puzzleData, 80))
	fmt.Println()
	fmt.Printf("Second part: The number of lanternfishes after 256 days is equal to %d.", fishcounter.GetLanternfishNumberAfterDaysOptimal(puzzleData, 256))
}
