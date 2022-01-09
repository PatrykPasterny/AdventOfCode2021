package polymerization_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayFourteen/polymerization"
)

func TestCreatePolymerBruteForce(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("polymerization_test.txt")
	correctValue := 1588
	steps := 10
	value := polymerization.CreatePolymerBruteForce(puzzleData, steps)

	if correctValue != value {
		t.Fatalf("The differenece between the highest and lowest polymer element quantity after %d steps is equal to %d, but is equal to %d.", steps, correctValue, value)
	}

}

func TestCreatePolymerBetterSolution(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("polymerization_test.txt")
	correctValue := 1588
	steps := 10

	value := polymerization.CreatePolymerBetterSolution(puzzleData, steps)

	if correctValue != value {
		t.Fatalf("The differenece between the highest and lowest polymer element quantity after %d steps is equal to %d, but is equal to %d.", steps, correctValue, value)
	}
}

func TestCreatePolymerBestSolution(t *testing.T) {
	puzzleData := datareaders.GetDataFromFile("polymerization_test.txt")
	correctValue := 1588
	steps := 10

	value := polymerization.CreatePolymerBestSolution(puzzleData, steps)

	if correctValue != value {
		t.Fatalf("The differenece between the highest and lowest polymer element quantity after %d steps is equal to %d, but is equal to %d.", steps, correctValue, value)
	}

}
