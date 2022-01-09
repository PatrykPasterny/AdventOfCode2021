package dayfourteen

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayFourteen/polymerization"
)

func DayFourteen(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)
	steps := 40

	value := polymerization.CreatePolymerBestSolution(puzzleData, steps)
	fmt.Printf("First part: The differenece between the highest and lowest polymer element quantity after %d steps is equal to %d.", steps, value)
	fmt.Println()
}
