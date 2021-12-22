package daytwelve

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	pathfinder "github.com/PatrykPasterny/AdventOfCode2021/dayTwelve/pathFinder"
)

func DayTwelve(filename string) {
	puzzleData := datareaders.GetDataFromFile(filename)

	fmt.Printf("First part: The number of all exit paths is equal to %d.", pathfinder.FindAllExitPathsForSingleSmallCaveEntrance(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: The step during which all octopuses have flashed is equal to %d.", pathfinder.FindAllExitPathsForDoubleSmallCaveEntrance(puzzleData))
}
