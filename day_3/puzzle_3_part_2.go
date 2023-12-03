package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var numberRegexPattern, _ = regexp.Compile("\\d+")
var gearRegexPattern, _ = regexp.Compile("\\*")
var lineLength = 0
var answer = 0
var correctGearDistances = []int{0, 1, -1}

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	prevLine := ""
	currentLine := ""
	nextLine := ""
	var currentGearIndices [][]int

	var prevLineNumberIndices [][]int
	var currLineNumberIndices [][]int
	var nextLineNumberIndices [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputLine := scanner.Text()
		if currentLine == "" {
			currentLine = inputLine
			currLineNumberIndices = numberRegexPattern.FindAllStringIndex(currentLine, -1)
			continue
		} else {
			nextLine = inputLine
			nextLineNumberIndices = numberRegexPattern.FindAllStringIndex(nextLine, -1)
		}

		currentGearIndices = gearRegexPattern.FindAllStringIndex(currentLine, -1)

		if lineLength == 0 {
			lineLength = len(currentLine)
		}

		if prevLine != "" {
			prevLineNumberIndices = numberRegexPattern.FindAllStringIndex(prevLine, -1)
		}

		for _, gearIndices := range currentGearIndices {
			gearNumbers := []int{}
			gearIndex := gearIndices[0]

			if len(prevLineNumberIndices) != 0 {
				findGearNumbers(prevLineNumberIndices, gearIndex, prevLine, &gearNumbers)
				if len(gearNumbers) > 2 {
					continue
				}
			}

			findGearNumbers(currLineNumberIndices, gearIndex, currentLine, &gearNumbers)
			if len(gearNumbers) > 2 {
				continue
			}

			findGearNumbers(nextLineNumberIndices, gearIndex, nextLine, &gearNumbers)

			if len(gearNumbers) == 2 {
				answer += gearNumbers[0] * gearNumbers[1]
			}
		}

		// Move current line forward
		prevLine = currentLine
		currentLine = nextLine

		// Move number indices forward
		prevLineNumberIndices = currLineNumberIndices
		currLineNumberIndices = nextLineNumberIndices
	}

	// Evaluate last input line
	for _, gearIndices := range currentGearIndices {
		gearNumbers := []int{}
		gearIndex := gearIndices[0]

		if len(prevLineNumberIndices) != 0 {
			findGearNumbers(prevLineNumberIndices, gearIndex, prevLine, &gearNumbers)
			if len(gearNumbers) > 2 {
				continue
			}
		}

		findGearNumbers(currLineNumberIndices, gearIndex, currentLine, &gearNumbers)
		if len(gearNumbers) > 2 {
			continue
		}

		if len(gearNumbers) == 2 {
			answer += gearNumbers[0] * gearNumbers[1]
		}
	}
	fmt.Println(answer)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func addNewPartNumbersToAnswer(currentNumberIndices [][]int, currentLine string, partNumbersSet map[int]bool) {
	for index, _ := range partNumbersSet {
		partNumberIndices := currentNumberIndices[index]
		numberString := currentLine[partNumberIndices[0]:partNumberIndices[1]]
		number, _ := strconv.Atoi(numberString)
		answer += number
	}
}

func findGearNumbers(lineNumberIndices [][]int, gearIndex int, inputLine string, gearNumbers *[]int) {
	for _, numberIndices := range lineNumberIndices {
		if numberIndices[0] > gearIndex+1 {
			return
		}

		firstIndexDistance := numberIndices[0] - gearIndex
		lastIndexDistance := numberIndices[1] - 1 - gearIndex

		if slices.Contains(correctGearDistances, firstIndexDistance) || slices.Contains(correctGearDistances, lastIndexDistance) {
			numberString := inputLine[numberIndices[0]:numberIndices[1]]
			number, _ := strconv.Atoi(numberString)
			*gearNumbers = append(*gearNumbers, number)
		}

		if len(*gearNumbers) == 3 {
			return
		}
	}
}
