package day08

import (
	"fmt"
	"log"
	"regexp"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day08(path string) {
	fmt.Println("Day08")

	input, err := utils.ReadInput(path + "/day08")
	if err != nil {
		log.Fatal(err)
	}
	directions, nodes := parseInput(input)
	// part1(directions, nodes)
	part2(directions, nodes)
}

func part1(directions string, nodes map[string]*node) {
	start := nodes["AAA"]
	end := nodes["ZZZ"]
	currentNode := start
	i := 0
	stepCount := 0
	for currentNode != end {
		direction := rune(directions[i])
		switch direction {
		case 'L':
			currentNode = currentNode.left
		case 'R':
			currentNode = currentNode.right
		default:
			panic("invalid direction character")
		}
		stepCount += 1
		i += 1
		if i == len(directions) {
			i = 0
		}
	}
	fmt.Println(stepCount)
}

func part2(directions string, nodes map[string]*node) {
	var currentNodes []*node
	for _, n := range nodes {
		if n.id[2] == 'A' {
			currentNodes = append(currentNodes, n)
		}
	}
	stepsToEnd := make([]int, len(currentNodes))
	i, stepCount := 0, 0
	for !allCyclesFound(stepsToEnd) {
		stepCount += 1
		direction := rune(directions[i])
		for j, n := range currentNodes {
			switch direction {
			case 'L':
				currentNodes[j] = n.left
			case 'R':
				currentNodes[j] = n.right
			default:
				panic("invalid direction character")
			}
			if currentNodes[j].id[2] == 'Z' && stepsToEnd[j] == 0 {
				stepsToEnd[j] = stepCount
			}
		}
		i += 1
		if i == len(directions) {
			i = 0
		}
	}
	fmt.Println(lcm(stepsToEnd[0], stepsToEnd[1], stepsToEnd[2:]...))
}

func allCyclesFound(stepCounts []int) bool {
	for _, count := range stepCounts {
		if count == 0 {
			return false
		}
	}
	return true
}

type node struct {
	id          string
	left, right *node
}

func parseInput(input []string) (string, map[string]*node) {
	directions := input[0]
	re := regexp.MustCompile(`([\w\d]{3}) = \(([\w\d]{3}), ([\w\d]{3})\)`)
	nodes := make(map[string]*node, len(input)-2)
	for i := 2; i < len(input); i += 1 {
		match := re.FindStringSubmatch(input[i])
		id, left, right := match[1], match[2], match[3]
		currentNode := nodes[id]
		if currentNode == nil {
			n := node{id, nil, nil}
			nodes[id] = &n
			currentNode = &n
		}

		leftNode := nodes[left]
		if leftNode == nil {
			n := node{left, nil, nil}
			nodes[left] = &n
			leftNode = &n
		}
		currentNode.left = leftNode

		rightNode := nodes[right]
		if rightNode == nil {
			n := node{right, nil, nil}
			nodes[right] = &n
			rightNode = &n
		}
		currentNode.right = rightNode
	}
	return directions, nodes
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via gcd
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
