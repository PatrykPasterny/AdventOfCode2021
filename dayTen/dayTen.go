package dayten

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	syntaxscoring "github.com/PatrykPasterny/AdventOfCode2021/dayTen/syntaxScoring"
)

func DayTen(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)
	firstResult, secondResult := syntaxscoring.ScoreCorruptedLines(puzzleData)

	fmt.Printf("First part: The incorrect score counted from given syntax data is equal to %d.", firstResult)
	fmt.Println()
	fmt.Printf("Second part: The incomplete score counted from given syntax data is equal to %d.", secondResult)
}
