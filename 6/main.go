package main

import (
	"log"
	"os"
	"regexp"
)

type buffer struct {
	data     string
	capacity int
}

func (b *buffer) Add(c rune) {
	var size = b.Size()

	if size+1 > b.capacity {
		b.data = b.data[size-b.capacity+1:] + string(c)
		return
	}

	b.data += string(c)
}

func (b *buffer) Clear() {
	b.data = ""
}

func (b *buffer) Size() int {
	return len(b.data)
}

func (b *buffer) AreUnique() bool {
	var m = map[rune]bool{}

	for _, c := range b.data {
		if m[c] {
			return false

		}
		m[c] = true
	}

	return true
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

	var capacity = 4

	var buf = buffer{
		data:     "",
		capacity: capacity,
	}

	for _, line := range lines {
		buf.Clear()

		for i, c := range line {
			buf.Add(c)

			if buf.Size() == capacity && buf.AreUnique() {
				log.Println("result =", i+1)
				break
			}
		}
	}

}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var capacity = 14
	var buf = buffer{
		data:     "",
		capacity: capacity,
	}

	for _, line := range lines {
		buf.Clear()

		for i, c := range line {
			buf.Add(c)

			if buf.Size() == capacity && buf.AreUnique() {
				log.Println("result =", i+1)
				break
			}
		}
	}
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
