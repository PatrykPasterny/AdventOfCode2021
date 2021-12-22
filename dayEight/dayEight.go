package dayeight

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	fourdigitdisplay "github.com/PatrykPasterny/AdventOfCode2021/dayEight/fourDigitDisplay"
)

func DayEight(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)

	fmt.Printf("First part: The amount of easy digits in the given signal should equal be to %d.", fourdigitdisplay.CountEasyDigits(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: The sum of decoded four digits in the given signal should be equal to %d.", fourdigitdisplay.GetAllSummedFourDigits((puzzleData)))
}
