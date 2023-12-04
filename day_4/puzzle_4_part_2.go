package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var answer = 0
var numbersSplitPattern, _ = regexp.Compile("\\s+")
var winningGuessedSplitPattern, _ = regexp.Compile("\\s+\\|\\s+")

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentCardNumber := 1
	collectedCards := make(map[int]int)

	for scanner.Scan() {
		inputLine := scanner.Text()
		cardInputLine := strings.Split(inputLine, ": ")[1]
		foundNumbers := []string{}

		collectedCards[currentCardNumber] = collectedCards[currentCardNumber] + 1

		winningNumbers := retrieveWinningNumbers(cardInputLine)
		guessedNumbers := retrieveGuessedNumbers(cardInputLine)

		for _, guessedNumber := range guessedNumbers {
			if len(winningNumbers) <= 0 {
				break
			}

			for i, winningNumber := range winningNumbers {
				if guessedNumber == winningNumber {
					foundNumbers = append(foundNumbers, guessedNumber)
					winningNumbers = removeFromSlice(winningNumbers, i)
					break
				}
			}
		}

		for i, _ := range foundNumbers {
			collectedCards[currentCardNumber+i+1] = collectedCards[currentCardNumber+i+1] + collectedCards[currentCardNumber]
		}

		currentCardNumber += 1
	}

	// Calculate answer
	for _, cards := range collectedCards {
		answer += cards
	}

	fmt.Println(answer)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func retrieveWinningNumbers(inputLine string) []string {
	winningNumbersString := winningGuessedSplitPattern.Split(inputLine, -1)[0]
	return numbersSplitPattern.Split(winningNumbersString, -1)
}

func retrieveGuessedNumbers(inputLine string) []string {
	guessedNumbersString := winningGuessedSplitPattern.Split(inputLine, -1)[1]
	return numbersSplitPattern.Split(guessedNumbersString, -1)
}

func removeFromSlice(slice []string, i int) []string {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
