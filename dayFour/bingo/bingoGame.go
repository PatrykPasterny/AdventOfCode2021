package bingo

import (
	"github.com/PatrykPasterny/AdventOfCode2021/dayFour/bingo/models"
)

type MatrixPoint struct {
	idx int
	jdx int
}

func GetBingoWinner(data models.BingoGame) int {
	return FindNthBingoWinner(data, 1)
}

func GetLastBingoWinner(data models.BingoGame) int {
	return FindNthBingoWinner(data, len(data.Boards))
}

func FindNthBingoWinner(data models.BingoGame, n int) int {
	dataToManipulate := data

	for _, val := range dataToManipulate.BingoNumbers {
		winningBoardIdxs := make([]int, 0)

		for idx, board := range dataToManipulate.Boards {
			markedIndices := FindIndicesToMark(board.Board, val)

			if len(markedIndices) > 0 {
				MarkFoundIndices(board.Marked, markedIndices)
				if IsBoardWinning(board.Marked) {
					n -= 1
					if n == 0 {
						sumOfUnmarkedNumbers := GetSumOfUnmarkedNumbers(board)

						return sumOfUnmarkedNumbers * val
					} else {
						winningBoardIdxs = append(winningBoardIdxs, idx)
					}
				}
			}
		}

		if len(winningBoardIdxs) > 0 {
			deletedIdxs := 0
			for i := range winningBoardIdxs {
				dataToManipulate.Boards = append(dataToManipulate.Boards[:winningBoardIdxs[i]-deletedIdxs], dataToManipulate.Boards[winningBoardIdxs[i]-deletedIdxs+1:]...)
				deletedIdxs++
			}
		}
	}

	return -1
}

func GetSumOfUnmarkedNumbers(board models.BingoBoard) (sum int) {
	for i := range board.Marked {
		for j := range board.Marked[i] {
			if board.Marked[i][j] == 0 {
				sum += board.Board[i][j]
			}
		}
	}

	return
}

func IsBoardWinning(boards [][]int) bool {
	for j := range boards[0] {
		if boards[0][j] == 1 {
			isColWinning := true
			for i := 1; i < len(boards); i++ {
				if boards[i][j] != 1 {
					isColWinning = false
					break
				}
			}

			if isColWinning {
				return true
			}
		}

	}

	for i := range boards {
		if boards[i][0] == 1 {
			isRowWinning := true
			for j := 1; j < len(boards[i]); j++ {
				if boards[i][j] != 1 {
					isRowWinning = false
					break
				}
			}

			if isRowWinning {
				return true
			}
		}
	}

	return false
}

func FindIndicesToMark(matrix [][]int, searchingValue int) (result []MatrixPoint) {
	result = make([]MatrixPoint, 0)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == searchingValue {
				result = append(result, MatrixPoint{idx: i, jdx: j})
			}
		}
	}

	return
}

func MarkFoundIndices(marked [][]int, markedIndices []MatrixPoint) {
	for _, indices := range markedIndices {
		marked[indices.idx][indices.jdx] = 1
	}
}
