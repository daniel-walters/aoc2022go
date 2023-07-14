package day02

import (
	"aoc/utils"
	"bufio"
	"errors"
	"os"
	"strings"
)

func decodeHand(encodedHand string) (string, error) {
	switch encodedHand {
	case "A", "X":
		return "rock", nil
	case "B", "Y":
		return "paper", nil
	case "C", "Z":
		return "scissors", nil
	default:
		return "", errors.New("Invalid hand: " + encodedHand)
	}
}

func getPoints(hand string) (int, error) {
	switch hand {
	case "rock":
		return 1, nil
	case "paper":
		return 2, nil
	case "scissors":
		return 3, nil
	default:
		return 0, errors.New("Invalid hand: " + hand)
	}
}

func getWinner(opponent, player string) string {
	if player == "rock" {
		if opponent == "rock" {
			return "draw"
		} else if opponent == "paper" {
			return "lose"
		} else {
			return "win"
		}
	} else if player == "paper" {
		if opponent == "rock" {
			return "win"
		} else if opponent == "paper" {
			return "draw"
		} else {
			return "lose"
		}
	} else {
		if opponent == "rock" {
			return "lose"
		} else if opponent == "paper" {
			return "win"
		} else {
			return "draw"
		}
	}
}

func calculateScore(opponent, player string) int {
	handScore, err := getPoints(player)
	utils.PanicIfError(err)
	outcome := getWinner(opponent, player)
	if outcome == "win" {
		handScore += 6
	} else if outcome == "draw" {
		handScore += 3
	}

	return handScore
}

func Main(input string) (int, int) {
	slnOne := PartOne(input)
	slnTwo := PartTwo(input)

	return slnOne, slnTwo
}

func PartOne(inputFile string) int {
	file, err := os.Open(inputFile)
	defer utils.HandleFileClose(file)
	utils.PanicIfError(err)

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, " ", 2)

		opponent, err := decodeHand(splitLine[0])
		utils.PanicIfError(err)
		player, err := decodeHand(splitLine[1])
		utils.PanicIfError(err)

		score := calculateScore(opponent, player)
		totalScore += score
	}

	return totalScore
}

func getWinningHand(opponent, result string) (string, error) {
	switch result {
	case "X":
		if opponent == "rock" {
			return "scissors", nil
		} else if opponent == "paper" {
			return "rock", nil
		} else {
			return "paper", nil
		}
	case "Y":
		if opponent == "rock" {
			return "rock", nil
		} else if opponent == "paper" {
			return "paper", nil
		} else {
			return "scissors", nil
		}
	case "Z":
		if opponent == "rock" {
			return "paper", nil
		} else if opponent == "paper" {
			return "scissors", nil
		} else {
			return "rock", nil
		}
	default:
		return "", errors.New("Invalid result: " + result)
	}
}

func PartTwo(inputFile string) int {
	file, err := os.Open(inputFile)
	defer utils.HandleFileClose(file)
	utils.PanicIfError(err)

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, " ", 2)

		opponent, err := decodeHand(splitLine[0])
		utils.PanicIfError(err)
		player, err := getWinningHand(opponent, splitLine[1])
		utils.PanicIfError(err)

		score := calculateScore(opponent, player)
		totalScore += score
	}

	return totalScore
}
