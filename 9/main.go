package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Direction [2]int

var (
	DirectionUp    = Direction{0, 1}
	DirectionDown  = Direction{0, -1}
	DirectionLeft  = Direction{-1, 0}
	DirectionRight = Direction{1, 0}
)

var charToDirection = map[string]Direction{
	"L": DirectionLeft,
	"R": DirectionRight,
	"U": DirectionUp,
	"D": DirectionDown,
}

type Position struct {
	X int
	Y int
}

func (p *Position) Move(d Direction) {
	p.X += d[0]
	p.Y += d[1]
}

func (p *Position) Distance(other *Position) int {
	return max(abs(p.X-other.X), abs(p.Y-other.Y))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {

	if x < 0 {
		return -x
	}
	return x
}

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

	var head = Position{
		X: 0,
		Y: 0,
	}
	var tail = head
	var visited = map[Position]bool{}

	visited[tail] = true

	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var dir = charToDirection[parts[0]]
		var count, _ = strconv.Atoi(parts[1])

		for i := 0; i < count; i++ {
			var prevHead = head
			head.Move(dir)

			if head.Distance(&tail) > 1 {
				tail = prevHead
				visited[tail] = true
			}
		}
	}

	log.Println(len(visited))
}

func part2() {

	lines, err := readAllLines("./input.test")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var head = Position{
		X: 0,
		Y: 0,
	}
	var visited = map[Position]bool{}
	var history = []Position{}

	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var dir = charToDirection[parts[0]]
		var count, _ = strconv.Atoi(parts[1])

		for i := 0; i < count; i++ {
			history = append(history, head)
			head.Move(dir)

			// if len(history) > 8 {
			// 	var tail = prevHead
			// 	visited[tail] = true
			// }
		}
	}

	log.Println(history)
	log.Println(len(visited))
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
