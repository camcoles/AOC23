package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// "regexp"
	// "strconv"
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

	freq := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
	fmt.Println(freq)

	// var numbers []int
	handRegex := regexp.MustCompile(`A|K|Q|J|T|\d`)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		handMatch := handRegex.FindAllString(line, -1)
		hand := 0
		for _, match := range handMatch {
			hand = hand*10 + freq[match]
		}

		fmt.Println(hand, handMatch)

		// 	result := ""
		// 	result = matches[0]
		// 	for i := 1; i < len(matches); i++ {
		// 		result += matches[i]
		// 	}
		// 	digit, _ := strconv.Atoi(result)
		// 	numbers = append(numbers, digit)
	}
}
