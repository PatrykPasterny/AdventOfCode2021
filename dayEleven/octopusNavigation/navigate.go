package octopusnavigation

import (
	findlowestpoint "github.com/PatrykPasterny/AdventOfCode2021/dayNine/findLowestPoint"
)

type Point struct {
	idx int
	jdx int
}

func CountOctopusFlashesAfterNSteps(data [][]int, n int, rowLength int, colLength int) (result int) {
	for n > 0 {
		flashedThisStep := findlowestpoint.CreateVisitedMatrix(rowLength, colLength)
		for i := range data {
			for j := range data[i] {
				result += riseNeighbourEnergy(data, flashedThisStep, rowLength, colLength, i, j)
			}
		}

		n--
	}

	return result
}

func GetStepWhenAllOctopusesFlash(data [][]int, rowLength int, colLength int) (step int) {
	previousResult := 0
	result := 0
	for result-previousResult != rowLength*colLength {
		previousResult = result
		flashedThisStep := findlowestpoint.CreateVisitedMatrix(rowLength, colLength)
		for i := range data {
			for j := range data[i] {
				result += riseNeighbourEnergy(data, flashedThisStep, rowLength, colLength, i, j)
			}
		}

		step++
	}

	return
}

func riseNeighbourEnergy(data [][]int, visited [][]int, rowLength int, colLength int, i int, j int) (result int) {
	if data[i][j] == 9 {
		data[i][j] = 0
		visited[i][j] = 1
		for _, neighbour := range getUnvisitedNeighbours(data, visited, rowLength, colLength, i, j) {
			result += riseNeighbourEnergy(data, visited, rowLength, colLength, neighbour.idx, neighbour.jdx)
		}

		result++
	} else {
		if visited[i][j] == 0 {
			data[i][j]++
		}
	}

	return
}

func getUnvisitedNeighbours(data [][]int, visited [][]int, rowLength int, colLength int, i int, j int) (result []Point) {
	result = make([]Point, 0)
	if i-1 >= 0 {
		if visited[i-1][j] == 0 {
			result = append(result, Point{idx: i - 1, jdx: j})
		}

		if j-1 >= 0 && visited[i-1][j-1] == 0 {
			result = append(result, Point{idx: i - 1, jdx: j - 1})
		}

		if j+1 < colLength && visited[i-1][j+1] == 0 {
			result = append(result, Point{idx: i - 1, jdx: j + 1})
		}
	}

	if i+1 < rowLength {
		if visited[i+1][j] == 0 {
			result = append(result, Point{idx: i + 1, jdx: j})
		}

		if j-1 >= 0 && visited[i+1][j-1] == 0 {
			result = append(result, Point{idx: i + 1, jdx: j - 1})
		}

		if j+1 < colLength && visited[i+1][j+1] == 0 {
			result = append(result, Point{idx: i + 1, jdx: j + 1})
		}
	}

	if j-1 >= 0 && visited[i][j-1] == 0 {
		result = append(result, Point{idx: i, jdx: j - 1})
	}

	if j+1 < colLength && visited[i][j+1] == 0 {
		result = append(result, Point{idx: i, jdx: j + 1})
	}

	return
}
