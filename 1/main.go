package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {

	f, err := os.Open("./input.txt")

	if err != nil {
		log.Fatalln("error opening the input file", err)
	}

	data, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalln("error reading the data", err)
	}

	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")

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

	f, err := os.Open("./input.txt")

	if err != nil {
		log.Fatalln("error opening the input file", err)
	}

	data, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatalln("error reading the data", err)
	}

	fileContent := string(data)
	lines := strings.Split(fileContent, "\n")

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

	for i := size - 3; i < size; i++ {
		result += elfsCalories[i]
	}

	log.Println("top 3 calories =", result)
}

func main() {
	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}
