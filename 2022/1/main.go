package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

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

func part1() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	maxCalories := 0
	elfIndex := 0

	currentCalories := 0
	currentElf := 0
	for _, line := range lines {
		if line == "" {
			currentElf++

			if currentCalories > maxCalories {
				maxCalories = currentCalories
				elfIndex = currentElf
			}

			currentCalories = 0
			continue
		}

		n, err := strconv.Atoi(line)

		if err != nil {
			log.Println(err)
			continue
		}
		currentCalories += n
	}

	log.Println("elfNumber =", elfIndex)
	log.Println("maxCalories =", maxCalories)
}

func part2() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	elfsCalories := []int{}

	currentCalories := 0
	currentElf := 0
	for _, line := range lines {
		if line == "" {
			currentElf++

			elfsCalories = append(elfsCalories, currentCalories)

			currentCalories = 0
			continue
		}

		n, err := strconv.Atoi(line)

		if err != nil {
			log.Println(err)
			continue
		}
		currentCalories += n
	}

	sort.Ints(elfsCalories)
	size := len(elfsCalories)
	result := 0

	for i := max(size-3, 0); i < size; i++ {
		result += elfsCalories[i]
	}

	log.Println("top 3 calories =", result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}
