package dayseven

import (
	"fmt"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	crabfueloptimizer "github.com/PatrykPasterny/AdventOfCode2021/daySeven/crabFuelOptimizer"
)

func DaySeven(filename string) {
	puzzleData := datareaders.GetIntDataWithSeparatorFromFile(filename, ",")

	fmt.Printf("First part: The crab allignment value that uses the less amount of fuel is equal to %d.", crabfueloptimizer.GetOptimizeAmountOfFuelToAllignCrabs(puzzleData))
	fmt.Println()
	fmt.Printf("Second part: The crab allignment value that uses the less amount of fuel taking the given fuel distribution is equal to %d.", crabfueloptimizer.GetOptimizeAmountOfFuelToAllignCrabsConsideringRaisingFuelBurn(puzzleData))
}
