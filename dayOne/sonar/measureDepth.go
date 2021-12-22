package sonar

func CountDepthIncreases(data []int) int {
	result := 0
	for idx := 1; idx < len(data); idx++ {
		if data[idx] > data[idx-1] {
			result++
		}
	}

	return result
}

func CountDepthIncreasesWithNoise(data []int) int {
	result := 0
	noiseData := [3]int{data[0], data[1], data[2]}
	for idx := 3; idx < len(data); idx++ {
		prevNoiseData := noiseData
		noiseData[idx%3] = data[idx]

		if sumOfNoiseDepthElements(prevNoiseData) < sumOfNoiseDepthElements(noiseData) {
			result++
		}
	}

	return result
}

func sumOfNoiseDepthElements(data [3]int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}

	return sum
}
