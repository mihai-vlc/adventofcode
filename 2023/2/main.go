package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}

func part1() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int

	for _, line := range lines {
		game := Game{}
		err := game.Load(line)
		if err != nil {
			continue
		}

		if game.IsValid() {
			result += game.ID
		}
	}

	log.Println("result =", result)
}

func part2() {
	_, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int

	log.Println("result =", result)
}

type Selection struct {
	Count     int
	ColorName string
}

func (s *Selection) Load(selectionText string) {
	parts := strings.Split(selectionText, " ")
	s.Count, _ = strconv.Atoi(parts[0])
	s.ColorName = parts[1]
}

var limitsMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func (s *Selection) IsValid() bool {
	limit := limitsMap[s.ColorName]

	return s.Count <= limit
}

type Round struct {
	Selections []Selection
}

func (r *Round) Load(roundText string) {
	for _, selectionText := range strings.Split(roundText, ",") {
		s := Selection{}
		s.Load(strings.TrimSpace(selectionText))
		r.Selections = append(r.Selections, s)
	}
}

func (r *Round) IsValid() bool {
	for _, selection := range r.Selections {
		if !selection.IsValid() {
			return false
		}
	}
	return true
}

type Game struct {
	ID     int
	Rounds []Round
}

func (g *Game) Load(line string) error {
	parts := strings.Split(line, ":")
	gameName := strings.TrimSpace(parts[0])
	rounds := strings.TrimSpace(parts[1])

	id, err := strconv.Atoi(strings.Split(gameName, " ")[1])
	if err != nil {
		return err
	}
	g.ID = id

	for _, roundText := range strings.Split(rounds, ";") {
		currentRound := Round{}
		currentRound.Load(strings.TrimSpace(roundText))
		g.Rounds = append(g.Rounds, currentRound)
	}

	return nil
}

func (g *Game) IsValid() bool {
	for _, round := range g.Rounds {
		if !round.IsValid() {
			return false
		}
	}
	return true
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
