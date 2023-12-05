package day05

import (
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/willsmith28/goaoc2023/pkg/utils"
)

func Day05(path string) {
	fmt.Println("Day05")
	input, err := utils.ReadInput(path + "/day05")
	if err != nil {
		log.Fatal(err)
	}
	a := parseAlmanac(input)
	part1(&a)
	start := time.Now()
	part2(&a)
	duration := time.Since(start)
	fmt.Println("part 2 took", duration)
}

func part1(a *almanac) {
	var locations []int
	for _, seed := range a.seeds {
		soil := a.sourceToDestination("seedToSoil", seed)
		fertilizer := a.sourceToDestination("soilToFertilizer", soil)
		water := a.sourceToDestination("fertilizerToWater", fertilizer)
		light := a.sourceToDestination("waterToLight", water)
		temp := a.sourceToDestination("lightToTemp", light)
		humidity := a.sourceToDestination("tempToHumidity", temp)
		location := a.sourceToDestination("humidityToLocation", humidity)
		locations = append(locations, location)
	}
	fmt.Println(slices.Min(locations))
}

func part2(a *almanac) {
	minLocation := math.MaxInt
	for i := 0; i < len(a.seeds); i += 2 {
		start, length := a.seeds[i], a.seeds[i+1]
		for seed := start; seed < start+length; seed += 1 {
			soil := a.sourceToDestination("seedToSoil", seed)
			fertilizer := a.sourceToDestination("soilToFertilizer", soil)
			water := a.sourceToDestination("fertilizerToWater", fertilizer)
			light := a.sourceToDestination("waterToLight", water)
			temp := a.sourceToDestination("lightToTemp", light)
			humidity := a.sourceToDestination("tempToHumidity", temp)
			location := a.sourceToDestination("humidityToLocation", humidity)
			if location < minLocation {
				minLocation = location
			}
		}
	}
	fmt.Println(minLocation)
}

type mapRange struct {
	destinationStart, destinationEnd, sourceStart, sourceEnd int
}

type almanac struct {
	seeds                                                                                                          []int
	seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation []mapRange
}

func (a *almanac) sourceToDestination(mapName string, source int) int {
	var requestedMap *[]mapRange
	switch mapName {
	case "seedToSoil":
		requestedMap = &a.seedToSoil
	case "soilToFertilizer":
		requestedMap = &a.soilToFertilizer
	case "fertilizerToWater":
		requestedMap = &a.fertilizerToWater
	case "waterToLight":
		requestedMap = &a.waterToLight
	case "lightToTemp":
		requestedMap = &a.lightToTemp
	case "tempToHumidity":
		requestedMap = &a.tempToHumidity
	case "humidityToLocation":
		requestedMap = &a.humidityToLocation

	default:
		panic("invalid mapName")
	}

	for _, mRange := range *requestedMap {
		if mRange.sourceStart <= source && source < mRange.sourceEnd {
			distance := source - mRange.sourceStart
			return mRange.destinationStart + distance
		}
	}
	return source
}

func parseRangeLine(line string) mapRange {
	vals := strings.Split(line, " ")
	if len(vals) != 3 {
		panic("strings.Split produced more than 3 items")
	}
	destinationStart, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal(err)
	}
	sourceStart, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal(err)
	}
	rangeLength, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Fatal(err)
	}
	return mapRange{
		destinationStart,
		destinationStart + rangeLength,
		sourceStart,
		sourceStart + rangeLength,
	}
}

func parseAlmanac(input []string) almanac {
	// read seeds
	var seeds []int
	for _, num := range strings.Split(input[0][7:], " ") {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, number)
	}
	// read seeds to soil
	i := 3
	var seedToSoilLines []string
	for input[i] != "" {
		seedToSoilLines = append(seedToSoilLines, input[i])
		i += 1
	}
	var seedToSoil []mapRange
	for _, line := range seedToSoilLines {
		seedToSoil = append(seedToSoil, parseRangeLine(line))
	}
	// read soil to fertilizer
	i += 2
	var soilToFertilizerLines []string
	for input[i] != "" {
		soilToFertilizerLines = append(soilToFertilizerLines, input[i])
		i += 1
	}
	var soilToFertilizer []mapRange
	for _, line := range soilToFertilizerLines {
		soilToFertilizer = append(soilToFertilizer, parseRangeLine(line))
	}
	// read fertilizer to water
	i += 2
	var fertilizerToWaterLines []string
	for input[i] != "" {
		fertilizerToWaterLines = append(fertilizerToWaterLines, input[i])
		i += 1
	}
	var fertilizerToWater []mapRange
	for _, line := range fertilizerToWaterLines {
		fertilizerToWater = append(fertilizerToWater, parseRangeLine(line))
	}
	// read water to light
	i += 2
	var waterToLightLines []string
	for input[i] != "" {
		waterToLightLines = append(waterToLightLines, input[i])
		i += 1
	}
	var waterToLight []mapRange
	for _, line := range waterToLightLines {
		waterToLight = append(waterToLight, parseRangeLine(line))
	}
	// read light to temperature
	i += 2
	var lightToTempLines []string
	for input[i] != "" {
		lightToTempLines = append(lightToTempLines, input[i])
		i += 1
	}
	var lightToTemp []mapRange
	for _, line := range lightToTempLines {
		lightToTemp = append(lightToTemp, parseRangeLine(line))
	}
	// read temp to humidity
	i += 2
	var tempToHumidityLines []string
	for input[i] != "" {
		tempToHumidityLines = append(tempToHumidityLines, input[i])
		i += 1
	}
	var tempToHumidity []mapRange
	for _, line := range tempToHumidityLines {
		tempToHumidity = append(tempToHumidity, parseRangeLine(line))
	}
	// read humidity to location
	i += 2
	var humidityToLocationLines []string
	for i < len(input) && input[i] != "" {
		humidityToLocationLines = append(humidityToLocationLines, input[i])
		i += 1
	}
	var humidityToLocation []mapRange
	for _, line := range humidityToLocationLines {
		humidityToLocation = append(humidityToLocation, parseRangeLine(line))
	}
	sort.Slice(seedToSoil, func(i, j int) bool {
		return seedToSoil[i].sourceStart < seedToSoil[j].sourceStart
	})

	sort.Slice(soilToFertilizer, func(i, j int) bool {
		return soilToFertilizer[i].sourceStart < soilToFertilizer[j].sourceStart
	})

	sort.Slice(fertilizerToWater, func(i, j int) bool {
		return fertilizerToWater[i].sourceStart < fertilizerToWater[j].sourceStart
	})

	sort.Slice(waterToLight, func(i, j int) bool {
		return waterToLight[i].sourceStart < waterToLight[j].sourceStart
	})

	sort.Slice(lightToTemp, func(i, j int) bool {
		return lightToTemp[i].sourceStart < lightToTemp[j].sourceStart
	})

	sort.Slice(tempToHumidity, func(i, j int) bool {
		return tempToHumidity[i].sourceStart < tempToHumidity[j].sourceStart
	})

	sort.Slice(humidityToLocation, func(i, j int) bool {
		return humidityToLocation[i].sourceStart < humidityToLocation[j].sourceStart
	})

	return almanac{seeds, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation}
}
