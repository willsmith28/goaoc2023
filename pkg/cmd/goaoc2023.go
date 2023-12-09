package cmd

import (
	"flag"
	"fmt"

	"github.com/willsmith28/goaoc2023/pkg/days/day01"
	"github.com/willsmith28/goaoc2023/pkg/days/day02"
	"github.com/willsmith28/goaoc2023/pkg/days/day03"
	"github.com/willsmith28/goaoc2023/pkg/days/day04"
	"github.com/willsmith28/goaoc2023/pkg/days/day05"
	"github.com/willsmith28/goaoc2023/pkg/days/day06"
	"github.com/willsmith28/goaoc2023/pkg/days/day07"
	"github.com/willsmith28/goaoc2023/pkg/days/day08"
	"github.com/willsmith28/goaoc2023/pkg/days/day09"
)

func DayRunner() {
	inputDir := flag.String("input", "", "Path to input directory.")
	runAll := flag.Bool("runall", false, "Run all days")
	runDay01 := flag.Bool("day01", false, "Run Day01")
	runDay02 := flag.Bool("day02", false, "Run Day02")
	runDay03 := flag.Bool("day03", false, "Run Day03")
	runDay04 := flag.Bool("day04", false, "Run Day04")
	runDay05 := flag.Bool("day05", false, "Run Day05")
	runDay06 := flag.Bool("day06", false, "Run Day06")
	runDay07 := flag.Bool("day07", false, "Run Day07")
	runDay08 := flag.Bool("day08", false, "Run Day08")
	runDay09 := flag.Bool("day09", false, "Run Day09")
	flag.Parse()

	if *runAll {
		*runDay01 = true
		*runDay02 = true
		*runDay03 = true
		*runDay04 = true
		*runDay05 = true
		*runDay06 = true
		*runDay07 = true
		*runDay08 = true
		*runDay09 = true
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
	if *runDay04 {
		day04.Day04(*inputDir)
	}
	if *runDay05 {
		day05.Day05(*inputDir)
	}
	if *runDay06 {
		day06.Day06(*inputDir)
	}
	if *runDay07 {
		day07.Day07(*inputDir)
	}
	if *runDay08 {
		day08.Day08(*inputDir)
	}
	if *runDay09 {
		day09.Day09(*inputDir)
	}
}
