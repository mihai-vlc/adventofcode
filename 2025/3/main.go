package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"slices"
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

	for _, line := range lines {
		var digits = textToDigits(line)
		var firstDigitValue = slices.Max(digits)
		var firstDigitIndex = slices.Index(digits, firstDigitValue)
		var secondDigitValue int

		if firstDigitIndex+1 == len(line) {
			secondDigitValue = firstDigitValue
			firstDigitValue = slices.Max(digits[0:firstDigitIndex])
		} else {
			secondDigitValue = slices.Max(digits[firstDigitIndex+1:])
		}

		result += firstDigitValue*10 + secondDigitValue
	}

	log.Println("result =", result)
}

func textToDigits(input string) []int {
	var digits = strings.Split(input, "")
	var result = []int{}

	for _, item := range digits {
		n, _ := strconv.Atoi(item)
		result = append(result, n)
	}

	return result
}

func part2() {
	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int

	for _, line := range lines {
		var digits = textToDigits(line)
		var n int
		var start = 0
		var end = len(digits) - 11
		for i := range 12 {
			value, index := findMax(digits[start:end])
			n += value * int(math.Pow(10.0, float64(12-i-1)))

			start += index + 1
			end = len(digits) - 11 + i + 1
		}

		result += n
	}

	log.Println("result =", result)
}

func findMax(items []int) (int, int) {
	var result = slices.Max(items)
	var index = slices.Index(items, result)
	return result, index
}

func main() {
	log.SetFlags(0)

	log.Println("== Part 1")
	part1()

	log.Println("== Part 2")
	part2()
}
