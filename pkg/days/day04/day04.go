package day04

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day04(path string) {
	fmt.Println("Day04")
	input, err := utils.ReadInput(path + "/day04")
	if err != nil {
		log.Fatal(err)
	}
	scratchCards := parseInput(input)
	part1(scratchCards)
	part2(scratchCards)
}

func part1(scratchCards []scratchCard) {
	var totalScore int
	for _, game := range scratchCards {
		var score int
		if game.matchCount > 0 {
			score = 1 << (game.matchCount - 1)
		}
		totalScore += score
	}
	fmt.Println(totalScore)

}

func part2(scratchCards []scratchCard) {
	cardCounts := make([]int, len(scratchCards))
	for i := range cardCounts {
		cardCounts[i] = 1
	}
	for i, game := range scratchCards {
		if game.matchCount == 0 {
			continue
		}
		for j := i + 1; j <= i+game.matchCount && j < len(cardCounts); j += 1 {
			cardCounts[j] += cardCounts[i]
		}
	}
	var totalCards int
	for _, cardCount := range cardCounts {
		totalCards += cardCount
	}
	fmt.Println(totalCards)
}

type scratchCard struct {
	winningNumbers, ourNumbers []int
	matchCount                 int
}

func parseInput(input []string) []scratchCard {
	scratchRegexp := regexp.MustCompile(`Card\s+\d+:\s+((?:\d+\s+)+)\|\s+((?:\d+\s*)+)`)
	var result []scratchCard
	for _, line := range input {
		match := scratchRegexp.FindStringSubmatch(line)
		var winningNumbers, ourNumbers []int

		for _, number := range strings.Split(match[1], " ") {
			if number == "" {
				continue
			}
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			winningNumbers = append(winningNumbers, num)
		}
		for _, number := range strings.Split(match[2], " ") {
			if number == "" {
				continue
			}
			num, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			ourNumbers = append(ourNumbers, num)
		}
		var matchCount int
		for _, number := range ourNumbers {
			if slices.Contains(winningNumbers, number) {
				matchCount += 1
			}
		}
		result = append(result, scratchCard{winningNumbers, ourNumbers, matchCount})
	}
	return result
}
