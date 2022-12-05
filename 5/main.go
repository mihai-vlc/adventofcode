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

	var stackElement = regexp.MustCompile(`[ A-Z\[\]]{3} ?`)
	var instruction = regexp.MustCompile(`move \d+ from \d+ to \d+`)
	var number = regexp.MustCompile(`\d+`)
	var elements [][]string

	for _, line := range lines {
		if stackElement.MatchString(line) {
			var parts = stackElement.FindAllString(line, -1)

			if elements == nil {
				elements = getElementsContainer(len(parts))
			}

			for i, el := range parts {
				el = strings.Trim(el, " []")
				if el == "" {
					continue
				}
				elements[i] = append(elements[i], el)
			}
		}

		if instruction.MatchString(line) {
			var parts = number.FindAllString(line, -1)
			var count, _ = strconv.Atoi(parts[0])
			var source, _ = strconv.Atoi(parts[1])
			var destination, _ = strconv.Atoi(parts[2])

			source--
			destination--

			for i := 0; i < count; i++ {
				var item = elements[source][0]
				elements[destination] = append([]string{item}, elements[destination]...)
				elements[source] = elements[source][1:]
			}
		}

	}

	var result = ""
	for _, v := range elements {
		result += v[0]
	}

	log.Println("result =", result)
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var stackElement = regexp.MustCompile(`[ A-Z\[\]]{3} ?`)
	var instruction = regexp.MustCompile(`move \d+ from \d+ to \d+`)
	var number = regexp.MustCompile(`\d+`)
	var elements [][]string

	for _, line := range lines {
		if stackElement.MatchString(line) {
			var parts = stackElement.FindAllString(line, -1)

			if elements == nil {
				elements = getElementsContainer(len(parts))
			}

			for i, el := range parts {
				el = strings.Trim(el, " []")
				if el == "" {
					continue
				}
				elements[i] = append(elements[i], el)
			}
		}

		if instruction.MatchString(line) {
			var parts = number.FindAllString(line, -1)
			var count, _ = strconv.Atoi(parts[0])
			var source, _ = strconv.Atoi(parts[1])
			var destination, _ = strconv.Atoi(parts[2])

			source--
			destination--

			var items = elements[source][0:count]
			var aa = append([]string{}, items...)
			elements[destination] = append(aa, elements[destination]...)
			elements[source] = elements[source][count:]
		}

	}

	var result = ""
	for _, v := range elements {
		result += v[0]
	}

	log.Println("result =", result)
}

func getElementsContainer(size int) [][]string {
	var result = make([][]string, size)

	for i := range result {
		result[i] = []string{}
	}

	return result
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
