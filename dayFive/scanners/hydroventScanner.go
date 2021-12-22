package scanners

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type MatrixPoint struct {
	idx int
	jdx int
}

type LineCoordinates struct {
	lineStart MatrixPoint
	lineEnd   MatrixPoint
}

func GetNumberOfOverlappingLinesWithoutDiagonals(data []string) int {
	lines, maxIdx, maxJdx := GetHydrothermalVentLines(data, false)

	pathMatrix := CreatePathMatrix(maxIdx, maxJdx)

	for _, val := range lines {
		for idx := Min(val.lineStart.idx, val.lineEnd.idx); idx <= Max(val.lineStart.idx, val.lineEnd.idx); idx++ {
			for jdx := Min(val.lineStart.jdx, val.lineEnd.jdx); jdx <= Max(val.lineStart.jdx, val.lineEnd.jdx); jdx++ {
				pathMatrix[idx][jdx] += 1
			}
		}
	}

	numberOfOverlappingLines := CountOverlappingLines(pathMatrix)

	return numberOfOverlappingLines
}

func GetNumberOfOverlappingLinesWithDiagonals(data []string) int {
	lines, maxIdx, maxJdx := GetHydrothermalVentLines(data, true)

	pathMatrix := CreatePathMatrix(maxIdx, maxJdx)

	for _, val := range lines {
		if val.lineStart.idx == val.lineEnd.idx || val.lineStart.jdx == val.lineEnd.jdx {
			for idx := Min(val.lineStart.idx, val.lineEnd.idx); idx <= Max(val.lineStart.idx, val.lineEnd.idx); idx++ {
				for jdx := Min(val.lineStart.jdx, val.lineEnd.jdx); jdx <= Max(val.lineStart.jdx, val.lineEnd.jdx); jdx++ {
					pathMatrix[idx][jdx] += 1
				}
			}
		} else {
			if val.lineStart.idx < val.lineEnd.idx {
				jdx := val.lineStart.jdx
				for idx := val.lineStart.idx; idx <= val.lineEnd.idx; idx++ {
					pathMatrix[idx][jdx] += 1
					if val.lineStart.jdx < val.lineEnd.jdx {
						jdx++
					} else {
						jdx--
					}
				}
			} else {
				jdx := val.lineStart.jdx
				for idx := val.lineStart.idx; idx >= val.lineEnd.idx; idx-- {
					pathMatrix[idx][jdx] += 1
					if val.lineStart.jdx < val.lineEnd.jdx {
						jdx++
					} else {
						jdx--
					}
				}
			}
		}
	}

	numberOfOverlappingLines := CountOverlappingLines(pathMatrix)

	return numberOfOverlappingLines
}

func CreatePathMatrix(maxIdx int, maxJdx int) (pathMatrix [][]int) {
	pathMatrix = make([][]int, maxIdx+1)
	for i := range pathMatrix {
		pathMatrix[i] = make([]int, maxJdx+1)
	}

	return
}

func CountOverlappingLines(pathMatrix [][]int) (numberOfOverlappingLines int) {
	numberOfOverlappingLines = 0
	for idx := range pathMatrix {
		for jdx := range pathMatrix[idx] {
			if pathMatrix[idx][jdx] > 1 {
				numberOfOverlappingLines++
			}
		}
	}

	return
}

func GetHydrothermalVentLines(data []string, countDiagonals bool) (lines []LineCoordinates, maxIdx int, maxJdx int) {
	lines = make([]LineCoordinates, 0)

	maxIdx = 0
	maxJdx = 0

	for _, val := range data {
		lineCoordinatesAsString := strings.Split(strings.Trim(val, " \n"), " -> ")
		lineStartAsString := strings.Split(lineCoordinatesAsString[0], ",")
		lineEndAsString := strings.Split(lineCoordinatesAsString[1], ",")
		lineStartIdx := ConvertStringToInt(lineStartAsString[0])
		lineStartJdx := ConvertStringToInt(lineStartAsString[1])
		lineEndIdx := ConvertStringToInt(lineEndAsString[0])
		lineEndJdx := ConvertStringToInt(lineEndAsString[1])

		maxIdx = Max(maxIdx, lineStartIdx)
		maxIdx = Max(maxIdx, lineEndIdx)
		maxJdx = Max(maxJdx, lineStartJdx)
		maxJdx = Max(maxJdx, lineEndJdx)

		if !countDiagonals && lineStartIdx != lineEndIdx && lineStartJdx != lineEndJdx {
			continue
		}

		lineCoordinates := LineCoordinates{lineStart: MatrixPoint{idx: lineStartIdx, jdx: lineStartJdx}, lineEnd: MatrixPoint{idx: lineEndIdx, jdx: lineEndJdx}}

		lines = append(lines, lineCoordinates)
	}

	return
}

func ConvertStringToInt(val string) (res int) {
	res, err := strconv.Atoi(val)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	return
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
