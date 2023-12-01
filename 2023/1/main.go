package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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

	for _, v := range lines {
		var digits = make(map[string]string)

		for _, c := range v {
			if !unicode.IsDigit(c) {
				continue
			}
			_, hasFirst := digits["first"]
			if !hasFirst {
				digits["first"] = string(c)
			}

			digits["last"] = string(c)
		}
		first, hasFirst := digits["first"]
		last := digits["last"]
		if !hasFirst {
			continue
		}

		n, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalln("ERROR: failed to convert to number", err)
		}
		result += n
	}

	log.Println("result =", result)
}

func part2() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var nameToDigit = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var result int

	for _, line := range lines {

		var digits = make(map[string]string)

		for index, c := range line {
			var digit string
			if unicode.IsDigit(c) {
				digit = string(c)
			} else {
				for key := range nameToDigit {
					if strings.HasPrefix(line[index:], key) {
						digit = nameToDigit[key]
					}
				}
			}

			if digit == "" {
				continue
			}

			_, hasFirst := digits["first"]
			if !hasFirst {
				digits["first"] = digit
			}

			digits["last"] = digit
		}

		n, err := strconv.Atoi(digits["first"] + digits["last"])
		if err != nil {
			log.Fatalln("failed to convert to number", err)
		}

		result += n
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
