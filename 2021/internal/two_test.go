package advent_of_code

import "testing"

func TestReadPositionChanges(t *testing.T) {
	rawChanges := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expectedPositionChanges := getPositionChanges()

	positionChanges := ReadPositionChangesFrom(rawChanges)

	expectedLen := len(expectedPositionChanges)
	actualLen := len(positionChanges)

	if actualLen != expectedLen {
		t.Errorf("Expected there to be % d elements but there were %d elements", expectedLen, actualLen)
	}

	for i := 0; i < expectedLen; i++ {
		if positionChanges[i] != expectedPositionChanges[i] {
			t.Errorf("Unexpected position changes. Was %v Expected %v", positionChanges, expectedPositionChanges)
		}
	}
}

func TestApplyPositionChanges(t *testing.T) {
	positionChanges := getPositionChanges()
	expectedPosition := Position{
		Horizontal: 15,
		Vertical:   10,
	}

	position := ApplyPositionChanges(positionChanges)

	if expectedPosition != position {
		t.Errorf("Expected position to be %+v but was %+v", expectedPosition, position)
	}

}

func TestApplyAimedPositionChanges(t *testing.T) {
	positionChanges := getPositionChanges()
	expectedPosition := Position{
		Horizontal: 15,
		Vertical:   60,
	}

	position := ApplyAimedPositionChanges(positionChanges)

	if expectedPosition != position {
		t.Errorf("Expected position to be %+v but was %+v", expectedPosition, position)
	}
}

func getPositionChanges() []PositionChange {
	expectedPositionChanges := []PositionChange{
		{
			command:   Forward,
			magnitude: 5,
		},
		{
			command:   Down,
			magnitude: 5,
		},
		{
			command:   Forward,
			magnitude: 8,
		},
		{
			command:   Up,
			magnitude: 3,
		},
		{
			command:   Down,
			magnitude: 8,
		},
		{
			command:   Forward,
			magnitude: 2,
		},
	}
	return expectedPositionChanges
}
