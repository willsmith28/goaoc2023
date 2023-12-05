package cmd

import (
	"flag"
	"fmt"

	"github.com/willsmith28/goaoc2023/pkg/days/day01"
	"github.com/willsmith28/goaoc2023/pkg/days/day02"
	"github.com/willsmith28/goaoc2023/pkg/days/day03"
)

func DayRunner() {
	inputDir := flag.String("input", "", "Path to input directory.")
	runAll := flag.Bool("runall", false, "Run all days")
	runDay01 := flag.Bool("day01", false, "Run all days")
	runDay02 := flag.Bool("day02", false, "Run all days")
	runDay03 := flag.Bool("day03", false, "Run all days")
	flag.Parse()
	if *runAll {
		*runDay01 = true
		*runDay02 = true
	}
	fmt.Println(*inputDir)
	if *runDay01 {
		day01.Day01(*inputDir)
	}
	if *runDay02 {
		day02.Day02(*inputDir)
	}
	if *runDay03 {
		day03.Day03(*inputDir)
	}
}
