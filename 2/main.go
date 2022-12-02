package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type Shape int
type Outcome int

const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

const (
	OutcomeLoss Outcome = 0
	OutcomeDraw Outcome = 3
	OutcomeWon  Outcome = 6
)

var letterToScore = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}

func part1() {
	var gameTable = map[Shape]map[Shape]Outcome{
		Rock: {
			Rock:     OutcomeDraw,
			Paper:    OutcomeLoss,
			Scissors: OutcomeWon,
		},
		Paper: {
			Rock:     OutcomeWon,
			Paper:    OutcomeDraw,
			Scissors: OutcomeLoss,
		},
		Scissors: {
			Rock:     OutcomeLoss,
			Paper:    OutcomeWon,
			Scissors: OutcomeDraw,
		},
	}

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	result := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		opponent := letterToScore[parts[0]]
		me := letterToScore[parts[1]]

		result += int(me) + int(gameTable[me][opponent])
	}

	log.Println("result =", result)
}

func part2() {

	var letterToOutcome = map[string]Outcome{
		"X": OutcomeLoss,
		"Y": OutcomeDraw,
		"Z": OutcomeWon,
	}

	var gameTable = map[Shape]map[Outcome]Shape{
		Rock: {
			OutcomeDraw: Rock,
			OutcomeLoss: Scissors,
			OutcomeWon:  Paper,
		},
		Paper: {
			OutcomeDraw: Paper,
			OutcomeLoss: Rock,
			OutcomeWon:  Scissors,
		},
		Scissors: {
			OutcomeDraw: Scissors,
			OutcomeLoss: Paper,
			OutcomeWon:  Rock,
		},
	}

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	result := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		opponent := letterToScore[parts[0]]
		neededOutcome := letterToOutcome[parts[1]]
		me := gameTable[opponent][neededOutcome]

		result += int(me) + int(neededOutcome)
	}

	log.Println("result =", result)
}

func readAllLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile("\\r?\\n")

	fileContent := string(data)
	lines := reg.Split(fileContent, -1)

	return lines, nil
}
