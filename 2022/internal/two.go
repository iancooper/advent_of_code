package advent_of_code_2022

type Move int
type Outcome int

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

const (
	Win  Outcome = 6
	Lose Outcome = 0
	Draw Outcome = 3
)

func Roshambo(opponent string, player string) (Outcome, int) {

	var opponentRps = toRPS(opponent)
	var playerRps = toRPS(player)
	var win = RPS(opponentRps, playerRps)

	return win, int(win) + int(playerRps)
}

func StealthRoshambo(opponent string, player string) (Outcome, int) {

	var opponentRps = toRPS(opponent)
	var outcome = toOutcome(player)

	var move = ForceRPS(opponentRps, outcome)

	return outcome, int(outcome) + int(move)
}

func RoshamboContest(games []string) int {

	total := 0
	for _, game := range games {
		_, score := Roshambo(game[0:1], game[2:3])
		total += score
	}
	return total
}

func StealthRoshamboContest(games []string) int {

	total := 0
	for _, game := range games {
		_, score := StealthRoshambo(game[0:1], game[2:3])
		total += score
	}
	return total
}

func RPS(opponentRps Move, playerRps Move) Outcome {

	if opponentRps == Rock {
		if playerRps == Paper {
			return Win
		} else if playerRps == Scissors {
			return Lose
		} else if playerRps == Rock {
			return Draw
		}
	}

	if opponentRps == Paper {
		if playerRps == Paper {
			return Draw
		} else if playerRps == Scissors {
			return Win
		} else if playerRps == Rock {
			return Lose
		}
	}

	if opponentRps == Scissors {
		if playerRps == Paper {
			return Lose
		} else if playerRps == Scissors {
			return Draw
		} else if playerRps == Rock {
			return Win
		}
	}

	panic("Not a legal contest in RPS")

}

func toRPS(move string) Move {

	switch move {
	case "A":
		fallthrough
	case "X":
		return Rock
	case "B":
		fallthrough
	case "Y":
		return Paper
	case "C":
		fallthrough
	case "Z":
		return Scissors
	}

	panic("Not a known move in RPS")

}

func ForceRPS(opponentRps Move, playerOutcome Outcome) Move {

	if opponentRps == Rock {
		if playerOutcome == Win {
			return Paper
		} else if playerOutcome == Lose {
			return Scissors
		} else if playerOutcome == Draw {
			return Rock
		}
	}

	if opponentRps == Paper {
		if playerOutcome == Win {
			return Scissors
		} else if playerOutcome == Lose {
			return Rock
		} else if playerOutcome == Draw {
			return Paper
		}
	}

	if opponentRps == Scissors {
		if playerOutcome == Win {
			return Rock
		} else if playerOutcome == Lose {
			return Paper
		} else if playerOutcome == Draw {
			return Scissors
		}
	}

	panic("Not a legal contest in RPS")

}

func toOutcome(move string) Outcome {

	switch move {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}

	panic("Not a known move in RPS")

}
