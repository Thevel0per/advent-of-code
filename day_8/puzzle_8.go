package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var answer = 0

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_8/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftRightMapping := make(map[string][]string)

	scanner.Scan()
	directionsString := scanner.Text()

	currentLocation := "AAA"
	currentDirectionIndex := 0

	for scanner.Scan() {
		inputLine := scanner.Text()

		if len(inputLine) == 0 {
			continue
		}

		splitInputLine := strings.Split(inputLine, " = ")

		locationKey := splitInputLine[0]
		destinationValues := createLeftRightArray(splitInputLine[1])
		leftRightMapping[locationKey] = destinationValues
	}

	for currentLocation != "ZZZ" {
		currentDirection := string(directionsString[currentDirectionIndex])
		nextAvailableLocations := leftRightMapping[currentLocation]
		nextLocation := resolveNewLocation(nextAvailableLocations, currentDirection)

		currentLocation = nextLocation
		currentDirectionIndex += 1

		if currentDirectionIndex == len(directionsString) {
			currentDirectionIndex = 0
		}

		answer += 1
	}

	fmt.Println(answer)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func createLeftRightArray(leftRightString string) []string {
	destinationValuesString := strings.Replace(strings.Replace(leftRightString, "(", "", 1), ")", "", 1)
	return strings.Split(destinationValuesString, ", ")
}

func resolveNewLocation(availableLocations []string, direction string) string {
	if direction == "L" {
		return availableLocations[0]
	} else if direction == "R" {
		return availableLocations[1]
	} else {
		return ""
	}
}
