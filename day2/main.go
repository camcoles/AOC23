package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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

	var numbers []int
	findRed := regexp.MustCompile(`(\d+)\s*(red)`)
	findBlue := regexp.MustCompile(`(\d+)\s*(blue)`)
	findGreen := regexp.MustCompile(`(\d+)\s*(green)`)
	findGame := regexp.MustCompile(`Game (\d+)`)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		gameNo := findGame.FindAllString(line, -1)
		matchRed := findRed.FindAllString(line, -1)
		matchBlue := findBlue.FindAllString(line, -1)
		matchGreen := findGreen.FindAllString(line, -1)

		filteredGameNo := strings.Replace(gameNo[0], "Game ", "", -1)
		gameNoAsInt, _ := strconv.Atoi(filteredGameNo)

		for _, val := range matchRed {
			filteredNum := strings.Replace(val, " red", "", -1)
			digit, _ := strconv.Atoi(filteredNum)
			if digit > 12 && !slices.Contains(numbers, gameNoAsInt) {
				numbers = append(numbers, gameNoAsInt)
			}
		}
		for _, val := range matchBlue {
			filteredNum := strings.Replace(val, " blue", "", -1)
			digit, _ := strconv.Atoi(filteredNum)
			if digit > 14 && !slices.Contains(numbers, gameNoAsInt) {
				numbers = append(numbers, gameNoAsInt)
			}
		}
		for _, val := range matchGreen {
			filteredNum := strings.Replace(val, " green", "", -1)
			digit, _ := strconv.Atoi(filteredNum)
			if digit > 13 && !slices.Contains(numbers, gameNoAsInt) {
				numbers = append(numbers, gameNoAsInt)
			}
		}
	}

	total := 5050
	for _, val := range numbers {
		total -= val
	}

	fmt.Println(total)

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
	findRed := regexp.MustCompile(`(\d+)\s*(red)`)
	findBlue := regexp.MustCompile(`(\d+)\s*(blue)`)
	findGreen := regexp.MustCompile(`(\d+)\s*(green)`)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		matchRed := findRed.FindAllString(line, -1)
		matchBlue := findBlue.FindAllString(line, -1)
		matchGreen := findGreen.FindAllString(line, -1)

		minRedVal := 0
		minBlueVal := 0
		minGreenVal := 0

		for _, val := range matchRed {
			filteredNum := strings.Replace(val, " red", "", -1)
			digit, _ := strconv.Atoi(filteredNum)
			if digit > minRedVal {
				minRedVal = digit
			}
		}
		for _, val := range matchBlue {
			filteredNum := strings.Replace(val, " blue", "", -1)
			digit, _ := strconv.Atoi(filteredNum)
			if digit > minBlueVal {
				minBlueVal = digit
			}
		}
		for _, val := range matchGreen {
			filteredNum := strings.Replace(val, " green", "", -1)
			digit, _ := strconv.Atoi(filteredNum)
			if digit > minGreenVal {
				minGreenVal = digit
			}
		}
		power := minRedVal * minGreenVal * minBlueVal
		numbers = append(numbers, power)

	}

	var total int
	for _, val := range numbers {
		total += val
	}

	fmt.Println(total)
}
