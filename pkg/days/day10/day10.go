package day10

import (
	"fmt"
	"log"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day10(path string) {
	fmt.Println("Day 10")
	input, err := utils.ReadInput(path + "/day10")
	if err != nil {
		log.Fatal(err)
	}
	start, adjList := parseInput(input)
	part1(start, adjList)
}

func part1(start position, adjList map[position][]position) {
	visitedPositions := make(map[position]bool)
	visitedPositions[start] = true
	stepCount := 1
	queue := new(positionQueue)
	adjacentPositions := adjList[start]
	for _, adjPos := range adjacentPositions {
		queue.Push(adjPos, stepCount)
	}
	for {
		currentPos, distance := queue.Dequeue()
		if visitedPositions[currentPos] {
			fmt.Println(distance)
			return
		}
		visitedPositions[currentPos] = true
		adjacentPositions = adjList[currentPos]
		stepCount = distance + 1
		for _, adjPos := range adjacentPositions {
			if !visitedPositions[adjPos] {
				queue.Push(adjPos, stepCount)
			}
		}

	}
}

func parseInput(input []string) (position, map[position][]position) {
	var start position
	adjList := make(map[position][]position, len(input)*len(input[0]))
	for i, line := range input {
		for j, tile := range line {
			currentPos := position{i, j}
			var adjacentPositions []position
			switch tile {
			case '|':
				adjacentPositions = upTile(adjacentPositions, position{i - 1, j}, i, input)
				adjacentPositions = downTile(adjacentPositions, position{i + 1, j}, i, input)
			case '-':
				adjacentPositions = rightTile(adjacentPositions, position{i, j + 1}, i, input)
				adjacentPositions = leftTile(adjacentPositions, position{i, j - 1}, i, input)
			case 'L':
				adjacentPositions = upTile(adjacentPositions, position{i - 1, j}, i, input)
				adjacentPositions = rightTile(adjacentPositions, position{i, j + 1}, i, input)
			case 'J':
				adjacentPositions = upTile(adjacentPositions, position{i - 1, j}, i, input)
				adjacentPositions = leftTile(adjacentPositions, position{i, j - 1}, i, input)
			case '7':
				adjacentPositions = downTile(adjacentPositions, position{i + 1, j}, i, input)
				adjacentPositions = leftTile(adjacentPositions, position{i, j - 1}, i, input)
			case 'F':
				adjacentPositions = downTile(adjacentPositions, position{i + 1, j}, i, input)
				adjacentPositions = rightTile(adjacentPositions, position{i, j + 1}, i, input)
			case 'S':
				start = currentPos
				adjacentPositions = upTile(adjacentPositions, position{i - 1, j}, i, input)
				adjacentPositions = downTile(adjacentPositions, position{i + 1, j}, i, input)
				adjacentPositions = rightTile(adjacentPositions, position{i, j + 1}, i, input)
				adjacentPositions = leftTile(adjacentPositions, position{i, j - 1}, i, input)
			case '.':
				continue
			default:
				panic("Invalid character in input.")
			}
			if len(adjacentPositions) > 0 {
				adjList[currentPos] = adjacentPositions
			}
		}
	}
	return start, adjList
}

type position struct {
	i, j int
}

func inBounds(pos position, maxI int, maxJ int) bool {
	return pos.i >= 0 && pos.i < maxI && pos.j >= 0 && pos.j < maxJ
}

func upTile(adjacentPositions []position, up position, i int, input []string) []position {
	if inBounds(up, len(input), len(input[i])) {
		adjTile := input[up.i][up.j]
		if adjTile == '|' || adjTile == '7' || adjTile == 'F' || adjTile == 'S' {
			adjacentPositions = append(adjacentPositions, up)
		}
	}
	return adjacentPositions
}

func downTile(adjacentPositions []position, down position, i int, input []string) []position {
	if inBounds(down, len(input), len(input[i])) {
		adjTile := input[down.i][down.j]
		if adjTile == '|' || adjTile == 'L' || adjTile == 'J' || adjTile == 'S' {
			adjacentPositions = append(adjacentPositions, down)
		}
	}
	return adjacentPositions
}

func rightTile(adjacentPositions []position, right position, i int, input []string) []position {
	if inBounds(right, len(input), len(input[i])) {
		adjTile := input[right.i][right.j]
		if adjTile == '-' || adjTile == 'J' || adjTile == '7' || adjTile == 'S' {
			adjacentPositions = append(adjacentPositions, right)
		}
	}
	return adjacentPositions
}

func leftTile(adjacentPositions []position, left position, i int, input []string) []position {
	if inBounds(left, len(input), len(input[i])) {
		adjTile := input[left.i][left.j]
		if adjTile == '-' || adjTile == 'F' || adjTile == 'L' || adjTile == 'S' {
			adjacentPositions = append(adjacentPositions, left)
		}
	}
	return adjacentPositions
}

type tuple struct {
	pos      position
	distance int
}

type positionQueue struct {
	q []tuple
}

func (pq *positionQueue) Push(pos position, distance int) {
	pq.q = append(pq.q, tuple{pos, distance})
}

func (pq *positionQueue) Dequeue() (position, int) {
	tup := pq.q[0]
	pq.q = pq.q[1:]
	return tup.pos, tup.distance
}
