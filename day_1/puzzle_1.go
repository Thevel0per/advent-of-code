package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

var digitsMapping map[string]string = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

var digitInputs []string = []string{"\\d", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_1/input.txt")
	if err != nil  {
		panic(err)
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
	answer := 0

    for scanner.Scan() {
        inputLine := scanner.Text()
		
		firstNumberString := findDigit(true, inputLine)
		lastNumberString := findDigit(false, inputLine)

		nextNumber, _ := strconv.Atoi(firstNumberString + lastNumberString)

		answer += nextNumber
    }

	fmt.Println(answer)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func findDigit(first bool, inputLine string) string {
	minIndex := len(inputLine)
	maxIndex := -1

	var foundDigitIndexis []int


	if first {
		for _, regexString := range digitInputs { 
			regexPattern, _ := regexp.Compile(regexString)
			foundIndexis := regexPattern.FindStringIndex(inputLine)

			if foundIndexis != nil && minIndex > foundIndexis[0] {
				minIndex = foundIndexis[0]
				foundDigitIndexis = foundIndexis
			}
		}
	} else {
		for _, regexString := range digitInputs { 
			regexPattern, _ := regexp.Compile(regexString)
			allFoundIndexis := regexPattern.FindAllStringSubmatchIndex(inputLine, -1)

			if allFoundIndexis != nil {
				foundIndexis := allFoundIndexis[len(allFoundIndexis) - 1]

				if maxIndex < foundIndexis[0] {
					maxIndex = foundIndexis[0]
					foundDigitIndexis = foundIndexis
				}
			}
		}
	}
	return convertDigitWordToDigit(inputLine[foundDigitIndexis[0]:foundDigitIndexis[1]])
}

func convertDigitWordToDigit(digitString string) string {
	regexPattern, _ := regexp.Compile("\\d")
	if regexPattern.MatchString(digitString) {
		return digitString
	} else {
		return digitsMapping[digitString]
	}
}