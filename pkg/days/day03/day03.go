package day03

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day03(path string) {
	fmt.Println("Day 03")
	input, err := utils.ReadInput(path + "/day03")
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)
}

type partNumber struct {
	number int
	valid  bool
}

func parsePartNumbers(schema []string) []partNumber {
	var result []partNumber

	for i, line := range schema {
		var digits []rune
		start, end := -1, -1
		for j, char := range line {
			if unicode.IsDigit(char) {
				if len(digits) == 0 {
					start = j
				}
				digits = append(digits, char)
				if j == len(line)-1 {
					end = j + 1
					partNum, err := strconv.Atoi(string(digits))
					if err != nil {
						log.Fatal(err)
					}
					var valid bool
					for y := i - 1; y <= i+1; y += 1 {
						if y < 0 || y >= len(schema) {
							continue
						}
						for x := start - 1; x <= end; x += 1 {
							if x < 0 || x >= len(schema[y]) {
								continue
							}
							potentialSymbol := rune(schema[y][x])
							if !unicode.IsDigit(potentialSymbol) && potentialSymbol != '.' {
								valid = true
							}
						}
					}
					result = append(result, partNumber{partNum, valid})
				}

			} else if len(digits) != 0 {
				end = j
				partNum, err := strconv.Atoi(string(digits))
				if err != nil {
					log.Fatal(err)
				}
				var valid bool
				for y := i - 1; y <= i+1; y += 1 {
					if y < 0 || y >= len(schema) {
						continue
					}
					for x := start - 1; x <= end; x += 1 {
						if x < 0 || x >= len(schema[y]) {
							continue
						}
						potentialSymbol := rune(schema[y][x])
						if !unicode.IsDigit(potentialSymbol) && potentialSymbol != '.' {
							valid = true
						}
					}
				}

				result = append(result, partNumber{partNum, valid})
				digits = nil
				start, end = -1, -1
			}
		}
	}

	return result
}

func part1(schema []string) {
	var sum int
	numbers := parsePartNumbers(schema)
	for _, number := range numbers {
		if number.valid {
			sum += number.number
		}
	}
	fmt.Println(sum)
}

type position struct {
	i, j int
}

type gear struct {
	number int
	used   bool
}

func parseSchemaMap(schema []string) [][]*gear {
	result := make([][]*gear, len(schema))
	for i := 0; i < len(schema); i += 1 {
		result[i] = make([]*gear, len(schema[i]))
	}

	for i, line := range schema {
		var digits []rune
		start, end := -1, -1
		for j, char := range line {
			if unicode.IsDigit(char) {
				if len(digits) == 0 {
					start = j
				}
				digits = append(digits, char)
				if j == len(line)-1 {
					end = j + 1
					partNum, err := strconv.Atoi(string(digits))
					if err != nil {
						log.Fatal(err)
					}

					foundGear := gear{partNum, false}
					for n := start; n < end; n += 1 {
						result[i][n] = &foundGear
					}
				}

			} else if len(digits) != 0 {
				end = j
				partNum, err := strconv.Atoi(string(digits))
				if err != nil {
					log.Fatal(err)
				}
				foundGear := gear{partNum, false}
				for n := start; n < end; n += 1 {
					result[i][n] = &foundGear
				}
				digits = nil
				start, end = -1, -1
			}
		}
	}

	return result
}

func part2(schema []string) {
	schemaMap := parseSchemaMap(schema)
	var gearRatios []int
	for i, line := range schema {
		for j, symbol := range line {
			if symbol != '*' {
				continue
			}
			adjacencies := [8]position{
				{i - 1, j - 1},
				{i - 1, j},
				{i - 1, j + 1},
				{i, j - 1},
				{i, j + 1},
				{i + 1, j - 1},
				{i + 1, j},
				{i + 1, j + 1},
			}
			var gears []*gear
			for _, adjacentPos := range adjacencies {
				if adjacentPos.i < 0 || adjacentPos.i >= len(schema) || adjacentPos.j < 0 || adjacentPos.j >= len(line) {
					continue
				}
				possibleGear := schemaMap[adjacentPos.i][adjacentPos.j]
				if possibleGear != nil && !possibleGear.used {
					gears = append(gears, possibleGear)
					possibleGear.used = true
				}
			}
			if len(gears) >= 2 {
				var gearRatio int = 1
				for _, g := range gears {
					gearRatio *= g.number
				}
				gearRatios = append(gearRatios, gearRatio)
			}
		}
	}
	var sum int
	for _, ratio := range gearRatios {
		sum += ratio
	}
	fmt.Println(sum)
}
