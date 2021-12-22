package dayfour

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayFour/bingo"
)

func DayFour(filename string) {
	puzzleData := datareaders.GetBingoDataFromFile(filename)

	fmt.Printf("First part: The bingo game winning board sum of all unmarked numbers multiplied by the number that was just called when this board won equals %d.", bingo.GetBingoWinner(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: The bingo game last winning board sum of all unmarked numbers multiplied by the number that was just called when this board won equals %d.", bingo.GetLastBingoWinner(puzzleData))
}
