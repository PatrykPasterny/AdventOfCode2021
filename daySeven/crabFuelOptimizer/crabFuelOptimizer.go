package crabfueloptimizer

import (
	"math"
	"sort"
)

func GetOptimizeAmountOfFuelToAllignCrabs(data []int) int {
	sort.Ints(data)
	median := 0
	if len(data)%2 == 0 && len(data) > 0 {
		median = (data[len(data)/2] + data[(len(data)/2)-1]) / 2
	} else {
		median = data[(len(data)-1)/2]
	}
	potentialAnswers := []int{median, median + 1}
	firstPotentialResult := CountCrabFuelUsageForGivenAllignmentValue(data, potentialAnswers[0])
	secondPotentialResult := CountCrabFuelUsageForGivenAllignmentValue(data, potentialAnswers[1])

	return Min(firstPotentialResult, secondPotentialResult)
}

func GetOptimizeAmountOfFuelToAllignCrabsConsideringRaisingFuelBurn(data []int) int {
	average := SumArrayElements(data) / len(data)

	potentialAnswers := []int{average, average + 1}
	firstPotentialResult := CountCrabFuelUsageForGivenAllignmentValueWithChangingFuelBurningSystem(data, potentialAnswers[0])
	secondPotentialResult := CountCrabFuelUsageForGivenAllignmentValueWithChangingFuelBurningSystem(data, potentialAnswers[1])

	return Min(firstPotentialResult, secondPotentialResult)
}

func CountCrabFuelUsageForGivenAllignmentValue(data []int, value int) (usedFuel int) {
	for _, fuel := range data {
		usedFuel += int(math.Abs(float64(value - fuel)))
	}

	return
}

func CountCrabFuelUsageForGivenAllignmentValueWithChangingFuelBurningSystem(data []int, value int) (usedFuel int) {
	for _, fuel := range data {
		n := int(math.Abs(float64(value - fuel)))
		usedFuel += n * (n + 1) / 2
	}

	return
}

func SumArrayElements(data []int) (sum int) {
	for _, val := range data {
		sum += val
	}

	return
}

func Min(f int, s int) int {
	if f < s {
		return f
	} else {
		return s
	}
}
