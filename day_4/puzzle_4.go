package main

import (
	"bufio"
	"fmt"
	"math"
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

	for scanner.Scan() {
		inputLine := scanner.Text()
		cardInputLine := strings.Split(inputLine, ": ")[1]
		foundNumbers := []string{}

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

		answer += int(math.Pow(2.0, float64(len(foundNumbers)-1)))
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
