package diagnostic

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func CountPowerConsumption(data []string) int {
	if len(data) == 0 {
		return 0
	}

	sumArray := GetSummedBitsArray(data)

	var sb strings.Builder
	var xorHelperBuilder strings.Builder

	for idx := 0; idx < len(sumArray); idx++ {
		if sumArray[idx] > len(data)/2 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}

		xorHelperBuilder.WriteString("1")
	}

	xorHelper := ParseBitStringToInt(xorHelperBuilder.String())

	gamma := ParseBitStringToInt(sb.String())
	epsilon := gamma ^ xorHelper

	return int(gamma * epsilon)
}

func CountLifeSupportRating(data []string) int {
	if len(data) == 0 {
		return 0
	}

	lineLength := len(data[0])
	oxygenGeneratorData := FilterLifeSupportingData(lineLength, data, "1", "0")
	co2scubberData := FilterLifeSupportingData(lineLength, data, "0", "1")

	oxygenGeneratorRating := oxygenGeneratorData[0]
	co2scubberRating := co2scubberData[0]

	return ParseBitStringToInt(oxygenGeneratorRating) * ParseBitStringToInt(co2scubberRating)
}

func FilterLifeSupportingData(lineLength int, data []string, leadingBit string, secondaryBit string) []string {
	for idx := 0; idx < lineLength; idx++ {
		if len(data) == 1 {
			break
		}

		currentIdxSum := 0
		for _, val := range data {
			currentIdxSum += GetSingleBitFromStringByIdx(val, idx)
		}

		if float64(currentIdxSum) > (float64(len(data))/2.0)-0.0001 {
			data = FilterNumbersWithNOnCurrentIdx(data, leadingBit, idx)
		} else {
			data = FilterNumbersWithNOnCurrentIdx(data, secondaryBit, idx)
		}
	}

	return data
}

func FilterNumbersWithNOnCurrentIdx(array []string, bit string, idx int) []string {
	filteredArray := []string{}

	for i := range array {
		if string(array[i][idx]) == bit {
			filteredArray = append(filteredArray, array[i])
		}
	}

	return filteredArray
}

func GetSummedBitsArray(data []string) []int {
	sumArray := make([]int, len(strings.Trim(data[0], " \n")))

	for _, val := range data {
		for idx := 0; idx < len(strings.Trim(val, " \n")); idx++ {
			sumArray[idx] += GetSingleBitFromStringByIdx(val, idx)
		}
	}

	return sumArray
}

func GetSingleBitFromStringByIdx(val string, idx int) int {
	bit, err := strconv.Atoi(string(val[idx]))
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return bit
}

func ParseBitStringToInt(s string) int {
	xorHelper, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return int(xorHelper)
}
