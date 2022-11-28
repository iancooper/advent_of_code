package advent_of_code

import "testing"

func TestDepthIncrease(t *testing.T) {

	changes := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	count := DepthIncreases(changes)
	if count != 7 {
		t.Errorf("Failed. Count was %d and expected 7", count)
	}
}

func TestGroupedDepthIncrease(t *testing.T) {

	changes := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	count := GroupedDepthIncreases(changes, 3)
	if count != 5 {
		t.Errorf("Failed count was %d and expected 5", count)
	}
}
