package sonar_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayOne/sonar"
)

func TestCountDepthIncreases(t *testing.T) {
	data := datareaders.GetIntDataFromFile("depthData_test.txt")
	correctResult := 7

	if correctResult != sonar.CountDepthIncreases(data) {
		t.Fatalf("Depth increases for test data should be: %d, but is %d.", correctResult, sonar.CountDepthIncreases(data))
	}
}

func TestCountDepthIncreasesWithNoise(t *testing.T) {
	data := datareaders.GetIntDataFromFile("depthData_test.txt")
	correctResult := 5

	if correctResult != sonar.CountDepthIncreasesWithNoise(data) {
		t.Fatalf("Depth increases for test data should be: %d, but is %d.", correctResult, sonar.CountDepthIncreasesWithNoise(data))
	}
}
