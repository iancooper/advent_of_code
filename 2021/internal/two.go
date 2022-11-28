package advent_of_code

import (
	"errors"
	"strconv"
	"strings"
)

type Direction int

const (
	Forward Direction = iota
	Down
	Up
)

func DirectionFromString(s string) (Direction, error) {

	switch s {
	case "forward":
		return Forward, nil
	case "down":
		return Down, nil
	case "up":
		return Up, nil
	}

	return 0, errors.New("unknown direction")
}

type PositionChange struct {
	command   Direction
	magnitude int
}

func MagnitudeFromString(s string) (int, error) {
	i, e := strconv.Atoi(s)
	return i, e
}

func ReadPositionChangesFrom(rawChanges []string) []PositionChange {
	positionChanges := make([]PositionChange, 0, 1000)
	for _, s := range rawChanges {
		v := strings.Fields(s)
		cmd, _ := DirectionFromString(v[0])
		mag, _ := MagnitudeFromString(v[1])
		posChange := PositionChange{
			command:   cmd,
			magnitude: mag,
		}
		positionChanges = append(positionChanges, posChange)
	}
	return positionChanges
}

type Position struct {
	Horizontal int
	Vertical   int
}

func ApplyPositionChanges(positionChanges []PositionChange) Position {

	pos := Position{
		Horizontal: 0,
		Vertical:   0,
	}

	for _, v := range positionChanges {
		switch v.command {
		case Forward:
			pos.Horizontal += v.magnitude
		case Down:
			pos.Vertical += v.magnitude
		case Up:
			pos.Vertical -= v.magnitude
		}
	}

	return pos

}

type AimedPosition struct {
	Pos Position
	Aim int
}

func ApplyAimedPositionChanges(positionChanges []PositionChange) Position {

	var apos AimedPosition
	apos.Pos.Vertical = 0
	apos.Pos.Horizontal = 0
	apos.Aim = 0

	for _, v := range positionChanges {
		switch v.command {
		case Forward:
			apos.Pos.Horizontal += v.magnitude
			apos.Pos.Vertical += apos.Aim * v.magnitude
		case Down:
			apos.Aim += v.magnitude
		case Up:
			apos.Aim -= v.magnitude
		}
	}

	pos := Position{
		Horizontal: apos.Pos.Horizontal,
		Vertical:   apos.Pos.Vertical,
	}

	return pos

}
