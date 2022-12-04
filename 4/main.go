package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

func NewInterval(start string, end string) interval {
	var parsedStart, _ = strconv.Atoi(start)
	var parsedEnd, _ = strconv.Atoi(end)

	return interval{
		start: parsedStart,
		end:   parsedEnd,
	}
}

func (a interval) Contains(b interval) bool {
	return (a.start <= b.start) && (a.end >= b.end)
}

func (a interval) Overlaps(b interval) bool {
	if a.start <= b.start && a.end >= b.end {
		return true
	}

	if a.start <= b.start && a.end >= b.start {
		return true
	}

	if a.start <= b.end && a.end >= b.end {
		return true
	}

	if a.start >= b.start && a.end <= b.end {
		return true
	}

	return false
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

	var result = 0
	var slots = [2]interval{}

	for _, line := range lines {
		var pairs = strings.Split(line, ",")

		for i, p := range pairs {
			var intervalParts = strings.Split(p, "-")
			slots[i] = NewInterval(intervalParts[0], intervalParts[1])
		}

		if slots[0].Contains(slots[1]) || slots[1].Contains(slots[0]) {
			result++
		}
	}

	log.Println("result =", result)
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result = 0
	var slots = [2]interval{}

	for _, line := range lines {
		var pairs = strings.Split(line, ",")

		for i, p := range pairs {
			var intervalParts = strings.Split(p, "-")
			slots[i] = NewInterval(intervalParts[0], intervalParts[1])
		}

		if slots[0].Overlaps(slots[1]) {
			result++
		}
	}

	log.Println("result =", result)
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
