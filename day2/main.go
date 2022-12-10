package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ROCK     = "rock"
	PAPER    = "paper"
	SCISSORS = "scissors"
	WIN      = "win"
	DRAW     = "draw"
	LOSE     = "lose"
)

type Round struct {
	OpponentPlay  string
	PlayerPlay    string
	WantedOutcome string
}

func check(e error) {
	if e != nil {
		log.Fatalf("Gros naze: %v", e)
	}
}

func decipherPlay(letter string) string {
	var outcome string
	switch letter {
	case "A", "X":
		outcome = ROCK
	case "B", "Y":
		outcome = PAPER
	case "C", "Z":
		outcome = SCISSORS
	default:
		log.Fatalf("Play NOK: " + letter)
	}
	return outcome
}

func decipherResult(letter string) string {
	var outcome string
	switch letter {
	case "A":
		outcome = ROCK
	case "B":
		outcome = PAPER
	case "C":
		outcome = SCISSORS
	case "X":
		outcome = LOSE
	case "Y":
		outcome = DRAW
	case "Z":
		outcome = WIN
	default:
		log.Fatalf("Play NOK: " + letter)
	}
	return outcome
}

func getShapeScore(shape string) int {
	var shapeScore int
	switch shape {
	case ROCK:
		shapeScore = 1
	case PAPER:
		shapeScore = 2
	case SCISSORS:
		shapeScore = 3
	}
	return shapeScore
}

func computePlayScore(round Round) int {
	shapeScore := getShapeScore(round.PlayerPlay)

	var resultScore int
	if (round.OpponentPlay == ROCK && round.PlayerPlay == PAPER) || (round.OpponentPlay == PAPER && round.PlayerPlay == SCISSORS) || (round.OpponentPlay == SCISSORS && round.PlayerPlay == ROCK) {
		resultScore = 6
	} else if round.OpponentPlay == round.PlayerPlay {
		resultScore = 3
	} else {
		resultScore = 0
	}

	return shapeScore + resultScore
}

func computeResultScore(round Round) int {
	var playerShape string
	var resultScore int
	if round.WantedOutcome == WIN {
		switch round.OpponentPlay {
		case ROCK:
			playerShape = PAPER
		case PAPER:
			playerShape = SCISSORS
		case SCISSORS:
			playerShape = ROCK
		}
		resultScore = 6
	} else if round.WantedOutcome == LOSE {
		switch round.OpponentPlay {
		case ROCK:
			playerShape = SCISSORS
		case PAPER:
			playerShape = ROCK
		case SCISSORS:
			playerShape = PAPER
		}
		resultScore = 0
	} else {
		playerShape = round.OpponentPlay
		resultScore = 3
	}

	shapeScore := getShapeScore(playerShape)

	return shapeScore + resultScore
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	playScore := 0
	resultScore := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, " ")
		playRound := Round{
			OpponentPlay: decipherPlay(values[0]),
			PlayerPlay:   decipherPlay(values[1]),
		}
		resultRound := Round{
			OpponentPlay:  decipherResult(values[0]),
			WantedOutcome: decipherResult(values[1]),
		}

		playScore += computePlayScore(playRound)
		resultScore += computeResultScore(resultRound)
	}

	fmt.Println("Total play score: " + strconv.Itoa(playScore))
	fmt.Println("Total result score: " + strconv.Itoa(resultScore))
}
