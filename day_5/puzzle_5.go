package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	Number      int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

var seeds = []*Seed{}

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds = retrieveSeeds(scanner.Text())

	sourceAttribute := ""
	destinationAttribute := ""

	for scanner.Scan() {
		inputLine := scanner.Text()
		if inputLine == "" {
			if destinationAttribute != "" {
				fillInUnmapped(sourceAttribute, destinationAttribute)
			}
			continue
		}

		if strings.Contains(inputLine, "map") {
			mappedAttributes := strings.Split(strings.Split(inputLine, " ")[0], "-")
			sourceAttribute = mappedAttributes[0]
			destinationAttribute = mappedAttributes[2]
			continue
		}

		mappingNumberStrings := strings.Split(inputLine, " ")
		rangeLength, _ := strconv.Atoi(mappingNumberStrings[2])

		destinationFrom, _ := strconv.Atoi(mappingNumberStrings[0])

		sourceFrom, _ := strconv.Atoi(mappingNumberStrings[1])
		sourceTo := sourceFrom + rangeLength

		for _, seed := range seeds {
			sourceValue := seed.getAttribute(sourceAttribute)
			destinationValue := seed.getAttribute(destinationAttribute)

			if destinationValue != -1 {
				continue
			}

			if sourceValue >= sourceFrom && sourceValue <= sourceTo {
				offset := sourceValue - sourceFrom
				seed.setAttribute(destinationAttribute, destinationFrom+offset)
			}
		}
	}

	fillInUnmapped(sourceAttribute, destinationAttribute)

	answer := seeds[0].Location

	for _, seed := range seeds {
		location := seed.Location

		if location < answer {
			answer = location
		}
	}

	fmt.Println(answer)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func retrieveSeeds(seedsInputLine string) []*Seed {
	seedsString := strings.Split(seedsInputLine, ": ")[1]
	seedStringsArray := strings.Split(seedsString, " ")
	seedsArray := []*Seed{}

	for _, seedString := range seedStringsArray {
		seed := Seed{-1, -1, -1, -1, -1, -1, -1, -1}
		seedNumber, _ := strconv.Atoi(seedString)
		seed.Number = seedNumber

		seedsArray = append(seedsArray, &seed)
	}

	return seedsArray
}

func fillInUnmapped(sourceAttribute string, destinationAttribute string) {
	for _, seed := range seeds {
		if seed.getAttribute(destinationAttribute) != -1 {
			continue
		}

		seed.setAttribute(destinationAttribute, seed.getAttribute(sourceAttribute))
	}
}

func (seed Seed) getAttribute(attribute string) int {
	switch attribute {
	case "seed":
		return seed.Number
	case "soil":
		return seed.Soil
	case "fertilizer":
		return seed.Fertilizer
	case "water":
		return seed.Water
	case "light":
		return seed.Light
	case "temperature":
		return seed.Temperature
	case "humidity":
		return seed.Humidity
	case "location":
		return seed.Location
	default:
		panic("Invalid seed attribute")
	}
}

func (seed *Seed) setAttribute(attribute string, value int) {
	switch attribute {
	case "seed":
		seed.Number = value
	case "soil":
		seed.Soil = value
	case "fertilizer":
		seed.Fertilizer = value
	case "water":
		seed.Water = value
	case "light":
		seed.Light = value
	case "temperature":
		seed.Temperature = value
	case "humidity":
		seed.Humidity = value
	case "location":
		seed.Location = value
	default:
		panic("Invalid seed attribute")
	}
}
