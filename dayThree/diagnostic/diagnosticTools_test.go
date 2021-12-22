package diagnostic_test

import (
	"testing"

	datareaders "github.com/PatrykPasterny/AdventOfCode2021/dataReaders"
	"github.com/PatrykPasterny/AdventOfCode2021/dayThree/diagnostic"
)

func TestCountPowerConsumption(t *testing.T) {
	testData := datareaders.GetDataFromFile("powerConsumption_test.txt")
	correctValue := 198

	if diagnostic.CountPowerConsumption(testData) != correctValue {
		t.Fatalf("The submarine power consumption should be equal to %d, but is %d.", correctValue, diagnostic.CountPowerConsumption(testData))
	}
}

func TestCountLifeSupportRating(t *testing.T) {
	testData := datareaders.GetDataFromFile("powerConsumption_test.txt")
	correctValue := 230

	if diagnostic.CountLifeSupportRating(testData) != correctValue {
		t.Fatalf("The submarine life support rating should be equal to %d, but is %d.", correctValue, diagnostic.CountLifeSupportRating(testData))
	}
}
