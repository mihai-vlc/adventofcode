package main

import (
	"log"
	"os"
	"regexp"
	"sort"
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

type Monkey struct {
	Items              []int
	OperationType      string
	OperationRhs       string
	Divisor            int
	DestinationTrue    int
	DestinationFalse   int
	InspectCount       int
	StressLevelControl func(int) int
}

type InspectionResult struct {
	Value       int
	Destination int
}

var operationRegex = regexp.MustCompile(`new = old (\+|\*) (\d+|old)`)
var testRegex = regexp.MustCompile(`divisible by (\d+)`)
var testResultRegex = regexp.MustCompile(`throw to monkey (\d+)`)

func (m *Monkey) Parse(line string) {
	if strings.Contains(line, "Starting items:") {
		var parts = strings.Split(line, ":")
		var items = strings.Split(parts[1], ",")

		for _, v := range items {
			var n, _ = strconv.Atoi(strings.Trim(v, " "))
			m.Items = append(m.Items, n)
		}
	}

	if strings.Contains(line, "Operation:") {
		var matches = operationRegex.FindStringSubmatch(line)

		m.OperationType = matches[1]
		m.OperationRhs = matches[2]
	}

	if strings.Contains(line, "Test:") {
		var matches = testRegex.FindStringSubmatch(line)
		m.Divisor, _ = strconv.Atoi(matches[1])
	}

	if strings.Contains(line, "If true") {
		var matches = testResultRegex.FindStringSubmatch(line)
		m.DestinationTrue, _ = strconv.Atoi(matches[1])
	}

	if strings.Contains(line, "If false") {
		var matches = testResultRegex.FindStringSubmatch(line)
		m.DestinationFalse, _ = strconv.Atoi(matches[1])
	}
}

func (m *Monkey) InspectItems() []InspectionResult {
	var result = []InspectionResult{}

	for _, v := range m.Items {
		m.InspectCount++

		var dest = m.DestinationFalse
		var val = m.CalculateWorryLevel(v)

		if val%m.Divisor == 0 {
			dest = m.DestinationTrue
		}

		result = append(result, InspectionResult{
			Value:       val,
			Destination: dest,
		})
	}

	m.Items = []int{}

	return result
}

func (m *Monkey) CalculateWorryLevel(input int) int {
	var result = input
	if m.OperationType == "+" {
		if m.OperationRhs == "old" {
			result = result + result
		} else {
			var n, _ = strconv.Atoi(m.OperationRhs)
			result = result + n
		}
	}

	if m.OperationType == "*" {
		if m.OperationRhs == "old" {
			result = result * result
		} else {
			var n, _ = strconv.Atoi(m.OperationRhs)
			result = result * n
		}
	}

	// monkey's inspection didn't damage the item
	result = m.StressLevelControl(result)

	return result
}

type Game struct {
	players []*Monkey
}

func (g *Game) PlayRound() {
	for _, p := range g.players {
		var results = p.InspectItems()

		for _, r := range results {
			var receiver = g.players[r.Destination]

			receiver.Items = append(receiver.Items, r.Value)
		}
	}
}

func part1() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var players = []*Monkey{}

	var currentMonkey *Monkey
	for _, line := range lines {
		if strings.Contains(line, "Monkey") {
			currentMonkey = &Monkey{
				StressLevelControl: func(result int) int {
					return result / 3
				},
			}
			players = append(players, currentMonkey)
			continue
		}

		currentMonkey.Parse(line)
	}

	var game = &Game{
		players: players,
	}

	for i := 0; i < 20; i++ {
		game.PlayRound()
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].InspectCount < players[j].InspectCount
	})

	// log.Printf("%#v \n", game.players[3])

	var size = len(players)
	// most active monkeys
	var result = players[size-1].InspectCount * players[size-2].InspectCount

	log.Println(result)
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var players = []*Monkey{}

	var currentMonkey *Monkey
	for _, line := range lines {
		if strings.Contains(line, "Monkey") {
			currentMonkey = &Monkey{}
			players = append(players, currentMonkey)
			continue
		}

		currentMonkey.Parse(line)
	}

	var mod = 1
	for _, p := range players {
		mod *= p.Divisor
	}

	for _, p := range players {
		p.StressLevelControl = func(result int) int {
			return result % mod
		}
	}

	var game = &Game{
		players: players,
	}

	for i := 0; i < 10000; i++ {
		game.PlayRound()
	}

	// log.Printf("%#v \n", game.players[0])

	sort.Slice(players, func(i, j int) bool {
		return players[i].InspectCount < players[j].InspectCount
	})

	var size = len(players)

	// most active monkeys
	var result = players[size-1].InspectCount * players[size-2].InspectCount

	log.Println(result)
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
