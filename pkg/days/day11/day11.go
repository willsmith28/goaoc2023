package day11

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day11(path string) {
	fmt.Println("Day 11")
	input, err := utils.ReadInput(path + "/day11")
	if err != nil {
		log.Fatal(err)
	}
	emptyRows, emptyColumns, galaxies := parseInput(input)
	solve(emptyRows, emptyColumns, galaxies, 1)
	solve(emptyRows, emptyColumns, galaxies, 999_999)

}

func solve(emptyRows []int, emptyColumns []int, galaxies []position, expansionFactor int) {
	var galaxyPairs [][2]position
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxyPairs = append(galaxyPairs, [2]position{galaxies[i], galaxies[j]})
		}
	}

	var sum int
	for _, pair := range galaxyPairs {
		minI, maxI := min(pair[0].i, pair[1].i), max(pair[0].i, pair[1].i)
		minJ, maxJ := min(pair[0].j, pair[1].j), max(pair[0].j, pair[1].j)
		var iExpansions, jExpansions int
		for i := minI; i < maxI; i++ {
			if slices.Contains(emptyRows, i) {
				iExpansions += expansionFactor
			}
		}
		for j := minJ; j < maxJ; j++ {
			if slices.Contains(emptyColumns, j) {
				jExpansions += expansionFactor
			}
		}

		distance := (abs(pair[0].i-pair[1].i) + iExpansions) + (abs(pair[0].j-pair[1].j) + jExpansions)
		sum += distance
	}
	fmt.Println(sum)
}

func parseInput(input []string) (emptyRows []int, emptyColumns []int, galaxies []position) {
	// find all empty rows
	for i, line := range input {
		if !strings.Contains(line, "#") {
			emptyRows = append(emptyRows, i)
		}
	}
	// find all empty columns
	for j := 0; j < len(input[0]); j++ {
		var containsGalaxy bool
		for _, line := range input {
			if line[j] == '#' {
				containsGalaxy = true
			}
		}
		if !containsGalaxy {
			emptyColumns = append(emptyColumns, j)
		}
	}
	// find all galaxies
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, position{i, j})
			}
		}
	}
	return
}

type position struct {
	i, j int
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
