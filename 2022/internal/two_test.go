package advent_of_code_2022

import "testing"

func TestRockPaperScissors_GameOne(t *testing.T) {

	var win, score = Roshambo("A", "Y")
	if win != Win && score != 8 {
		t.Errorf("Expected a game of A/Y to end in a win with 8 points")
	}

}

func TestRockPaperScissors_GameTwo(t *testing.T) {

	var win, score = Roshambo("B", "X")
	if win != Lose && score != 1 {
		t.Errorf("Expected a game of B/X to end in a loss with 1 point")
	}

}

func TestRockPaperScissors_GameThree(t *testing.T) {

	var win, score = Roshambo("C", "Z")
	if win != Draw && score != 6 {
		t.Errorf("Expected a game of C/Z to end in a draw with 6 points")
	}

}

func TestSetOfGames(t *testing.T) {
	games := []string{"A Y", "B X", "C Z"}

	result := RoshamboContest(games)

	if result != 15 {
		t.Errorf("Expected score of 15, but was %d", result)
	}
}

func TestRockPaperScissors_StrategyTwo_GameOne(t *testing.T) {

	var win, score = StealthRoshambo("A", "Y")
	if win != Draw && score != 4 {
		t.Errorf("Expected a game of A/Y to end in a draw with 4 points")
	}
}

func TestRockPaperScissors_StrategyTwo_GameTwo(t *testing.T) {

	var win, score = StealthRoshambo("B", "X")
	if win != Lose && score != 1 {
		t.Errorf("Expected a game of B/X to end in a loss with 1 point")
	}
}

func TestRockPaperScissors_StrategyTwo_GameThree(t *testing.T) {

	var win, score = StealthRoshambo("C", "Z")
	if win != Win && score != 7 {
		t.Errorf("Expected a game of B/X to end in a win with 7 points")
	}

}

func TestSetOfGames_StrategyTwo(t *testing.T) {
	games := []string{"A Y", "B X", "C Z"}

	result := StealthRoshamboContest(games)

	if result != 12 {
		t.Errorf("Expected score of 15, but was %d", result)
	}
}
