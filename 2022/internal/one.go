package advent_of_code_2022

import "strconv"

func GreatestCaloriesCounter(packs []string) int {

	highestCalories := 0
	packCalories := 0
	for _, ration := range packs {
		if ration == "" {
			packCalories = 0
			continue
		}
		if r, err := strconv.Atoi(ration); err == nil {
			packCalories += r
			if packCalories > highestCalories {
				highestCalories = packCalories
			}
		}
	}

	return highestCalories
}
