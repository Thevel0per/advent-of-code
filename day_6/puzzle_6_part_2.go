package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var answer = 1
var labelNumbersSplitPattern, _ = regexp.Compile(":\\s+")

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	raceData := []string{}

	for scanner.Scan() {
		inputLine := scanner.Text()
		raceDataLine := labelNumbersSplitPattern.Split(inputLine, -1)[1]
		raceData = append(raceData, raceDataLine)
	}

	raceTime, _ := strconv.Atoi(strings.Replace(raceData[0], " ", "", -1))
	raceHighScore, _ := strconv.Atoi(strings.Replace(raceData[1], " ", "", -1))

	leftBound := findLeftBound(raceTime, raceHighScore)
	rightBound := findRightBound(raceTime, raceHighScore, leftBound)

	fmt.Println(rightBound - leftBound + 1)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func findRightBound(raceTime int, highScore int, leftBound int) int {
	low := leftBound
	high := raceTime - 1

	for {
		mid := low + (high-low)/2

		midScore := mid * (raceTime - mid)
		nextScore := (mid + 1) * (raceTime - (mid + 1))

		if midScore > highScore {
			low = mid
		} else {
			high = mid
		}

		if midScore == highScore || (midScore > highScore && nextScore < highScore) {
			return mid
		}
	}
}

func findLeftBound(raceTime int, highScore int) int {
	low := 0 + 1
	high := raceTime - 1

	for {
		mid := low + (high-low)/2

		midScore := mid * (raceTime - mid)
		nextScore := (mid + 1) * (raceTime - (mid + 1))

		if midScore > highScore {
			high = mid
		} else {
			low = mid
		}

		if midScore == highScore || (midScore < highScore && nextScore > highScore) {
			return mid + 1
		}
	}
}
