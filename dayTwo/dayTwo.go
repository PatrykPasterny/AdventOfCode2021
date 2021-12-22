package daytwo

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayTwo/move"
)

func DayTwo(fileName string) {
	puzzleData := datareaders.GetDataFromFile(fileName)

	fmt.Printf("First part: Submarine depth multiplied by its horizontal position equals %d.", move.MoveSubmarine(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: Submarine depth multiplied by its horizontal position equals %d.", move.MoveSubmarineIncludingAim(puzzleData))
}
