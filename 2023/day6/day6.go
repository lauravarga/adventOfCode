package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var times []int
var distances []int

func races() int {
	totalPossibleWinningStrategies := 1
	for i := 0; i < len(times); i++ {
		winningStrategy := 0
		for time := 1; time < times[i]; time++ {

			if (times[i]-time)*time > distances[i] {
				winningStrategy++
			}
		}
		totalPossibleWinningStrategies *= winningStrategy
	}
	return totalPossibleWinningStrategies
}

func partTwo() int {
	bigTime, _ := strconv.Atoi(timesString)
	bigDistance, _ := strconv.Atoi(distancesString)
	winningStrategy := 0
	for time := 1; time < bigTime; time++ {
		if (bigTime-time)*time > bigDistance {
			winningStrategy++
		}
	}
	return winningStrategy
}

var timesString string
var distancesString string

func main() {
	fmt.Println("welcome 2023! day 6!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	readingTime := true
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%v\n", line)
		_, dataString, _ := strings.Cut(line, ":")
		listString := strings.Fields(strings.TrimSpace(dataString))
		if readingTime {
			timesString = strings.Join(listString, "")
		} else {
			distancesString = strings.Join(listString, "")
		}

		for i := 0; i < len(listString); i++ {
			n, _ := strconv.Atoi(listString[i])
			if readingTime {
				times = append(times, n)
			} else {
				distances = append(distances, n)
			}
		}
		readingTime = false
	}

	fmt.Printf("times are: %v\n", times)
	fmt.Printf("distances are: %v\n", distances)
	fmt.Printf("string time: %s \n", timesString)
	fmt.Printf("string distance: %s \n", distancesString)
	// fmt.Printf("races: %v", races())
	fmt.Printf("part two: %v", partTwo())
}
