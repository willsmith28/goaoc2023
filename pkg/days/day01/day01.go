package day01

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day01(path string) {
	input, err := utils.ReadInput(path + "/day01")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 01")
	part1(input)
	part2(input)
}

func part1(input []string) {
	numbers := make([]int, len(input))
	for _, line := range input {
		var digits []rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}

		number, err := strconv.Atoi(string([]rune{digits[0], digits[len(digits)-1]}))
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}

func part2(input []string) {
	digits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
	}
	digits_regexp := make(map[string]*regexp.Regexp, len(digits))
	for digit := range digits {
		digits_regexp[digit] = regexp.MustCompile(digit)
	}
	var numbers []int
	for _, line := range input {
		var matchedDigits []foundDigit
		for digit := range digits {
			current_regexp := digits_regexp[digit]
			matches := current_regexp.FindAllStringIndex(line, -1)
			for _, match := range matches {
				matchedDigits = append(matchedDigits, foundDigit{match[0], digit})
			}
		}
		min, max := minMax(matchedDigits)
		number, err := strconv.Atoi(digits[min.digit] + digits[max.digit])
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	var sum int
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}

type foundDigit struct {
	index int
	digit string
}

func minMax(array []foundDigit) (foundDigit, foundDigit) {
	min := array[0]
	max := array[0]
	for _, value := range array {
		if max.index < value.index {
			max = value
		}
		if min.index > value.index {
			min = value
		}
	}
	return min, max
}
