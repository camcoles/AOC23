package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var numbers []int
	numberRegex := regexp.MustCompile(`\d`)

	fileScanner := bufio.NewScanner(readFile)
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		matches := numberRegex.FindAllString(line, -1)

		if len(matches) == 1 {
			digit, _ := strconv.Atoi(matches[0])
			concatenatedNumber := digit*10 + digit
			numbers = append(numbers, concatenatedNumber)
		} else {
			firstDigit, _ := strconv.Atoi(matches[0])
			lastDigit, _ := strconv.Atoi(matches[len(matches)-1])

			concatenatedNumber := firstDigit*10 + lastDigit
			numbers = append(numbers, concatenatedNumber)
		}
	}

	var total int
	for _, val := range numbers {
		total += val
	}

	fmt.Println("Part 1:", total)

	part2()
}

type StringMap map[string]int

func getValueByKey(key string, m StringMap) (int, bool) {
	value, ok := m[key]
	return value, ok
}

func part2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	strNumVals := StringMap{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var numbers []int
	numberRegex := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)

	fileScanner := bufio.NewScanner(readFile)
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	findReplacePairs := map[string]string{
		"oneight":   "18",
		"nineight":  "98",
		"eighthree": "83",
		"fiveight":  "58",
		"eightwo":   "82",
		"twone":     "21",
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		for find, replace := range findReplacePairs {
			line = strings.Replace(line, find, replace, -1)
		}

		matches := numberRegex.FindAllString(line, -1)

		if len(matches) == 1 {
			digit, err := strconv.Atoi(matches[0])
			concatenatedNumber := digit*10 + digit
			numbers = append(numbers, concatenatedNumber)
			if err != nil {
				panic("error converting string to int")
			}
		} else {
			var firstNum int
			var secondNum int
			firstDigit, err := strconv.Atoi(matches[0])
			lastDigit, err2 := strconv.Atoi(matches[len(matches)-1])

			if err != nil {
				convertedFirstValue, _ := getValueByKey(matches[0], strNumVals)
				firstNum = convertedFirstValue
			} else {
				firstNum = firstDigit
			}
			if err2 != nil {
				convertedSecondValue, _ := getValueByKey(matches[len(matches)-1], strNumVals)
				secondNum = convertedSecondValue
			} else {
				secondNum = lastDigit
			}
			concatenatedNumber := firstNum*10 + secondNum
			numbers = append(numbers, concatenatedNumber)
		}
	}

	var total int
	for _, val := range numbers {
		total += val
	}

	fmt.Println("Part 2:", total)
}
