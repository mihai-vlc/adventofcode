package main

import (
	"log"
	"os"
	"regexp"
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

	result := 0
	for _, line := range lines {
		size := len(line)
		c1 := line[0 : size/2]
		c2 := line[size/2:]

		for _, r := range c1 {
			if strings.IndexRune(c2, r) > -1 {
				result += getPriority(r)
				break
			}
		}
	}

	log.Println("result =", result)
}

func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	}

	return int(r - 38)
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result = 0
	var groupLines = [3]string{}

	for i, line := range lines {
		var idx = i % 3
		groupLines[idx] = line

		if idx == 2 {
			for _, r := range groupLines[0] {
				var isInSecond = strings.IndexRune(groupLines[1], r) > -1
				var isInThird = strings.IndexRune(groupLines[2], r) > -1
				if isInSecond && isInThird {
					result += getPriority(r)
					break
				}
			}
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
