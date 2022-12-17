package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos [2]int

const x, y = 0, 1

func main() {
	log.SetFlags(0)

	var file, err = os.Open("./input.txt")

	if err != nil {
		log.Fatalln("failed to open the file", err)
	}
	defer file.Close()

	var sc = bufio.NewScanner(file)

	var knots [10]Pos
	var set = map[Pos]bool{}

	for sc.Scan() {
		var line = sc.Text()
		var parts = strings.Split(line, " ")
		var dir = parts[0]
		var step, _ = strconv.Atoi(parts[1])

		var move Pos
		switch dir {
		case "U":
			move[y] = step
		case "D":
			move[y] = -step
		case "R":
			move[x] = step
		case "L":
			move[x] = -step
		}

		for {
			var d int
			for i := x; i <= y; i++ {
				if move[i] != 0 {
					d = delta(move[i])
					move[i] += -d
					knots[0][i] += d
					break
				}
			}

			if d == 0 {
				break
			}

			for i := 0; i < len(knots)-1; i++ {
				var dx = knots[i][x] - knots[i+1][x]
				var dy = knots[i][y] - knots[i+1][y]

				if abs(dx) > 1 || abs(dy) > 1 {
					knots[i+1][x] += delta(dx)
					knots[i+1][y] += delta(dy)
				}
			}

			set[knots[len(knots)-1]] = true
		}
	}

	log.Println(len(set))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func delta(n int) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	return 1
}
