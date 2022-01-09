package optimalpath

import (
	"math"
)

type Point struct {
	idx int
	jdx int
}

func GetRiskLevelOfShortestPath(data [][]int) int {
	return findOptimalPath(data)
}

func GetRiskLevelOfShortestPathForExtendedMatrix(data [][]int) int {
	data = extendInputNTimes(data, 5)

	return findOptimalPath(data)
}

func extendInputNTimes(matrix [][]int, n int) (newMatrix [][]int) {
	rowLength := len(matrix)
	colLength := len(matrix[0])
	newMatrix = make([][]int, rowLength*n)
	for idx := range newMatrix {
		newMatrix[idx] = make([]int, colLength*n)
	}

	for odx := 0; odx < n; odx++ {
		for idx := range matrix {
			for jdx := range matrix[idx] {
				for kdx := 0; kdx < n; kdx++ {
					increment := matrix[idx][jdx] + kdx + odx
					if increment >= 10 {
						increment = increment - 9
					}

					newMatrix[(odx*rowLength)+idx][(kdx*colLength)+jdx] = increment
				}
			}
		}
	}

	return
}

func findOptimalPath(matrix [][]int) int {
	rowLength := len(matrix)
	colLength := len(matrix[0])
	visited := make([][]int, rowLength)
	distances := make([][]int, rowLength)
	for idx := range visited {
		visited[idx] = make([]int, colLength)
		distances[idx] = make([]int, colLength)
		for jdx := range visited {
			distances[idx][jdx] = math.MaxInt
		}
	}

	startingPoint := Point{idx: 0, jdx: 0}
	visited[startingPoint.idx][startingPoint.jdx] = 1
	distances[0][0] = 0
	distances[0][1] = 0
	distances[1][0] = 0
	for _, neighbour := range getNeighbours(startingPoint, visited, rowLength, colLength) {
		exploreMatrix(matrix, distances, neighbour, visited, rowLength, colLength)
	}

	return distances[len(matrix)-1][len(matrix[0])-1] + matrix[len(matrix)-1][len(matrix[0])-1]
}

func exploreMatrix(matrix [][]int, distances [][]int, currentPoint Point, visited [][]int, rowLength int, colLength int) {
	if currentPoint.idx == rowLength-1 && currentPoint.jdx == colLength-1 {
		return
	}

	visited[currentPoint.idx][currentPoint.jdx] = 1
	neighbours := getNeighbours(currentPoint, visited, rowLength, colLength)

	for _, neighbour := range neighbours {
		distanceThroughCurrent := matrix[currentPoint.idx][currentPoint.jdx] + distances[currentPoint.idx][currentPoint.jdx]
		if distanceThroughCurrent < distances[neighbour.idx][neighbour.jdx] {
			distances[neighbour.idx][neighbour.jdx] = distanceThroughCurrent
		}
	}

	smallest := math.MaxInt
	var smallestPoint Point

	for idx := range visited {
		for jdx := range visited[idx] {
			if visited[idx][jdx] == 1 {
				continue
			}

			if distances[idx][jdx] < smallest {
				smallest = distances[idx][jdx]
				smallestPoint = Point{idx: idx, jdx: jdx}
			}
		}
	}

	if smallest == math.MaxInt {
		return
	}

	exploreMatrix(matrix, distances, smallestPoint, visited, rowLength, colLength)
}

func getNeighbours(currentPoint Point, visited [][]int, rowLength int, colLength int) (neighbours []Point) {
	neighbours = make([]Point, 0)
	if currentPoint.idx-1 > 0 && visited[currentPoint.idx-1][currentPoint.jdx] != 1 {
		neighbours = append(neighbours, Point{idx: currentPoint.idx - 1, jdx: currentPoint.jdx})
	}

	if currentPoint.idx+1 < rowLength && visited[currentPoint.idx+1][currentPoint.jdx] != 1 {
		neighbours = append(neighbours, Point{idx: currentPoint.idx + 1, jdx: currentPoint.jdx})
	}

	if currentPoint.jdx-1 > 0 && visited[currentPoint.idx][currentPoint.jdx-1] != 1 {
		neighbours = append(neighbours, Point{idx: currentPoint.idx, jdx: currentPoint.jdx - 1})
	}

	if currentPoint.jdx+1 < colLength && visited[currentPoint.idx][currentPoint.jdx+1] != 1 {
		neighbours = append(neighbours, Point{idx: currentPoint.idx, jdx: currentPoint.jdx + 1})
	}

	return
}
