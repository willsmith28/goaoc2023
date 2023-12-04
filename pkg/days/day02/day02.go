package day02

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

type gameResults struct {
	redPulls, greenPulls, bluePulls []int
}

func Day02(path string) {
	input, err := utils.ReadInput(path + "/day02")
	if err != nil {
		log.Fatal(err)
	}
	results := gameParser(input)
	fmt.Println("Day 02")
	part1(results)
	part2(results)
}

func gameParser(input []string) []gameResults {
	var redRegexp *regexp.Regexp = regexp.MustCompile(`(\d+) red`)
	var greenRegexp *regexp.Regexp = regexp.MustCompile(`(\d+) green`)
	var blueRegexp *regexp.Regexp = regexp.MustCompile(`(\d+) blue`)
	var results []gameResults
	for _, line := range input {
		var gameResult gameResults
		redMatches := redRegexp.FindAllStringSubmatch(line, -1)
		for _, redMatch := range redMatches {
			redCount, err := strconv.Atoi(redMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			gameResult.redPulls = append(gameResult.redPulls, redCount)
		}

		greenMatches := greenRegexp.FindAllStringSubmatch(line, -1)
		for _, greenMatch := range greenMatches {
			greenCount, err := strconv.Atoi(greenMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			gameResult.greenPulls = append(gameResult.greenPulls, greenCount)
		}

		blueMatches := blueRegexp.FindAllStringSubmatch(line, -1)
		for _, blueMatch := range blueMatches {
			blueCount, err := strconv.Atoi(blueMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			gameResult.bluePulls = append(gameResult.bluePulls, blueCount)
		}

		results = append(results, gameResult)
	}
	return results
}

func part1(input []gameResults) {
	totalRed := 12
	totalGreen := 13
	totalBlue := 14
	var possibleGames []int
	for i, pulls := range input {
		redPossible := true
		for _, redCount := range pulls.redPulls {
			if redCount > totalRed {
				redPossible = false
				break
			}
		}
		if !redPossible {
			continue
		}

		greenPossible := true
		for _, greenCount := range pulls.greenPulls {
			if greenCount > totalGreen {
				greenPossible = false
				break
			}
		}
		if !greenPossible {
			continue
		}

		bluePossible := true
		for _, blueCount := range pulls.bluePulls {
			if blueCount > totalBlue {
				bluePossible = false
				break
			}
		}
		if bluePossible {
			possibleGames = append(possibleGames, i+1)
		}
	}

	var sum int
	for _, gameID := range possibleGames {
		sum += gameID
	}
	fmt.Println(sum)
}

func part2(input []gameResults) {
	var cubePower []int
	for _, pulls := range input {
		minRedCubes := slices.Max(pulls.redPulls)
		minGreenCubes := slices.Max(pulls.greenPulls)
		minBlueCubes := slices.Max(pulls.bluePulls)
		cubePower = append(cubePower, minRedCubes*minGreenCubes*minBlueCubes)
	}
	var sum int
	for _, power := range cubePower {
		sum += power
	}
	fmt.Println(sum)
}
