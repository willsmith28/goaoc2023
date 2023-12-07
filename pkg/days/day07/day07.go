package day07

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day07(path string) {
	fmt.Println("Day07")
	input, err := utils.ReadInput(path + "/day07")
	if err != nil {
		log.Fatal(err)
	}
	solve(input, 1)
	solve(input, 2)
}

func solve(input []string, part int) {
	hands := parseInput(input, part)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Less(&hands[j], part)
	})
	var totalWinnings int
	for i, h := range hands {
		rank := i + 1
		totalWinnings += (h.bid * rank)
	}
	fmt.Println(totalWinnings)
}

var cardValue1 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}
var cardValue2 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards    string
	handType int
	bid      int
}

// lhs < rhs
func (lhs *hand) Less(rhs *hand, part int) bool {
	if part != 1 && part != 2 {
		panic("invalid part number")
	}
	if lhs.handType != rhs.handType {
		return lhs.handType < rhs.handType
	}
	for i := 0; i < 5; i += 1 {
		lhsChar := rune(lhs.cards[i])
		rhsChar := rune(rhs.cards[i])
		if lhsChar == rhsChar {
			continue
		}
		if part == 1 {
			return cardValue1[lhsChar] < cardValue1[rhsChar]
		}
		return cardValue2[lhsChar] < cardValue2[rhsChar]

	}
	return false
}

func parseInput(input []string, part int) []hand {
	if part != 1 && part != 2 {
		panic("invalid part number")
	}
	hands := make([]hand, len(input))
	for i, line := range input {
		lineSplit := strings.Split(line, " ")
		bid, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			log.Fatal(err)
		}

		cardCounter := make(map[rune]int, 5)
		for _, card := range lineSplit[0] {
			cardCounter[card] += 1
		}
		if part == 2 {
			_, containsJoker := cardCounter['J']
			if containsJoker {
				jokerCount := cardCounter['J']
				delete(cardCounter, 'J')
				var mostCommon rune
				var commonCount int
				for card, count := range cardCounter {
					if count > commonCount {
						commonCount = count
						mostCommon = card
					}
				}
				cardCounter[mostCommon] += jokerCount
			}
		}
		var handType int
		if len(cardCounter) == 1 {
			handType = fiveOfAKind
		} else if len(cardCounter) == 2 {
			counts := make([]int, 2)
			var j int
			for _, val := range cardCounter {
				counts[j] = val
				j += 1
			}
			slices.Sort(counts)
			if counts[0] == 1 && counts[1] == 4 {
				handType = fourOfAKind
			} else if counts[0] == 2 && counts[1] == 3 {
				handType = fullHouse
			} else {
				panic("Unexpected values for cardCount == 2")
			}
		} else if len(cardCounter) == 3 {
			counts := make([]int, 3)
			var j int
			for _, val := range cardCounter {
				counts[j] = val
				j += 1
			}
			slices.Sort(counts)
			if counts[0] == 1 && counts[1] == 1 && counts[2] == 3 {
				handType = threeOfAKind
			} else if counts[0] == 1 && counts[1] == 2 && counts[2] == 2 {
				handType = twoPair
			} else {
				panic("Unexpected values for cardCount == 3")
			}
		} else if len(cardCounter) == 4 {
			handType = onePair
		} else if len(cardCounter) == 5 {
			handType = highCard
		} else {
			panic("invalid hand")
		}

		hands[i].cards = lineSplit[0]
		hands[i].bid = bid
		hands[i].handType = handType
	}
	return hands
}
