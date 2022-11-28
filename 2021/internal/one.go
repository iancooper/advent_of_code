package advent_of_code

import "container/list"

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

func GroupedDepthIncreases(changes []int, groupSize int) int {
	return findIncreases(sumGroups(groupReadings(changes, groupSize)))
}

func groupReadings(changes []int, groupSize int) []*list.List {
	var groupedDepths = make([]*list.List, len(changes))
	//populate a grouped map
	for index, value := range changes {
		//create a new group
		readings := list.New()
		groupedDepths[index] = readings

		//now iterate over the groups and add to any that are not full yet
		for _, rds := range groupedDepths {
			if rds != nil && rds.Len() < groupSize {
				rds.PushBack(value)
			}
		}
	}
	return groupedDepths
}

func findIncreases(summedGroup []int) int {
	increases := 0
	priorValue := 0
	for _, value := range summedGroup {
		if priorValue > 0 && value > priorValue {
			increases++
		}
		priorValue = value
	}

	return increases
}

func sumGroups(groupedDepths []*list.List) []int {
	var summedGroup = []int{}
	for _, group := range groupedDepths {
		if group != nil && group.Len() == 3 {
			totalDepth := 0
			el := group.Front()
			for el != nil {
				totalDepth += el.Value.(int)
				el = el.Next()
			}
			summedGroup = append(summedGroup, totalDepth)
		}
	}
	return summedGroup
}
