package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isDigit(s string) bool {
	match, err := regexp.MatchString(`\d`, s)
	return err == nil && match
}

func isSymbol(s string) bool {
	match, err := regexp.MatchString(`[*@=$#+%&/-]`, s)
	return err == nil && match
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

	var numbers [][]string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		var charArray []string
		for _, char := range line {
			charArray = append(charArray, string(char))
		}
		numbers = append(numbers, charArray)
	}

	var partNumsPerLine []int
	for i, arr := range numbers {
		var prevTotal int
		for subI, val := range arr {
			digit, err := strconv.Atoi(val)
			var isPartNum bool = false

			if i == 0 && err == nil {
				if subI > 0 && isSymbol(arr[subI-1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(arr[subI+1]) {
					isPartNum = true
				} else if isSymbol(numbers[i+1][subI]) {
					isPartNum = true
				} else if subI > 0 && isSymbol(numbers[i+1][subI-1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(numbers[i+1][subI+1]) {
					isPartNum = true
				}
			} else if i != 139 && err == nil {
				if subI > 0 && isSymbol(arr[subI-1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(arr[subI+1]) {
					isPartNum = true
				} else if isSymbol(numbers[i-1][subI]) {
					isPartNum = true
				} else if isSymbol(numbers[i+1][subI]) {
					isPartNum = true
				} else if subI > 0 && isSymbol(numbers[i-1][subI-1]) {
					isPartNum = true
				} else if subI > 0 && isSymbol(numbers[i+1][subI-1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(numbers[i-1][subI+1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(numbers[i+1][subI+1]) {
					isPartNum = true
				}
			} else if i == 139 && err == nil {
				if subI > 0 && isSymbol(arr[subI-1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(arr[subI+1]) {
					isPartNum = true
				} else if isSymbol(numbers[i-1][subI]) {
					isPartNum = true
				} else if subI > 0 && isSymbol(numbers[i-1][subI-1]) {
					isPartNum = true
				} else if subI != 139 && isSymbol(numbers[i-1][subI+1]) {
					isPartNum = true
				}
			}
			var total int

			if isPartNum == true {
				if !isDigit(arr[subI-1]) {
					total = digit
					if isDigit(arr[subI+1]) {
						rightDigit, _ := strconv.Atoi(arr[subI+1])
						total = total*10 + rightDigit
						if isDigit(arr[subI+2]) {
							secondRightDigit, _ := strconv.Atoi(arr[subI+2])
							total = total*10 + secondRightDigit
						}
					}
				} else if !isDigit(arr[subI+1]) {
					total = digit
					if isDigit(arr[subI-1]) {
						leftDigit, _ := strconv.Atoi(arr[subI-1])
						total = leftDigit*10 + total
						if isDigit(arr[subI-2]) {
							secondLeftDigit, _ := strconv.Atoi(arr[subI-2])
							total = secondLeftDigit*100 + total
						}
					}
				} else if isDigit(arr[subI-1]) && isDigit(arr[subI+1]) {
					leftDigit, _ := strconv.Atoi(arr[subI-1])
					rightDigit, _ := strconv.Atoi(arr[subI+1])
					var tempVal int
					tempVal = leftDigit*10 + digit
					total = tempVal*10 + rightDigit
				}
			}
			if prevTotal != total {
				partNumsPerLine = append(partNumsPerLine, total)
			}
			prevTotal = total
		}
	}
	var total int
	for _, val := range partNumsPerLine {
		total += val
	}
	fmt.Println(total)
}
