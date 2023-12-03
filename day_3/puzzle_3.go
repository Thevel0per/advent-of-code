package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

var numberRegexPattern, _ = regexp.Compile("\\d+")
var symbolRegexPattern, _ = regexp.Compile("[^0-9\\.]")
var lineLength = 0
var answer = 0

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_3/input.txt")
	if err != nil  {
		panic(err)
	}
	defer file.Close()

	prevLine := ""
	currentLine := ""
	nextLine := ""
	var currentNumberIndices [][]int

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        inputLine := scanner.Text()
		if currentLine == "" {
			currentLine = inputLine
			continue
		} else {
			nextLine = inputLine
		}

		var partNumbersSet map[int]bool = make(map[int]bool) 
		
		currentNumberIndices = numberRegexPattern.FindAllStringIndex(currentLine, -1)

		if lineLength == 0 {
			lineLength = len(currentLine)
		}

		// Check previous line symbols to find part numbers
		if prevLine != "" {
			checkPartNumbersAgainstLine(currentNumberIndices, prevLine, partNumbersSet)
		}

		// Check current line symbols to find part numbers
		checkPartNumbersAgainstLine(currentNumberIndices, currentLine, partNumbersSet)

		// Check next line symbols to find part numbers
		checkPartNumbersAgainstLine(currentNumberIndices, nextLine, partNumbersSet)

		// Add part numbers to the answer
		addNewPartNumbersToAnswer(currentNumberIndices, currentLine, partNumbersSet)

		// Move current line forward
		prevLine = currentLine
		currentLine = nextLine
    }

	// Evaluate last input line
	var partNumbersSet map[int]bool = make(map[int]bool) 
	currentNumberIndices = numberRegexPattern.FindAllStringIndex(currentLine, -1)
	checkPartNumbersAgainstLine(currentNumberIndices, prevLine, partNumbersSet)
	checkPartNumbersAgainstLine(currentNumberIndices, currentLine, partNumbersSet)
	addNewPartNumbersToAnswer(currentNumberIndices, currentLine, partNumbersSet)

	fmt.Println(answer)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func addNewPartNumbersToAnswer(currentNumberIndices [][]int, currentLine string, partNumbersSet map[int]bool){
	for index, _ := range partNumbersSet {
		partNumberIndices := currentNumberIndices[index]
		numberString := currentLine[partNumberIndices[0]:partNumberIndices[1]]
		number, _ := strconv.Atoi(numberString)
		answer += number
	}
}

func checkPartNumbersAgainstLine(currentNumberIndices [][]int, inputLine string, partNumbersSet map[int]bool) {
	for index, numberIndices := range currentNumberIndices {
		var sliceBeginIndex int
		var sliceEndIndex int

		if numberIndices[0] > 0 {
			sliceBeginIndex = numberIndices[0] - 1
		} else {
			sliceBeginIndex = numberIndices[0]
		}

		if numberIndices[1] < lineLength {
			sliceEndIndex = numberIndices[1] + 1
		} else {
			sliceEndIndex = numberIndices[1]
		}

		lineSlice := inputLine[sliceBeginIndex:sliceEndIndex]
		if symbolRegexPattern.MatchString(lineSlice) {
			partNumbersSet[index] = true
		}
	}
}