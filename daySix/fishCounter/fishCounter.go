package fishcounter

type Lanternfish struct {
	incubationDays int
}

func NewLanternfish(d int) *Lanternfish {
	return &Lanternfish{incubationDays: d}
}

func GetLanternfishNumberAfterDays(data []int, n int) int {
	lanternfishes := make([]Lanternfish, 0)

	for _, fish := range data {
		lanternfishes = append(lanternfishes, *NewLanternfish(fish))
	}

	for n > 0 {
		lanternfishesToAdd := make([]Lanternfish, 0)
		for i := range lanternfishes {
			if lanternfishes[i].incubationDays == 0 {
				lanternfishesToAdd = append(lanternfishesToAdd, *NewLanternfish(8))
				lanternfishes[i].incubationDays = 6
			} else {
				lanternfishes[i].incubationDays--
			}
		}

		lanternfishes = append(lanternfishes, lanternfishesToAdd...)
		n--
	}

	return len(lanternfishes)
}

func GetLanternfishNumberAfterDaysOptimal(data []int, n int) (sumOfAllLanternfishes int) {
	dailyLanternFishes := make([]int, 9)

	for i := range data {
		dailyLanternFishes[data[i]]++
	}

	for i := n; i > 0; i-- {
		oldFish := dailyLanternFishes[0]
		dailyLanternFishes[0] = dailyLanternFishes[1]
		dailyLanternFishes[1] = dailyLanternFishes[2]
		dailyLanternFishes[2] = dailyLanternFishes[3]
		dailyLanternFishes[3] = dailyLanternFishes[4]
		dailyLanternFishes[4] = dailyLanternFishes[5]
		dailyLanternFishes[5] = dailyLanternFishes[6]
		dailyLanternFishes[6] = dailyLanternFishes[7] + oldFish
		dailyLanternFishes[7] = dailyLanternFishes[8]
		dailyLanternFishes[8] = oldFish
	}

	for _, val := range dailyLanternFishes {
		sumOfAllLanternfishes += val
	}

	return
}

func CountAllChildrenOfSingleLanternfishInGivenDays(days int, startIncubationDay int, childrenIncubationDuration int, daysToBeAdultLanternfish int) (sumOfChildren int) {
	numberOfBirths := (days - startIncubationDay) / 8
	for i := range make([]int, numberOfBirths+1) {
		newChildrenKids := (days - startIncubationDay - (daysToBeAdultLanternfish+childrenIncubationDuration-1)*i) / childrenIncubationDuration
		if i >= 2 {
			newChildrenKids *= i
		}

		sumOfChildren += newChildrenKids
	}

	println(sumOfChildren)
	println(numberOfBirths)

	sumOfChildren += numberOfBirths

	return
}
