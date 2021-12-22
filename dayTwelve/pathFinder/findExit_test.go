package pathfinder_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	pathfinder "github.com/PatrykPasterny/AdventOfCode2021/dayTwelve/pathFinder"
)

func TestFindAllExitPathsForSingleSmallCaveEntrance(t *testing.T) {
	puzzleDataFirst := datareaders.GetDataFromFile("paths_test1.txt")
	firstCorrectValue := 10
	firstValue := pathfinder.FindAllExitPathsForSingleSmallCaveEntrance(puzzleDataFirst)

	puzzleDataSecond := datareaders.GetDataFromFile("paths_test2.txt")
	secondCorrectValue := 19
	secondValue := pathfinder.FindAllExitPathsForSingleSmallCaveEntrance(puzzleDataSecond)

	puzzleDataThird := datareaders.GetDataFromFile("paths_test3.txt")
	thirdCorrectValue := 226
	thirdValue := pathfinder.FindAllExitPathsForSingleSmallCaveEntrance(puzzleDataThird)

	if firstCorrectValue != firstValue {
		t.Fatalf("The number of all exit paths should be equal to %d, but is equal to %d.", firstCorrectValue, firstValue)
	}

	if secondCorrectValue != secondValue {
		t.Fatalf("The number of all exit paths should be equal to %d, but is equal to %d.", secondCorrectValue, secondValue)
	}

	if thirdCorrectValue != thirdValue {
		t.Fatalf("The number of all exit paths should be equal to %d, but is equal to %d.", thirdCorrectValue, thirdValue)
	}
}

func TestFindAllExitPathsForDoubleSmallCaveEntrance(t *testing.T) {
	puzzleDataFirst := datareaders.GetDataFromFile("paths_test1.txt")
	firstCorrectValue := 36
	firstValue := pathfinder.FindAllExitPathsForDoubleSmallCaveEntrance(puzzleDataFirst)

	puzzleDataSecond := datareaders.GetDataFromFile("paths_test2.txt")
	secondCorrectValue := 103
	secondValue := pathfinder.FindAllExitPathsForDoubleSmallCaveEntrance(puzzleDataSecond)

	puzzleDataThird := datareaders.GetDataFromFile("paths_test3.txt")
	thirdCorrectValue := 3509
	thirdValue := pathfinder.FindAllExitPathsForDoubleSmallCaveEntrance(puzzleDataThird)

	if firstCorrectValue != firstValue {
		t.Fatalf("The number of all exit paths should be equal to %d, but is equal to %d.", firstCorrectValue, firstValue)
	}

	if secondCorrectValue != secondValue {
		t.Fatalf("The number of all exit paths should be equal to %d, but is equal to %d.", secondCorrectValue, secondValue)
	}

	if thirdCorrectValue != thirdValue {
		t.Fatalf("The number of all exit paths should be equal to %d, but is equal to %d.", thirdCorrectValue, thirdValue)
	}
}
