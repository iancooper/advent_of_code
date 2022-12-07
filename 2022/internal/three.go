package advent_of_code_2022

import "unicode"

func BackpackDuplicates(backpack string) (string, string, []rune) {

	compartmentOne, compartmentTwo := findCompartments(backpack)

	duplicates := findDuplicates(compartmentOne, compartmentTwo)

	return compartmentOne, compartmentTwo, duplicates
}

func findDuplicates(one string, two string) []rune {
	matches := make(map[rune]bool)
	compartmentOne := []rune(one)
	compartmentTwo := []rune(two)

	for i := 0; i < len(compartmentOne); i++ {
		item := compartmentOne[i]
		if _, exists := matches[item]; !exists {
			for j := 0; j < len(compartmentTwo); j++ {
				if item == compartmentTwo[j] {
					matches[item] = true
					break
				}
			}
		}
	}

	uniqueRunes := make([]rune, 0, 1)

	for item, _ := range matches {
		uniqueRunes = append(uniqueRunes, item)
	}

	return uniqueRunes
}

func PriorityOfDuplicates(backpacks []string) int {

	prioritySum := 0
	groupPrioritySum := 0
	for _, backpack := range backpacks {
		compartmentOne, compartmentTwo := findCompartments(backpack)
		duplicates := findDuplicates(compartmentOne, compartmentTwo)
		groupPrioritySum = 0
		for _, item := range duplicates {
			groupPrioritySum += RuneToPriority(item)
		}
		prioritySum += groupPrioritySum
	}

	return prioritySum
}

func findCompartments(backpack string) (string, string) {
	var split int = len(backpack) / 2

	compartmentOne := backpack[0:split]
	compartmentTwo := backpack[split:len(backpack)]

	return compartmentOne, compartmentTwo
}

func RuneToPriority(item rune) int {

	if unicode.IsUpper(item) {
		return int(item) - 65 + 27
	} else if unicode.IsLower(item) {
		return int(item) - 97 + 1
	}

	panic("Was not a letter")

}

func PriorityByBadge(backpacks []string, groupSize int) int {
	badges := make([]rune, 0, len(backpacks)/3)

	for i := 0; i < len(backpacks)-1; i = i + groupSize {
		badges = append(badges, CommonItem(backpacks[i:i+groupSize]))
	}

	priorityCount := 0
	for _, badge := range badges {
		priorityCount += RuneToPriority(badge)
	}

	return priorityCount
}

func CommonItem(backpacks []string) rune {

	//determine the common runes across the array of strings given by backpacks
	runeCounts := make(map[rune]int, 26)

	for _, backpack := range backpacks {
		//Only add a letter once for a pack
		seen := make(map[rune]bool)
		for _, item := range backpack {
			if _, found := seen[item]; !found {
				runeCounts[item] = runeCounts[item] + 1
				seen[item] = true
			}
		}
	}

	var badge rune

	for key, value := range runeCounts {
		if value == 3 {
			badge = key
			break
		}
	}

	return badge

}
