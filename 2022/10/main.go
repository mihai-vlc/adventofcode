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

	var result = 0
	var cycle = 1
	var registerX = 1

	var trackedCycles = []int{20, 60, 100, 140, 180, 220}
	var trackSignalStrength = func() {
		if contains(trackedCycles, cycle) {
			result += cycle * registerX
		}
	}

	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var instruction = parts[0]

		if instruction == "noop" {
			cycle++
			trackSignalStrength()
			continue
		}
		var param, _ = strconv.Atoi(parts[1])

		if instruction == "addx" {
			cycle++
			trackSignalStrength()

			cycle++
			registerX += param
			trackSignalStrength()
		}
	}

	log.Println(result)
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var cycle = 1
	var registerX = 1

	var crtLine = ""

	var draw = func() {
		var linePos = cycle % 40
		var start = registerX
		var end = registerX + 2

		if linePos >= start && linePos <= end {
			crtLine += "#"
		} else {
			crtLine += "."
		}

		if linePos == 0 {
			log.Println(crtLine)
			crtLine = ""
		}
	}

	draw()

	for _, line := range lines {
		var parts = strings.Split(line, " ")
		var instruction = parts[0]

		if instruction == "noop" {
			cycle++
			draw()
			continue
		}
		var param, _ = strconv.Atoi(parts[1])

		if instruction == "addx" {
			cycle++
			draw()

			cycle++
			registerX += param
			draw()
		}
	}

}

func contains(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}

	return false
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
