package foldmanual

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	idx int
	jdx int
}

func FoldManualNTimes(data []string, n int) (dotNumber int, lastManual [][]string) {
	manual, folds := getManualAndFolds(data)
	idx := 0
	for idx < n {
		manual = FoldManual(manual, folds[idx])
		dotNumber += countDots(manual)
		idx++
	}

	return dotNumber, manual
}

func FoldManual(manual [][]string, fold Point) (foldedManual [][]string) {
	foldLengthIdx, foldLengthJdx := getFoldData(manual, fold)
	rewriteDots(manual, fold, foldLengthIdx, foldLengthJdx)
	foldedManual = createNewManualAfterFold(manual, fold)

	return
}

func getFoldData(manual [][]string, fold Point) (foldLengthIdx int, foldLengthJdx int) {
	if fold.idx > 0 {
		for jdx := 0; jdx < len(manual[fold.idx]); jdx++ {
			manual[fold.idx][jdx] = "-"
		}

		if fold.idx > len(manual)-fold.idx-1 {
			foldLengthIdx = len(manual) - fold.idx - 1
		} else {
			foldLengthIdx = fold.idx
		}
	}

	if fold.jdx > 0 {
		for idx := 0; idx < len(manual); idx++ {
			manual[idx][fold.jdx] = "-"
		}

		if fold.jdx > len(manual[0])-fold.jdx-1 {
			foldLengthJdx = len(manual[0]) - fold.jdx - 1
		} else {
			foldLengthJdx = fold.jdx
		}
	}

	return
}

func PrintManual(manual [][]string) {
	for i := range manual {
		fmt.Println(manual[i])
	}
}

func rewriteDots(manual [][]string, fold Point, foldLengthIdx int, foldLengthJdx int) {
	for idx := 1; idx <= foldLengthIdx; idx++ {
		for jdx := 0; jdx < len(manual[0]); jdx++ {
			if manual[fold.idx+idx][jdx] == "#" {
				manual[fold.idx-idx][jdx] = manual[fold.idx+idx][jdx]
			}
		}
	}

	for jdx := 1; jdx <= foldLengthJdx; jdx++ {
		for idx := 0; idx < len(manual); idx++ {
			if manual[idx][fold.jdx+jdx] == "#" {
				manual[idx][fold.jdx-jdx] = manual[idx][fold.jdx+jdx]
			}
		}
	}
}

func createNewManualAfterFold(manual [][]string, fold Point) (foldedManual [][]string) {
	maxFoldIdx := 0
	maxFoldJdx := 0
	if fold.idx > 0 {
		maxFoldIdx = fold.idx
	} else {
		maxFoldIdx = len(manual)
	}

	if fold.jdx > 0 {
		maxFoldJdx = fold.jdx
	} else {
		maxFoldJdx = len(manual[0])
	}

	foldedManual = make([][]string, maxFoldIdx)
	for idx := 0; idx < maxFoldIdx; idx++ {
		foldedManual[idx] = make([]string, maxFoldJdx)
		for jdx := 0; jdx < maxFoldJdx; jdx++ {
			foldedManual[idx][jdx] = manual[idx][jdx]
		}
	}

	return
}

func countDots(manual [][]string) (result int) {
	for _, row := range manual {
		for _, val := range row {
			if val == "#" {
				result += 1
			}
		}
	}

	return
}

func getManualAndFolds(data []string) (manual [][]string, folds []Point) {
	idxMax := 0
	jdxMax := 0
	breakIdx := 0
	dots := make([]Point, 0)
	for i, val := range data {
		if val == "" {
			breakIdx = i
			break
		}

		idx, err := strconv.Atoi(strings.Split(val, ",")[1])
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		if idxMax < idx {
			idxMax = idx
		}

		jdx, err := strconv.Atoi(strings.Split(val, ",")[0])
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}

		if jdxMax < jdx {
			jdxMax = jdx
		}

		dots = append(dots, Point{idx: idx, jdx: jdx})
	}

	manual = make([][]string, idxMax+1)
	for idx := 0; idx < idxMax+1; idx++ {
		manual[idx] = make([]string, jdxMax+1)
		for jdx := 0; jdx < jdxMax+1; jdx++ {
			manual[idx][jdx] = "."
		}
	}

	for _, dot := range dots {
		manual[dot.idx][dot.jdx] = "#"
	}

	folds = make([]Point, 0)
	for idx := breakIdx + 1; idx < len(data); idx++ {
		foldString := strings.Split(data[idx], " ")[2]
		if strings.Split(foldString, "=")[0] == "y" {
			i, err := strconv.Atoi(strings.Split(foldString, "=")[1])
			if err != nil {
				log.Fatal(err)
				os.Exit(2)
			}

			folds = append(folds, Point{idx: i, jdx: 0})
		} else if strings.Split(foldString, "=")[0] == "x" {
			j, err := strconv.Atoi(strings.Split(foldString, "=")[1])
			if err != nil {
				log.Fatal(err)
				os.Exit(2)
			}

			folds = append(folds, Point{idx: 0, jdx: j})
		}
	}

	return
}
