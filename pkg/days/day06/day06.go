package day06

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day06(path string) {
	fmt.Println("Day06")
	input, err := utils.ReadInput(path + "/day06")
	if err != nil {
		log.Fatal(err)
	}
	records := parseInput1(input)
	solve(records)
	record := parseInput2(input)
	records = []raceRecord{record}
	solve(records)
}

func solve(records []raceRecord) {
	var winCounts []int
	for _, race := range records {
		var possibleWins int
		for rate := 0; rate < race.time; rate += 1 {
			timeRemaining := race.time - rate
			distance := rate * timeRemaining
			if distance > race.distance {
				possibleWins += 1
			}
		}
		winCounts = append(winCounts, possibleWins)
	}
	var result int = 1
	for _, wins := range winCounts {
		result *= wins
	}
	fmt.Println(result)
}

type raceRecord struct {
	time, distance int
}

func parseInput1(input []string) []raceRecord {
	regex := regexp.MustCompile(`(\d+)`)
	times := regex.FindAllString(input[0], -1)
	distances := regex.FindAllString(input[1], -1)
	if len(times) != len(distances) {
		panic("number of time matches does not match number of distance matches")
	}
	records := make([]raceRecord, len(times))
	for i := 0; i < len(times); i += 1 {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal(err)
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal(err)
		}
		records[i] = raceRecord{time, distance}
	}
	return records
}

func parseInput2(input []string) raceRecord {
	var timeDigits []rune
	for _, char := range input[0] {
		if unicode.IsDigit(char) {
			timeDigits = append(timeDigits, char)
		}
	}
	var distanceDigits []rune
	for _, char := range input[1] {
		if unicode.IsDigit(char) {
			distanceDigits = append(distanceDigits, char)
		}
	}
	time, err := strconv.Atoi(string(timeDigits))
	if err != nil {
		log.Fatal(err)
	}
	distance, err := strconv.Atoi(string(distanceDigits))
	if err != nil {
		log.Fatal(err)
	}
	return raceRecord{time, distance}
}
