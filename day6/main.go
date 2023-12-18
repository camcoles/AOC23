package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var numbers [][]int
	numberRegex := regexp.MustCompile(`\d+`)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		var lineNums []int
		matches := numberRegex.FindAllString(line, -1)
		for i := range matches {
			digit, _ := strconv.Atoi(matches[i])
			lineNums = append(lineNums, digit)
		}
		numbers = append(numbers, lineNums)

	}

	var possibleWins []int

	for i, race := range numbers[0] {
		var numWaysToWin int
		for j := 1; j <= race; j++ {
			timeToMove := race - j
			distanceTravelled := timeToMove * j
			if distanceTravelled > numbers[1][i] {
				numWaysToWin++
			}
		}
		possibleWins = append(possibleWins, numWaysToWin)
	}
	marginOfError := possibleWins[0] * possibleWins[1] * possibleWins[2] * possibleWins[3]
	fmt.Println(marginOfError)
	part2()
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var numbers []int
	numberRegex := regexp.MustCompile(`\d+`)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := numberRegex.FindAllString(line, -1)

		result := ""
		result = matches[0]
		for i := 1; i < len(matches); i++ {
			result += matches[i]
		}
		digit, _ := strconv.Atoi(result)
		numbers = append(numbers, digit)
	}

	var numWaysToWin int
	race := numbers[0]
	record := numbers[1]
	for j := 1; j <= race; j++ {
		timeToMove := race - j
		distanceTravelled := timeToMove * j
		if distanceTravelled > record {
			numWaysToWin++
		}
	}

	fmt.Println(numWaysToWin)
}
