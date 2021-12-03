package advent_of_code

func DepthIncreases(changes []int) int {
	increases := 0
	priorValue := 0
	for _, value := range changes {
		if priorValue > 0 && value > priorValue {
			increases++
		}
		priorValue = value
	}

	return increases
}
