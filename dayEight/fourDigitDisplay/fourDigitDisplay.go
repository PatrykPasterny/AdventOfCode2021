package fourdigitdisplay

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CountEasyDigits(data []string) int {
	_, fourDigitOutput := GetInfoAboutSignal(data)

	var sum int
	for _, val := range fourDigitOutput {
		outputs := strings.Split(val, " ")
		for _, digit := range outputs {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				sum++
			}
		}
	}

	return sum
}

func GetAllSummedFourDigits(data []string) int {
	uniqueSignalPatterns, fourDigitOutput := GetInfoAboutSignal(data)

	var sum int
	for i, val := range uniqueSignalPatterns {
		uniqueSignals := strings.Split(val, " ")
		codeToValueMap := DecodeUniqueSignals(uniqueSignals)

		digitOutput := strings.Split(fourDigitOutput[i], " ")
		sb := strings.Builder{}
		for _, code := range digitOutput {
			for key := range codeToValueMap {
				isMatch := true
				if len(code) != len(key) {
					isMatch = false
					continue
				}

				for _, ch := range code {
					if !strings.Contains(key, string(ch)) {
						isMatch = false
						break
					}
				}

				if isMatch {
					sb.WriteString(fmt.Sprintf("%d", codeToValueMap[key]))
					break
				} else {
					continue
				}
			}
		}

		decodedValue, err := strconv.Atoi(sb.String())
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		sum += decodedValue
	}

	return sum
}

func DecodeUniqueSignals(signals []string) (codeToDigitMap map[string]int) {
	codeToDigitMap = make(map[string]int)
	digitToCodeMap := make(map[int]string)
	for _, code := range signals {
		if len(code) == 2 {
			digitToCodeMap[1] = code
			codeToDigitMap[code] = 1
		} else if len(code) == 3 {
			digitToCodeMap[7] = code
			codeToDigitMap[code] = 7
		} else if len(code) == 4 {
			digitToCodeMap[4] = code
			codeToDigitMap[code] = 4
		} else if len(code) == 7 {
			digitToCodeMap[8] = code
			codeToDigitMap[code] = 8
		}
	}

	for _, code := range signals {
		if _, ok := codeToDigitMap[code]; ok {
			continue
		}

		if len(code) == 5 {
			DecodeFiveDigitSignals(digitToCodeMap, codeToDigitMap, code)
		}
		if len(code) == 6 {
			DecodeSixDigitSignals(digitToCodeMap, codeToDigitMap, code)
		}
	}

	return
}

func DecodeFiveDigitSignals(digitToCodeMap map[int]string, codeToDigitMap map[string]int, code string) {
	var oneCode = digitToCodeMap[1]
	_, ok3 := digitToCodeMap[3]
	_, ok2 := digitToCodeMap[2]
	_, ok5 := digitToCodeMap[5]
	if !ok3 && strings.Contains(code, string(oneCode[0])) && strings.Contains(code, string(oneCode[1])) {
		digitToCodeMap[3] = code
		codeToDigitMap[code] = 3
		return
	}

	if !ok2 {
		fourCode := digitToCodeMap[4]
		sumOfSimilarities := 0
		for _, ch := range fourCode {
			if strings.Contains(code, string(ch)) {
				sumOfSimilarities++
			}
		}

		if sumOfSimilarities == 2 {
			digitToCodeMap[2] = code
			codeToDigitMap[code] = 2
			return
		}
	}

	if !ok5 {
		digitToCodeMap[5] = code
		codeToDigitMap[code] = 5
	}
}

func DecodeSixDigitSignals(digitToCodeMap map[int]string, codeToDigitMap map[string]int, code string) {
	var oneCode = digitToCodeMap[1]
	_, ok6 := digitToCodeMap[6]
	_, ok9 := digitToCodeMap[9]
	_, ok0 := digitToCodeMap[0]
	if !ok6 && ((strings.Contains(code, string(oneCode[0])) && !strings.Contains(code, string(oneCode[1]))) ||
		(!strings.Contains(code, string(oneCode[0])) && strings.Contains(code, string(oneCode[1])))) {
		digitToCodeMap[6] = code
		codeToDigitMap[code] = 6
		return
	}

	if !ok9 {
		fourCode := digitToCodeMap[4]
		sumOfSimilarities := 0
		for _, ch := range fourCode {
			if strings.Contains(code, string(ch)) {
				sumOfSimilarities++
			}
		}

		if sumOfSimilarities == 4 {
			digitToCodeMap[9] = code
			codeToDigitMap[code] = 9
			return
		}
	}

	if !ok0 {
		digitToCodeMap[0] = code
		codeToDigitMap[code] = 0
	}
}

func GetInfoAboutSignal(data []string) (uniqueSignalPatterns []string, fourDigitOutputs []string) {
	for _, val := range data {
		uniqueSignalPatterns = append(uniqueSignalPatterns, strings.Split(val, " | ")[0])
		fourDigitOutputs = append(fourDigitOutputs, strings.Split(val, " | ")[1])
	}

	return
}
