package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/qba/Code/advent-of-code/day_2/input.txt")
	if err != nil  {
		panic(err)
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
	answer := 0

    for scanner.Scan() {
        inputLine := scanner.Text()

		gameSessionsString := strings.Split(inputLine, ": ")[1]
		gameSessionsData := strings.Split(gameSessionsString, "; ")

		gamePower := calculateGamePower(gameSessionsData)

		answer += gamePower
    }

	fmt.Println(answer)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func calculateGamePower(gameSessionsData []string) int {
	var minCubeAmounts map[string]int = map[string]int{
		"red": -1,
		"green": -1,
		"blue": -1,
	}

	for _, gameSessionData := range gameSessionsData {
		cubesData := strings.Split(gameSessionData, ", ")
		for _, cubeColorData := range cubesData {
			cubeAmountAndColor := strings.Split(cubeColorData, " ")
			cubeAmount, _ := strconv.Atoi(cubeAmountAndColor[0])
			cubeColor := cubeAmountAndColor[1]

			if minCubeAmounts[cubeColor] == -1 || minCubeAmounts[cubeColor] < cubeAmount {
				minCubeAmounts[cubeColor] = cubeAmount
			}
		}
	}

	result := 1

	for  _, cubeAmount := range minCubeAmounts {
		result *= cubeAmount
	 }

	return result
}