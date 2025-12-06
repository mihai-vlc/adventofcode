package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	var productRanges = strings.Split(lines[0], ",")

	for _, item := range productRanges {
		var parts = strings.Split(item, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidId(i) {
				result += i
			}
		}
	}

	log.Println("result =", result)
}

func isInvalidId(id int) bool {
	var value = strconv.Itoa(id)
	var midPoint = len(value) / 2

	return value[0:midPoint] == value[midPoint:]
}

func part2() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int

	log.Println("lines =", lines)
	log.Println("result =", result)
}

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}
