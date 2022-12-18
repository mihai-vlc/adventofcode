package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position2 [2]int

func (p *Position2) Move(d Direction2) {
	p[0] += d[0]
	p[1] += d[1]
}

func (p *Position2) Distance(other *Position2) (int, int) {
	var dx = p[0] - other[0]
	var dy = p[1] - other[1]

	return dx, dy
}

type Direction2 [2]int

type Rope [10]Position2

func (r *Rope) Move(d Direction2) {
	var head = &r[0]

	head.Move(d)

	// update the tail position
	for i := 0; i < len(r)-1; i++ {
		var current = &r[i]
		var next = &r[i+1]
		var dx, dy = current.Distance(next)

		if abs(dx) > 1 || abs(dy) > 1 {
			next[0] += normalize(dx)
			next[1] += normalize(dy)
		}
	}

}

func (r *Rope) Last() Position2 {
	return r[len(r)-1]
}

var keyToDirection = map[string]Direction2{
	"U": {0, 1},
	"D": {0, -1},
	"R": {1, 0},
	"L": {-1, 0},
}

func main() {
	log.SetFlags(0)

	var file, err = os.Open("./input.txt")

	if err != nil {
		log.Fatalln("failed to open the file", err)
	}
	defer file.Close()

	var sc = bufio.NewScanner(file)

	var knots Rope
	var visited = map[Position2]bool{}

	for sc.Scan() {
		var line = sc.Text()
		var parts = strings.Split(line, " ")
		var direction = keyToDirection[parts[0]]
		var stepsCount, _ = strconv.Atoi(parts[1])

		for k := 0; k < stepsCount; k++ {
			knots.Move(direction)

			visited[knots.Last()] = true
		}
	}

	log.Println("== Part 2")
	log.Println(len(visited))
}

func abs(x int) int {

	if x < 0 {
		return -x
	}
	return x
}

func normalize(n int) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	return 1
}
