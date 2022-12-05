package advent_of_code_2022

import "testing"

func TestGreatestCalorieCounter(t *testing.T) {

	packs := []string{"1000", "2000", "3000", "", "4000", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	greatestCalories := GreatestCaloriesCounter(packs)
	if greatestCalories != 24000 {
		t.Errorf("Failed. Calories was %d and expected 24000", greatestCalories)
	}
}

func TestHighestThreeCalorieCounter(t *testing.T) {
	packs := []string{"1000", "2000", "3000", "", "4000", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	highestCalories := HighestCaloriesCounter(packs, 3)
	if highestCalories != 45000 {
		t.Errorf("Failed. Calories was %d and expected 24000", highestCalories)
	}
}
