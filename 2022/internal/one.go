package advent_of_code_2022

import (
	"sort"
	"strconv"
)

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

func HighestCaloriesCounter(packs []string, i int) int {

	var totalHighestCalories = 0
	highestCalories := make([]int, 0, 10)
	packCalories := 0
	for _, ration := range packs {
		if ration == "" {
			highestCalories = append(highestCalories, packCalories)
			packCalories = 0
			continue
		}
		if r, err := strconv.Atoi(ration); err == nil {
			packCalories += r
		}
	}

	sort.Ints(highestCalories)
	reverse(highestCalories)
	totalHighestCalories = sum(highestCalories[:i])

	return totalHighestCalories
}

func reverse(input []int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
