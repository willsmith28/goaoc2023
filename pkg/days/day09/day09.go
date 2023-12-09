package day09

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day09(path string) {
	fmt.Println("Day 09")
	input, err := utils.ReadInput(path + "/day09")
	if err != nil {
		log.Fatal(err)
	}
	readings := parseInput(input)
	part1(readings)
	part2(readings)
}

func part1(readings [][]int) {
	var sum int
	for _, reading := range readings {
		sum += predictNextReading(reading)
	}
	fmt.Println(sum)
}

func part2(readings [][]int) {
	var sum int
	for _, reading := range readings {
		sum += predictPrevReading(reading)
	}
	fmt.Println(sum)
}

func predictNextReading(reading []int) int {
	if allTheSame(reading) {
		return reading[len(reading)-1]
	}
	sequenceDifferences := make([]int, len(reading)-1)
	for i := 1; i < len(reading); i += 1 {
		difference := reading[i] - reading[i-1]
		sequenceDifferences[i-1] = difference
	}
	return reading[len(reading)-1] + predictNextReading(sequenceDifferences)
}

func predictPrevReading(reading []int) int {
	if allTheSame(reading) {
		return reading[0]
	}
	sequenceDifferences := make([]int, len(reading)-1)
	for i := 1; i < len(reading); i += 1 {
		difference := reading[i] - reading[i-1]
		sequenceDifferences[i-1] = difference
	}
	return reading[0] - predictPrevReading(sequenceDifferences)
}

func allTheSame(numbers []int) bool {
	if len(numbers) == 0 {
		panic("numbers is empty")
	}
	first := numbers[0]
	for _, number := range numbers {
		if number != first {
			return false
		}
	}
	return true
}

func parseInput(input []string) [][]int {
	result := make([][]int, len(input))
	for i, line := range input {
		readings := strings.Split(line, " ")
		for _, number := range readings {
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			result[i] = append(result[i], n)
		}
	}
	return result
}
