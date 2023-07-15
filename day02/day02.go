package day02

import (
	"aoc/utils"
	"errors"
	"strings"
)

type Hand int
type Result int

const (
	WIN Result = iota
	DRAW
	LOSE
)

const (
	ROCK Hand = iota
	PAPER
	SCISSORS
)

var ErrNotAHand error = errors.New("Encountered a Hand not in ROCK, PAPER, SCISSORS")

var ErrNotAResult error = errors.New("Encountered a Result not in WIN, LOSE, DRAW")

func (h Hand) score() int {
	switch h {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	}

	panic(ErrNotAHand)
}

func (r Result) score() int {
	switch r {
	case WIN:
		return 6
	case DRAW:
		return 3
	case LOSE:
		return 0
	}

	panic(ErrNotAResult)
}

func (h Hand) losesTo() Hand {
	switch h {
	case ROCK:
		return PAPER
	case PAPER:
		return SCISSORS
	case SCISSORS:
		return ROCK
	}

	panic(ErrNotAHand)
}

func (h Hand) beats() Hand {
	switch h {
	case ROCK:
		return SCISSORS
	case PAPER:
		return ROCK
	case SCISSORS:
		return PAPER
	}

	panic(ErrNotAHand)
}

func (h Hand) scoreAgainst(h2 Hand) int {
	var outcome Result

	switch {
	case h.beats() == h2:
		outcome = WIN
	case h.losesTo() == h2:
		outcome = LOSE
	default:
		outcome = DRAW
	}

	return outcome.score() + h.score()
}

func decodeHand(encodedHand string) (Hand, error) {
	switch encodedHand {
	case "A", "X":
		return ROCK, nil
	case "B", "Y":
		return PAPER, nil
	case "C", "Z":
		return SCISSORS, nil
	}

	return 0, errors.New("Invalid hand: " + encodedHand)
}

func getWinner(opponent, player Hand) Result {
	if player.beats() == opponent {
		return WIN
	}

	if player.losesTo() == opponent {
		return LOSE
	}

	return DRAW
}

func calculateScore(opponent, player Hand) int {
	return player.scoreAgainst(opponent)
}

func Main(input string) (int, int) {
	slnOne := PartOne(input)
	slnTwo := PartTwo(input)

	return slnOne, slnTwo
}

func PartOne(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, " ", 2)

		opponent, err := decodeHand(splitLine[0])
		if err != nil {
			panic(err)
		}

		player, err := decodeHand(splitLine[1])
		if err != nil {
			panic(err)
		}

		totalScore += player.scoreAgainst(opponent)
	}

	return totalScore
}

func getWinningHand(opponent Hand, result string) (Hand, error) {
	switch result {
	case "X":
		return opponent.beats(), nil
	case "Y":
		return opponent, nil
	case "Z":
		return opponent.losesTo(), nil
	}

	return -1, errors.New("Invalid result: " + result)
}

func PartTwo(inputFile string) int {
	scanner, closeFile := utils.GetLineScanner(inputFile)
	defer closeFile()

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, " ", 2)

		opponent, err := decodeHand(splitLine[0])
		if err != nil {
			panic(err)
		}

		player, err := getWinningHand(opponent, splitLine[1])
		if err != nil {
			panic(err)
		}

		totalScore += player.scoreAgainst(opponent)
	}

	return totalScore
}
