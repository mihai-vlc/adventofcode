package main

import (
	"log"
	"os"
	"regexp"
)

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}

type Position struct {
	X int
	Y int
}

func part1() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var area [][]int
	var start Position

	for i, line := range lines {
		var vals = []int{}
		for j, v := range line {
			if v == 'S' {
				start.X = i
				start.Y = j
			}

			vals = append(vals, int(v))
		}
		area = append(area, vals)
	}

	var rows = len(area)
	var columns = len(area[0])

	var pos = start
	var val = 96
	// var steps = 0
	var dest = int('E')
	var visited = map[Position]bool{}
	visited[pos] = true

	var isValidStep = func(dx int, dy int) bool {
		var newX = pos.X + dx
		var newY = pos.Y + dy
		var newPos = Position{
			X: newX,
			Y: newY,
		}

		if visited[newPos] {
			return false
		}

		if newX < 0 || newY < 0 {
			return false
		}

		if newX >= rows || newY >= columns {
			return false
		}

		if val == int('z') && area[newX][newY] == int('E') {
			return true
		}

		var cost = area[newX][newY] - val

		if cost != 0 && cost != 1 {
			return false
		}

		return true
	}

	var history = []Position{}
	var steps = 0
	for {
		if val == dest {
			log.Println("found destination")
			break
		}

		// try up
		if isValidStep(-1, 0) {
			pos = Position{pos.X - 1, pos.Y}
			val = area[pos.X][pos.Y]
			visited[pos] = true
			history = append(history, pos)
			steps++
			continue
		}

		// try right
		if isValidStep(0, 1) {
			pos = Position{pos.X, pos.Y + 1}
			val = area[pos.X][pos.Y]
			visited[pos] = true
			history = append(history, pos)
			steps++
			continue
		}

		// try down
		if isValidStep(1, 0) {
			pos = Position{pos.X + 1, pos.Y}
			val = area[pos.X][pos.Y]
			visited[pos] = true
			history = append(history, pos)
			steps++
			continue
		}

		// try left
		if isValidStep(0, -1) {
			pos = Position{pos.X, pos.Y - 1}
			val = area[pos.X][pos.Y]
			visited[pos] = true
			history = append(history, pos)
			steps++
			continue
		}

		if len(history) > 0 {
			pos, history = history[len(history)-1], history[:len(history)-1]
			steps--
			continue
		}

		log.Println("no valid direction, stopping the search", pos)
		break
	}

	log.Println(steps)
}

func part2() {

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
