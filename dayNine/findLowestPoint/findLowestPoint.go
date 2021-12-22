package findlowestpoint

type Point struct {
	idx int
	jdx int
}

func GetSumOfRiskLevels(data [][]int, rowLength int, colLength int) int {
	visited := CreateVisitedMatrix(rowLength, colLength)
	lowPoints := findHeightmapLowPoints(data, visited, rowLength, colLength)

	sum := 0
	for _, val := range lowPoints {
		sum += data[val.idx][val.jdx] + 1
	}

	return sum
}

func GetThreeLargestBasins(data [][]int, rowLength int, colLength int) int {
	visited := CreateVisitedMatrix(rowLength, colLength)
	lowPoints := findHeightmapLowPoints(data, visited, rowLength, colLength)
	result := make([]int, 3)

	for _, val := range lowPoints {
		clearVisitedMatrix(visited)

		basinDepth := exploreBasin(val.idx, val.jdx, rowLength, colLength, data, visited)
		reviewDeepestBasins(result, basinDepth)
	}

	return result[0] * result[1] * result[2]
}

func CreateVisitedMatrix(rowLength int, colLength int) (visited [][]int) {
	visited = make([][]int, rowLength)

	for i := 0; i < rowLength; i++ {
		visited[i] = make([]int, colLength)
		for j := 0; j < colLength; j++ {
			visited[i][j] = 0
		}
	}

	return
}

func reviewDeepestBasins(result []int, basinDepth int) {
	if result[2] < basinDepth {
		t1 := result[2]
		t2 := result[1]
		result[2] = basinDepth
		result[1] = t1
		result[0] = t2
	} else if result[1] < basinDepth {
		t1 := result[1]
		result[1] = basinDepth
		result[0] = t1
	} else if result[0] < basinDepth {
		result[0] = basinDepth
	}
}

func clearVisitedMatrix(visited [][]int) {
	for i := range visited {
		for j := range visited[i] {
			visited[i][j] = 0
		}
	}
}

func findHeightmapLowPoints(pointMatrix [][]int, visited [][]int, rowLength int, colLength int) (lowPoints []Point) {
	lowPoints = make([]Point, 0)
	for i := range pointMatrix {
		for j := range pointMatrix[i] {
			if visited[i][j] == 1 {
				continue
			}

			result := exploreMatrix(i, j, rowLength, colLength, pointMatrix, visited)
			if result.idx >= 0 && result.jdx >= 0 {
				lowPoints = append(lowPoints, result)
			}
		}
	}

	return
}

func exploreMatrix(i int, j int, rowLength int, colLength int, matrix [][]int, visited [][]int) Point {
	visited[i][j] = 1
	for _, point := range findNeighbours(i, j, rowLength, colLength, visited) {
		if matrix[point.idx][point.jdx] <= matrix[i][j] {
			if visited[point.idx][point.jdx] == 1 {
				return Point{idx: -1, jdx: -1}
			}

			return exploreMatrix(point.idx, point.jdx, rowLength, colLength, matrix, visited)
		}
	}

	return Point{idx: i, jdx: j}
}

func exploreBasin(i int, j int, rowLength int, colLength int, matrix [][]int, visited [][]int) int {
	visited[i][j] = 1
	sum := 0
	for _, point := range findNeighbours(i, j, rowLength, colLength, visited) {
		if matrix[point.idx][point.jdx] == 9 || visited[point.idx][point.jdx] == 1 {
			continue
		}

		sum += exploreBasin(point.idx, point.jdx, rowLength, colLength, matrix, visited)
	}

	if matrix[i][j] < 9 {
		return sum + 1
	} else {
		return sum
	}
}

func findNeighbours(i int, j int, rowLength int, colLength int, visited [][]int) []Point {
	neighbours := make([]Point, 0)
	if i-1 >= 0 {
		neighbours = append(neighbours, Point{idx: i - 1, jdx: j})
	}

	if j-1 >= 0 {
		neighbours = append(neighbours, Point{idx: i, jdx: j - 1})
	}

	if j+1 < colLength {
		neighbours = append(neighbours, Point{idx: i, jdx: j + 1})
	}

	if i+1 < rowLength {
		neighbours = append(neighbours, Point{idx: i + 1, jdx: j})
	}

	return neighbours
}
