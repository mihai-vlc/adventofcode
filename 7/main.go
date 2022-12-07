package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type nodeKind int

const (
	File      nodeKind = 0
	Directory nodeKind = 1
)

type node struct {
	Name     string `json:"name"`
	kind     nodeKind
	size     int
	children []*node
	parent   *node
}

func NewNode(name string, kind nodeKind, size int, parent *node) *node {
	return &node{
		Name:   name,
		kind:   kind,
		size:   size,
		parent: parent,
	}
}

func (n *node) AddChild(child *node) {
	if n.children == nil {
		n.children = []*node{}
	}

	n.children = append(n.children, child)
}

func (n *node) FindChild(name string) *node {
	for _, c := range n.children {
		if c.Name == name {
			return c
		}
	}

	return nil
}

func (n *node) Parent() *node {
	return n.parent
}

func (n *node) Kind() nodeKind {
	return n.kind
}

func (n *node) Size() int {
	if n.kind == File {
		return n.size
	}

	if n.children == nil {
		return 0
	}

	var result = 0
	for _, c := range n.children {
		result += c.Size()
	}

	return result
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

	var root = NewNode("/", Directory, 0, nil)
	var currentNode = root
	var numberRegex = regexp.MustCompile(`\d+`)

	for _, line := range lines {
		var parts = strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					currentNode = root
					continue
				}

				if parts[2] == ".." {
					currentNode = currentNode.Parent()
					continue
				}

				currentNode = currentNode.FindChild(parts[2])
				continue
			}
		}

		if parts[0] == "dir" {
			currentNode.AddChild(NewNode(parts[1], Directory, 0, currentNode))
			continue
		}

		if numberRegex.MatchString(parts[0]) {
			var size, _ = strconv.Atoi(parts[0])
			currentNode.AddChild(NewNode(parts[1], File, size, currentNode))
			continue
		}
	}

	var subDirs = getDirectoriesWithSizeLess(root, 100000)
	var result = 0

	for _, d := range subDirs {
		result += d.Size()
	}

	log.Println(result)
}

func getDirectoriesWithSizeLess(n *node, maxSize int) []*node {
	var result []*node = []*node{}

	if n.children == nil {
		return result
	}

	for _, c := range n.children {
		if c.Kind() == Directory {

			if c.Size() <= maxSize {
				result = append(result, c)
			}

			var subDirs = getDirectoriesWithSizeLess(c, maxSize)
			result = append(result, subDirs...)
		}
	}

	return result
}

func part2() {

	lines, err := readAllLines("./input.txt")

	if err != nil {
		log.Fatalln("input reading failed", err)
	}

	var root = NewNode("/", Directory, 0, nil)
	var currentNode = root
	var numberRegex = regexp.MustCompile(`\d+`)

	for _, line := range lines {
		var parts = strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					currentNode = root
					continue
				}

				if parts[2] == ".." {
					currentNode = currentNode.Parent()
					continue
				}

				currentNode = currentNode.FindChild(parts[2])
				continue
			}
		}

		if parts[0] == "dir" {
			currentNode.AddChild(NewNode(parts[1], Directory, 0, currentNode))
			continue
		}

		if numberRegex.MatchString(parts[0]) {
			var size, _ = strconv.Atoi(parts[0])
			currentNode.AddChild(NewNode(parts[1], File, size, currentNode))
			continue
		}
	}

	const totalDiskSpace = 70000000
	var availableDiskSpace = totalDiskSpace - root.Size()
	var neededDiskSpace = 30000000 - availableDiskSpace

	log.Println("availableDiskSpace", availableDiskSpace)
	log.Println("neededDiskSpace", neededDiskSpace)

	var subDirs = getDirectoriesWithSizeMore(root, neededDiskSpace)
	log.Println("result =", MinSize(subDirs...))
}

func getDirectoriesWithSizeMore(n *node, minSize int) []*node {
	var result []*node = []*node{}

	if n.children == nil {
		return result
	}

	for _, c := range n.children {
		if c.Kind() == Directory {

			if c.Size() >= minSize {
				result = append(result, c)
			}

			var subDirs = getDirectoriesWithSizeMore(c, minSize)
			result = append(result, subDirs...)
		}
	}

	return result
}

func MinSize(vars ...*node) int {
	min := vars[0].Size()

	for _, i := range vars {
		if min > i.Size() {
			min = i.Size()
		}
	}

	return min
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
