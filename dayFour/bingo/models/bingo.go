package models

type BingoBoard struct {
	Board  [][]int
	Marked [][]int
}

type BingoGame struct {
	Boards       []BingoBoard
	BingoNumbers []int
}
