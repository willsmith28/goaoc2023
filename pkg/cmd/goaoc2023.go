package cmd

import (
	"flag"
	"fmt"

	"github.com/willsmith28/goaoc2023/pkg/days"
)

func DayRunner() {
	inputDir := flag.String("input", "", "Path to input directory.")
	runAll := flag.Bool("runall", false, "Run all days")
	day01 := flag.Bool("day01", false, "Run all days")
	flag.Parse()
	if *runAll {
		*day01 = true
	}
	fmt.Println(*inputDir)
	if *day01 {
		days.Day01(*inputDir)
	}
}
