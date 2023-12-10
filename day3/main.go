package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
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

	scheme := [][]rune{}
	sum := 0
	for fileScanner.Scan() {
		scheme = append(scheme, []rune(fileScanner.Text()))
	}

	for i := 0; i < len(scheme); i++ {
		for j := 0; j < len(scheme[i]); j++ {
			//find symbols - only '*'s are adjacent to 2 nums!
			if scheme[i][j] != '.' && !unicode.IsDigit(scheme[i][j]) {
				sum += getNums(scheme, i, j)
			}
		}

	}
	fmt.Println(sum)

}

func getNums(scheme [][]rune, i, j int) int {

	nums := make([]int, 0)
	//check nums on line above if not first line
	if i != 0 {
		nums = append(nums, getHoriNums(scheme, i-1, j)[:]...)
	}
	//check nums on current line
	nums = append(nums, getHoriNums(scheme, i, j)[:]...)

	//check nums on line below if not last line
	if i+1 < len(scheme) {
		nums = append(nums, getHoriNums(scheme, i+1, j)[:]...)
	}

	if len(nums) == 2 {
		return nums[0] * nums[1]
	}
	return 0
}

func getHoriNums(sh [][]rune, i int, j int) []int {
	ln, rn := "", ""

	//look for numbers right
	for k := j + 1; k < len(sh[i]); k++ {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		rn += string(sh[i][k])
	}
	//look for numbers left
	for k := j - 1; k >= 0; k-- {
		if !unicode.IsDigit(sh[i][k]) {
			break
		}
		ln = string(sh[i][k]) + ln
	}
	//if current location is digit concat left nums/right nums if any, convert  and return
	if unicode.IsDigit(sh[i][j]) {
		n, _ := strconv.Atoi(ln + string(sh[i][j]) + rn)
		return []int{n}
	}
	//otherwise convert and return left or right num
	num := []int{}
	lnn, _ := strconv.Atoi(ln)
	rnn, _ := strconv.Atoi(rn)
	if lnn != 0 {
		num = append(num, lnn)
	}
	if rnn != 0 {
		num = append(num, rnn)
	}
	return num
}

func hasAdjacentNumbers(line string, index int, matchDigit *regexp.Regexp) bool {
	numMatches := matchDigit.FindAllStringIndex(line, -1)
	valueOfMatches := matchDigit.FindStringSubmatch(line)
	fmt.Println(numMatches, valueOfMatches)
	for _, numMatch := range numMatches {
		numStart, numEnd := numMatch[0], numMatch[1]
		if numStart <= index && index <= numEnd {
			return true
		} else if index-1 == numEnd || index+1 == numStart {
			return true
		}
	}

	return false
}

func getPreviousLine(file *os.File, currentLineNumber int) string {
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	for i := 1; i < currentLineNumber; i++ {
		if !scanner.Scan() {
			break
		}
	}

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func getNextLine(file *os.File, currentLineNumber int) string {
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	for i := 1; i <= currentLineNumber; i++ {
		if !scanner.Scan() {
			break
		}
	}

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}
