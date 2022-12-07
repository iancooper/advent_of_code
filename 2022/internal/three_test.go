package advent_of_code_2022

import "testing"

func TestDuplicatedItems_In_BackpackA(t *testing.T) {

	compartmentOne, compartmentTwo, duplicates := BackpackDuplicates("vJrwpWtwJgWrhcsFMMfFFhFp")

	if compartmentOne != "vJrwpWtwJgWr" {
		t.Errorf("Expected comparment one to be vJrwpWtwJgWr but was %s", compartmentOne)
	}

	if compartmentTwo != "hcsFMMfFFhFp" {
		t.Errorf("Expected comparment two to be hcsFMMfFFhFp but was %s", compartmentTwo)
	}

	if len(compartmentOne) != len(compartmentTwo) {
		t.Errorf("Expected compartments to be evenly divided")
	}

	if duplicates[0] != 'p' {
		t.Errorf("Expected a duplicate but got %d", duplicates[0])
	}
}

func TestDuplicatedItems_In_BackpackB(t *testing.T) {

	compartmentOne, compartmentTwo, duplicates := BackpackDuplicates("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")

	if compartmentOne != "jqHRNqRjqzjGDLGL" {
		t.Errorf("Expected comparment one to be jqHRNqRjqzjGDLGL but was %s", compartmentOne)
	}

	if compartmentTwo != "rsFMfFZSrLrFZsSL" {
		t.Errorf("Expected comparment two to be rsFMfFZSrLrFZsSL but was %s", compartmentTwo)
	}

	if len(compartmentOne) != len(compartmentTwo) {
		t.Errorf("Expected compartments to be evenly divided")
	}

	if duplicates[0] != 'L' {
		t.Errorf("Expected a duplicate but got %d", duplicates[0])
	}
}

func TestDuplicatedItems_In_BackpackC(t *testing.T) {

	compartmentOne, compartmentTwo, duplicates := BackpackDuplicates("PmmdzqPrVvPwwTWBwg")

	if compartmentOne != "PmmdzqPrV" {
		t.Errorf("Expected comparment one to be PmmdzqPrV but was %s", compartmentOne)
	}

	if compartmentTwo != "vPwwTWBwg" {
		t.Errorf("Expected comparment two to be vPwwTWBwg but was %s", compartmentTwo)
	}

	if len(compartmentOne) != len(compartmentTwo) {
		t.Errorf("Expected compartments to be evenly divided")
	}

	if duplicates[0] != 'P' {
		t.Errorf("Expected a duplicate but got %d", duplicates[0])
	}
}

func TestConvertRuneToPriority(t *testing.T) {
	priorityp := RuneToPriority('p')
	if priorityp != 16 {
		t.Errorf("Expected a value of 16 got %d", priorityp)
	}

	priorityL := RuneToPriority('L')
	if priorityL != 38 {
		t.Errorf("Expected a value of 38 got %d", priorityL)
	}

	priorityP := RuneToPriority('P')
	if priorityP != 42 {
		t.Errorf("Expected a value of 42 got %d", priorityP)
	}

	priorityv := RuneToPriority('v')
	if priorityv != 22 {
		t.Errorf("Expected a value of 22 got %d", priorityv)
	}

	priorityt := RuneToPriority('t')
	if priorityt != 20 {
		t.Errorf("Expected a value of 20 got %d", priorityt)
	}

	prioritys := RuneToPriority('s')
	if prioritys != 19 {
		t.Errorf("Expected a value of 19 got %d", prioritys)
	}

}

func TestDuplicatedItems_In_Backpacks(t *testing.T) {
	backpacks := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}

	priorityTotal := PriorityOfDuplicates(backpacks)

	if priorityTotal != 157 {
		t.Errorf("Expected a total of 157 but got %d", priorityTotal)
	}
}

func TestCommonItems_In_Backpacks_GroupA(t *testing.T) {
	backpacks := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}

	commonItem := CommonItem(backpacks)

	if commonItem != 'r' {
		t.Errorf("Expected r but was %c", commonItem)
	}
}

func TestCommonItems_In_Backpacks_GroupB(t *testing.T) {
	backpacks := []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}

	commonItem := CommonItem(backpacks)

	if commonItem != 'Z' {
		t.Errorf("Expected r but was %c", commonItem)
	}
}

func TestCommonItems_Across_Groups(t *testing.T) {
	backpacks := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}

	priority := PriorityByBadge(backpacks, 3)

	if priority != 70 {
		t.Errorf("Expected a badge priority of 70 but was %d", priority)
	}

}
