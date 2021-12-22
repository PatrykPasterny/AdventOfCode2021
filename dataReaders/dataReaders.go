package datareaders

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	bingo "github.com/PatrykPasterny/AdventOfCode2021/dayFour/bingo/models"
)

func GetDataFromFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]string, 0)
	for scanner.Scan() {
		value := scanner.Text()

		result = append(result, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func GetIntDataFromFile(fileName string) []int {
	stringData := GetDataFromFile(fileName)

	result := make([]int, 0)
	for _, val := range stringData {
		value, err := strconv.Atoi(val)
		if err != nil {
			// handle error
			log.Fatal(err)
			os.Exit(2)
		}

		result = append(result, value)
	}

	return result
}

func GetMatrixDataFromFile(fileName string) (result [][]int, rowLength int, colLength int) {
	stringData := GetDataFromFile(fileName)
	rowLength = len(stringData)
	colLength = len(stringData[0])

	result = make([][]int, rowLength)
	for i := range result {
		result[i] = make([]int, colLength)
		for j := range result[i] {
			singleVal, err := strconv.Atoi(string(stringData[i][j]))
			if err != nil {
				log.Fatal(err)
				os.Exit(2)
			}

			result[i][j] = singleVal
		}
	}

	return
}

func GetIntDataWithSeparatorFromFile(filename string, s string) []int {
	stringData := GetDataFromFile(filename)

	result := make([]int, 0)
	for _, val := range strings.Split(stringData[0], s) {
		value, err := strconv.Atoi(val)
		if err != nil {
			// handle error
			log.Fatal(err)
			os.Exit(2)
		}

		result = append(result, value)
	}

	return result
}

func GetBingoDataFromFile(filename string) bingo.BingoGame {
	resultBingoGame := bingo.BingoGame{Boards: []bingo.BingoBoard{}, BingoNumbers: make([]int, 0)}
	dataToParse := GetDataFromFile(filename)

	bingoNumbersAsString := strings.Split(dataToParse[0], ",")
	for _, numAsString := range bingoNumbersAsString {
		num, err := strconv.Atoi(numAsString)
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		resultBingoGame.BingoNumbers = append(resultBingoGame.BingoNumbers, num)
	}

	newBingoNumbersBoard := make([][]int, 0)
	singleNumberBoardRowIndex := 0
	for idx := 2; idx < len(dataToParse); idx++ {
		if len(strings.Trim(dataToParse[idx], " \n")) == 0 {
			AddNewBingoBoard(&resultBingoGame, newBingoNumbersBoard)
			newBingoNumbersBoard = make([][]int, 0)
			singleNumberBoardRowIndex = 0
			continue
		}

		bingoNumbersBoardRow := strings.Fields(dataToParse[idx])
		newBingoNumbersBoard = append(newBingoNumbersBoard, make([]int, 0))
		for i := range bingoNumbersBoardRow {
			bingoNumberBoardElement, err := strconv.Atoi(strings.Trim(bingoNumbersBoardRow[i], " "))
			if err != nil {
				log.Fatal(err)
				os.Exit(2)
			}

			newBingoNumbersBoard[singleNumberBoardRowIndex] = append(newBingoNumbersBoard[singleNumberBoardRowIndex], bingoNumberBoardElement)
		}

		singleNumberBoardRowIndex++
	}

	AddNewBingoBoard(&resultBingoGame, newBingoNumbersBoard)

	return resultBingoGame
}

func AddNewBingoBoard(resultBingoGame *bingo.BingoGame, newBingoNumbersBoard [][]int) {
	newMarkedBoard := make([][]int, len(newBingoNumbersBoard))
	for i := range newMarkedBoard {
		newMarkedBoard[i] = make([]int, len(newBingoNumbersBoard[0]))
	}

	newBingoBoard := bingo.BingoBoard{Board: newBingoNumbersBoard, Marked: newMarkedBoard}
	resultBingoGame.Boards = append(resultBingoGame.Boards, newBingoBoard)
}
