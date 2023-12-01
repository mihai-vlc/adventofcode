package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
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

	var data = [][]int{}

	for _, line := range lines {
		data = append(data, toIntSlice(line))
	}

	var result = 0

	for i := range data {
		for j := range data[i] {
			if isVisible(i, j, data) {
				result++
			}
		}
	}

	log.Println(result)
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var data = [][]int{}

	for _, line := range lines {
		data = append(data, toIntSlice(line))
	}

	var result = 0

	for i := range data {
		for j := range data[i] {
			var score = scenicScore(i, j, data)
			if score > result {
				result = score
			}
		}
	}

	log.Println(result)
}

func toIntSlice(input string) []int {
	result := make([]int, len(input))
	for i, v := range input {
		var n, _ = strconv.Atoi(string(v))
		result[i] = n
	}
	return result
}

func isVisible(i int, j int, data [][]int) bool {
	var lastColumn = len(data[0]) - 1
	var lastLine = len(data) - 1

	// detect the edge
	if i == 0 || j == 0 || i == lastColumn || j == lastLine {
		return true
	}

	var val = data[i][j]
	var visible bool
	var idx int
	var size int

	// visible up
	visible = true
	idx = i - 1
	for idx >= 0 {
		if data[idx][j] >= val {
			visible = false
		}
		idx--
	}

	if visible {
		return true
	}

	// visible down
	visible = true
	idx = i + 1
	size = len(data)
	for idx < size {
		if data[idx][j] >= val {
			visible = false
		}
		idx++
	}

	if visible {
		return true
	}

	// visible left
	visible = true
	idx = j - 1
	for idx >= 0 {
		if data[i][idx] >= val {
			visible = false
		}
		idx--
	}

	if visible {
		return true
	}

	// visible right
	visible = true
	idx = j + 1
	size = len(data[i])
	for idx < size {
		if data[i][idx] >= val {
			visible = false
		}
		idx++
	}

	if visible {
		return true
	}

	return false
}

func scenicScore(i int, j int, data [][]int) int {
	var lastColumn = len(data[0]) - 1
	var lastLine = len(data) - 1

	// detect the edge
	if i == 0 || j == 0 || i == lastColumn || j == lastLine {
		return 0
	}

	var val = data[i][j]
	var idx int
	var size int

	// visible up
	var scoreUp = 0
	idx = i - 1
	for idx >= 0 {
		scoreUp++
		if data[idx][j] >= val {
			break
		}
		idx--
	}

	// visible down
	var scoreDown = 0
	idx = i + 1
	size = len(data)
	for idx < size {
		scoreDown++
		if data[idx][j] >= val {
			break
		}
		idx++
	}

	// visible left
	var scoreLeft = 0
	idx = j - 1
	for idx >= 0 {
		scoreLeft++
		if data[i][idx] >= val {
			break
		}
		idx--
	}

	// visible right
	var scoreRight = 0
	idx = j + 1
	size = len(data[i])
	for idx < size {
		scoreRight++

		if data[i][idx] >= val {
			break
		}
		idx++
	}

	// log.Println(scoreUp, scoreRight, scoreDown, scoreLeft)

	return scoreUp * scoreRight * scoreDown * scoreLeft
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
