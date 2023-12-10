package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"slices"
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

	var result int

	for _, line := range lines {
		parts := strings.Split(line, "|")
		card := strings.Split(parts[0], ":")

		cardValues := strings.Split(strings.TrimSpace(card[1]), " ")
		cardNumbers := Filter(Map(cardValues, strings.TrimSpace), isNotEmpty)

		winningValues := strings.Split(strings.TrimSpace(parts[1]), " ")
		winningNumbers := Filter(Map(winningValues, strings.TrimSpace), isNotEmpty)

		count := 0
		for _, n := range cardNumbers {
			if slices.Contains(winningNumbers, n) {
				count++
			}
		}

		if count > 0 {
			result += int(math.Pow(2, float64(count)-1))
		}
	}

	log.Println("result =", result)
}

func isNotEmpty(s string) bool {
	return len(s) > 0
}

func Map[T, U any](input []T, f func(T) U) []U {
	output := make([]U, len(input))
	for i := range input {
		output[i] = f(input[i])
	}
	return output
}

func Filter[T any](input []T, f func(T) bool) []T {
	output := []T{}
	for i := range input {
		if f(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func part2() {
	return
	lines, err := readAllLines("./input.test")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var result int

	for _, line := range lines {
		log.Println(line)
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
