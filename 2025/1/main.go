package main

import (
	"log"
	"os"
	"regexp"
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

	var result int
	var dial = 50

	for _, line := range lines {
		var direction = line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		steps = steps % 100

		if direction == 'L' {
			dial -= steps
		} else {
			dial += steps
		}

		if dial < 0 {
			dial = 100 + dial
		} else {
			dial = dial % 100
		}

		if dial == 0 {
			result += 1
		}
	}

	log.Println("result =", result)
}

func part2() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int
	var dial = 50

	for _, line := range lines {
		var direction = line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		for i := 0; i < steps; i++ {
			if direction == 'L' {
				dial -= 1
			} else {
				dial += 1
			}

			if dial < 0 {
				dial = 100 + dial
			}

			if dial == 100 {
				dial = 0
			}

			if dial == 0 {
				result += 1
			}
		}

	}

	log.Println("result =", result)
}

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}
