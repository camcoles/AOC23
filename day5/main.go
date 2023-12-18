package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func expandSeeds(seeds []int, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done()

	var expandedSeeds []int

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + (seeds[i+1] - 1)

		fmt.Printf("Processing range: %d to %d\n", start, end)

		// Preallocate the expandedSeeds slice to avoid reallocations
		expandedSeeds = make([]int, 0, end-start+1)

		for j := start; j <= end; j++ {
			expandedSeeds = append(expandedSeeds, j)
		}
	}

	ch <- expandedSeeds
}

func adjustSeed(seed, soilStart, soilEnd, soilDiff int) (int, bool) {
	if seed >= soilStart && seed <= soilEnd {
		difference := (seed - soilStart)
		return soilDiff + difference, true
	}
	return seed, false
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

	var seeds []int
	var seedToSoil [][]int
	var soilToFert [][]int
	var fertToWater [][]int
	var waterToLight [][]int
	var lightToTemp [][]int
	var tempToHumid [][]int
	var humidToLocation [][]int
	numberRegex := regexp.MustCompile(`\d+`)

	var numbers [][]int
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
	seeds = numbers[0]
	seedToSoil = numbers[3:20]
	soilToFert = numbers[22:31]
	fertToWater = numbers[33:73]
	waterToLight = numbers[75:99]
	lightToTemp = numbers[101:121]
	tempToHumid = numbers[123:167]
	humidToLocation = numbers[169:210]

	var finalNums []int
	numThreads := 2

	var wg sync.WaitGroup
	ch := make(chan []int, numThreads)

	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go expandSeeds(seeds[i:i+2], &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var expandedSeeds []int
	for result := range ch {
		expandedSeeds = append(expandedSeeds, result...)
	}

	for _, seed := range expandedSeeds {
		val := seed
		for _, soil := range seedToSoil {
			soilLoc, matched := adjustSeed(seed, soil[1], soil[1]+(soil[2]-1), soil[0])
			if matched {
				val = soilLoc
				break
			}
		}
		for _, fert := range soilToFert {
			fertLoc, matched := adjustSeed(val, fert[1], fert[1]+(fert[2]-1), fert[0])
			if matched {
				val = fertLoc
				break
			}
		}
		for _, water := range fertToWater {
			waterLoc, matched := adjustSeed(val, water[1], water[1]+(water[2]-1), water[0])
			if matched {
				val = waterLoc
				break
			}
		}
		for _, light := range waterToLight {
			lightLoc, matched := adjustSeed(val, light[1], light[1]+(light[2]-1), light[0])
			if matched {
				val = lightLoc
				break
			}
		}
		for _, temp := range lightToTemp {
			tempLoc, matched := adjustSeed(val, temp[1], temp[1]+(temp[2]-1), temp[0])
			if matched {
				val = tempLoc
				break
			}
		}
		for _, humid := range tempToHumid {
			humidLoc, matched := adjustSeed(val, humid[1], humid[1]+(humid[2]-1), humid[0])
			if matched {
				val = humidLoc
				break
			}
		}
		for _, loc := range humidToLocation {
			finalLoc, matched := adjustSeed(val, loc[1], loc[1]+(loc[2]-1), loc[0])
			if matched {
				val = finalLoc
				break
			}
		}
		finalNums = append(finalNums, val)
	}
	smallest := finalNums[0]

	for _, num := range finalNums {
		if num < smallest {
			smallest = num
		}
	}
	fmt.Println(smallest)
}

//create array of int seed num using regex /d+
// create array for each category, then create int array for each line in category.
//const difference = line[1] - line[0]
//If seed num is => line[1] && seedNum =< line[1]+(line[2]-1) then o
