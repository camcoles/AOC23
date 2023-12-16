package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func findMatchingNumbers(arr1, arr2 []int, winCount bool) int {
	var points int
	var count int

	for _, num1 := range arr1 {
		for _, num2 := range arr2 {
			if num1 == num2 {
				if winCount {
					count++
				} else {
					if points != 0 {
						points = points * 2
					} else {
						points = 1
					}
				}
				break
			}
		}
	}
	if winCount {
		return count
	} else {
		return points
	}
}

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
	numberRegex := regexp.MustCompile(`\d+`)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		matches := numberRegex.FindAllString(line, -1)
		var lineNums []int
		for i := range matches {
			digit, _ := strconv.Atoi(matches[i])
			lineNums = append(lineNums, digit)
		}
		firstHalf := lineNums[1:11]
		secondHalf := lineNums[11:36]
		points := findMatchingNumbers(firstHalf, secondHalf, false)

		numbers = append(numbers, points)
	}

	var total int
	for _, points := range numbers {
		total += points
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

	numberRegex := regexp.MustCompile(`\d+`)
	freq := map[int]int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := numberRegex.FindAllString(line, -1)
		var lineNums []int
		for i := range matches {
			digit, _ := strconv.Atoi(matches[i])
			lineNums = append(lineNums, digit)
		}
		cardNum := lineNums[0]
		firstHalf := lineNums[1:11]
		secondHalf := lineNums[11:36]

		//add 1x played to each card num
		freq[cardNum] = freq[cardNum] + 1
		nextLineCount := 1
		for _, num1 := range firstHalf {
			for _, num2 := range secondHalf {
				if num1 == num2 {
					//if winning number is found add number of times played for current card to next line
					freq[cardNum+nextLineCount] += freq[cardNum]
					nextLineCount += 1
					break
				}
			}
		}
	}
	sum := 0
	for _, v := range freq {
		sum += v
	}
	fmt.Printf("sum: %v\n", sum)
}
