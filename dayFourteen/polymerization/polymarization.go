package polymerization

import (
	"strings"
)

func CreatePolymerBruteForce(data []string, steps int) int {
	polymer := data[0]

	polymerCreationMap := make(map[string]string, 0)
	polymerElementsQuantityMap := make(map[string]int, 0)

	for idx := 2; idx < len(data); idx++ {
		polymerCreation := strings.Split(data[idx], " -> ")
		polymerCreationMap[polymerCreation[0]] = polymerCreation[1]
	}

	bruteForce(polymerCreationMap, polymerElementsQuantityMap, steps, polymer)
	highestElementAmount, lowestElementAmount := countHighestAndLowestElementQuantity(polymerElementsQuantityMap)

	return highestElementAmount - lowestElementAmount
}

func CreatePolymerBetterSolution(data []string, steps int) int {
	polymer := data[0]

	polymerCreationMap := make(map[string]string, 0)
	polymerElementsQuantityMap := make(map[string]int, 0)

	for idx := 2; idx < len(data); idx++ {
		polymerCreation := strings.Split(data[idx], " -> ")
		polymerCreationMap[polymerCreation[0]] = polymerCreation[1]
	}

	betterSolution(polymerCreationMap, polymerElementsQuantityMap, steps, polymer)
	highestElementAmount, lowestElementAmount := countHighestAndLowestElementQuantity(polymerElementsQuantityMap)

	return highestElementAmount - lowestElementAmount
}

func CreatePolymerBestSolution(data []string, steps int) int {
	polymer := data[0]

	polymerCreationMap := make(map[string]string, 0)
	polymerElementsQuantityMap := make(map[string]int, 0)

	for idx := 2; idx < len(data); idx++ {
		polymerCreation := strings.Split(data[idx], " -> ")
		polymerCreationMap[polymerCreation[0]] = polymerCreation[1]
	}

	bestSolution(polymerCreationMap, polymerElementsQuantityMap, steps, polymer)
	highestElementAmount, lowestElementAmount := countHighestAndLowestElementQuantity(polymerElementsQuantityMap)

	return highestElementAmount - lowestElementAmount
}

func countHighestAndLowestElementQuantity(polymerElementsQuantityMap map[string]int) (highestElementAmount int, lowestElementAmount int) {
	lowestElementAmount = 1000000000000000000

	for _, v := range polymerElementsQuantityMap {
		if v > highestElementAmount {
			highestElementAmount = v
		}

		if v < lowestElementAmount {
			lowestElementAmount = v
		}
	}

	return
}

func bruteForce(polymerCreationMap map[string]string, polymerElementsQuantityMap map[string]int, steps int, polymer string) string {
	var newPolymerSb strings.Builder
	for steps > 0 {

		for idx := 1; idx < len(polymer); idx++ {
			newPolymerPart := polymerCreationMap[string(polymer[idx-1])+string(polymer[idx])]
			newPolymerSb.WriteString(string(polymer[idx-1]) + newPolymerPart)
			if steps == 1 {
				polymerElementsQuantityMap[string(polymer[idx-1])]++
				polymerElementsQuantityMap[string(newPolymerPart[0])]++
			}
		}

		newPolymerSb.WriteString(string(polymer[len(polymer)-1]))

		if steps == 1 {
			polymerElementsQuantityMap[string(polymer[len(polymer)-1])]++
		}
		polymer = newPolymerSb.String()
		newPolymerSb.Reset()

		steps--
	}

	return polymer
}

func betterSolution(polymerCreationMap map[string]string, polymerElementsQuantityMap map[string]int, steps int, polymer string) {
	for idx := 1; idx < len(polymer); idx++ {
		polymerElementsQuantityMap[string(polymer[idx-1])]++
		processPair(string(polymer[idx-1]), string(polymer[idx]), steps, polymerCreationMap, polymerElementsQuantityMap)
	}

	polymerElementsQuantityMap[string(polymer[len(polymer)-1])]++
}

func bestSolution(polymerCreationMap map[string]string, polymerElementsQuantityMap map[string]int, steps int, polymer string) {
	for idx := 0; idx < len(polymer); idx++ {
		polymerElementsQuantityMap[string(polymer[idx])]++
	}

	mapOfPairs := make(map[string]int, 0)
	for idx := 1; idx < len(polymer); idx++ {
		mapOfPairs[string(polymer[idx-1])+string(polymer[idx])]++
	}

	for steps > 0 {
		newMapOfPairs := make(map[string]int, 0)
		for k, val := range mapOfPairs {
			newVal := polymerCreationMap[k]
			leftPair := string(k[0]) + newVal
			rightPair := newVal + string(k[1])
			polymerElementsQuantityMap[newVal] += val
			newMapOfPairs[leftPair] += val
			newMapOfPairs[rightPair] += val
		}
		mapOfPairs = newMapOfPairs
		steps--
	}
}

func processPair(firstPairElement string, secPairElement string, steps int, polymerCreationMap map[string]string, polymerElementsQuantityMap map[string]int) {

	if steps == 0 {
		return
	}

	newElement := polymerCreationMap[firstPairElement+secPairElement]
	polymerElementsQuantityMap[newElement]++

	processPair(firstPairElement, newElement, steps-1, polymerCreationMap, polymerElementsQuantityMap)
	processPair(newElement, secPairElement, steps-1, polymerCreationMap, polymerElementsQuantityMap)
}
