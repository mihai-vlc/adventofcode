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

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int
	var dial = 50

	for _, v := range lines {
		var direction = v[0]
		steps, err := strconv.Atoi(v[1:])
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

	for _, line := range lines {
		result = len(line)
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
