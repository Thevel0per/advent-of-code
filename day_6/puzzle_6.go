package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var answer = 1
var numbersSplitPattern, _ = regexp.Compile("\\s+")
var labelNumbersSplitPattern, _ = regexp.Compile(":\\s+")

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	raceData := [][]string{}
	highScoreBeatWays := []int{}

	for scanner.Scan() {
		inputLine := scanner.Text()
		raceDataLine := labelNumbersSplitPattern.Split(inputLine, -1)[1]
		raceDataNumbers := numbersSplitPattern.Split(raceDataLine, -1)
		raceData = append(raceData, raceDataNumbers)
	}

	for i := 0; i < len(raceData[0]); i += 1 {
		raceTime, _ := strconv.Atoi(raceData[0][i])
		raceHighScore, _ := strconv.Atoi(raceData[1][i])

		leftBound := 0 + 1
		rightBound := raceTime - 1

		for {
			leftScore := leftBound * (raceTime - leftBound)
			rightScore := rightBound * (raceTime - rightBound)

			if leftScore <= raceHighScore {
				leftBound += 1
			}

			if rightScore <= raceHighScore {
				rightBound -= 1
			}

			if leftScore > raceHighScore && rightScore > raceHighScore {
				break
			}
		}

		highScoreBeatWays = append(highScoreBeatWays, rightBound-leftBound+1)
	}

	for _, result := range highScoreBeatWays {
		answer *= result
	}

	fmt.Println(answer)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
