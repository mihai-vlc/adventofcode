package main

import (
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"unicode"
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
	var engineMatrix = [][]string{}

	for i, line := range lines {
		engineMatrix = append(engineMatrix, []string{})

		var currentNumber string

		for _, c := range line {
			if unicode.IsDigit(c) {
				currentNumber += string(c)
			} else {
				if currentNumber != "" {
					for j := 0; j < len(currentNumber); j++ {
						engineMatrix[i] = append(engineMatrix[i], currentNumber)
					}
				}
				engineMatrix[i] = append(engineMatrix[i], string(c))
				currentNumber = ""
			}
		}
		if currentNumber != "" {
			for j := 0; j < len(currentNumber); j++ {
				engineMatrix[i] = append(engineMatrix[i], currentNumber)
			}
		}
	}

	var lastValue = ""
	for i := range engineMatrix {
		for j := range engineMatrix[i] {
			if !isInt(engineMatrix[i][j]) {
				lastValue = ""
				continue
			}
			if lastValue == engineMatrix[i][j] {
				continue
			}

			if hasAdjacentSymbol(engineMatrix, i, j) {
				lastValue = engineMatrix[i][j]
				n, _ := strconv.Atoi(engineMatrix[i][j])
				result += n
			}
		}
		lastValue = ""
	}

	log.Println("result =", result)
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func hasAdjacentSymbol(matrix [][]string, i int, j int) bool {
	var adjacentPositions = [][]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j + 1},
		{i + 1, j + 1},
		{i + 1, j},
		{i + 1, j - 1},
		{i, j - 1},
	}
	var rowCount = len(matrix)
	var colCount = len(matrix[0])

	for _, pos := range adjacentPositions {
		x, y := pos[0], pos[1]
		if x < 0 || x >= rowCount {
			continue
		}
		if y < 0 || y >= colCount {
			continue
		}

		if matrix[x][y] != "." && !isInt(matrix[x][y]) {
			return true
		}
	}

	return false
}

func part2() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var engineMatrix = [][]string{}

	for i, line := range lines {
		engineMatrix = append(engineMatrix, []string{})

		var currentNumber string

		for _, c := range line {
			if unicode.IsDigit(c) {
				currentNumber += string(c)
			} else {
				if currentNumber != "" {
					for j := 0; j < len(currentNumber); j++ {
						engineMatrix[i] = append(engineMatrix[i], currentNumber)
					}
				}
				engineMatrix[i] = append(engineMatrix[i], string(c))
				currentNumber = ""
			}
		}
		if currentNumber != "" {
			for j := 0; j < len(currentNumber); j++ {
				engineMatrix[i] = append(engineMatrix[i], currentNumber)
			}
		}
	}

	var result int
	for i := range engineMatrix {
		for j := range engineMatrix[i] {
			if engineMatrix[i][j] != "*" {
				continue
			}

			nums := getAdjacentNumbers(engineMatrix, i, j)
			if len(nums) == 2 {
				result += nums[0] * nums[1]
			}
		}
	}

	log.Println("result =", result)
}

func getAdjacentNumbers(matrix [][]string, i int, j int) []int {
	var adjacentPositions = [][]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j + 1},
		{i + 1, j + 1},
		{i + 1, j},
		{i + 1, j - 1},
		{i, j - 1},
	}
	var rowCount = len(matrix)
	var colCount = len(matrix[0])

	nums := []int{}

	for _, pos := range adjacentPositions {
		x, y := pos[0], pos[1]
		if x < 0 || x >= rowCount {
			continue
		}
		if y < 0 || y >= colCount {
			continue
		}
		if isInt(matrix[x][y]) {
			n, _ := strconv.Atoi(matrix[x][y])
			if !slices.Contains(nums, n) {
				nums = append(nums, n)
			}
		}
	}

	return nums
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
