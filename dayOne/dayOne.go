package dayOne

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayOne/sonar"
)

func DayOne(fileName string) {
	var puzzleData = datareaders.GetIntDataFromFile(fileName)

	fmt.Printf("First part: Submarine depth increases number is equal to %d.", sonar.CountDepthIncreases(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: Submarine depth increases considering noise is equal to %d.", sonar.CountDepthIncreasesWithNoise(puzzleData))
}
