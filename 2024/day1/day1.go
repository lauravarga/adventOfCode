package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome 2024! day 1!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var firstLocationsList []int
	var secondLocationsList []int
	for scanner.Scan() {
		line := scanner.Text()
		locations := strings.Fields(line)
		// fmt.Printf("split line: %v length: %v \n", locations, len(locations))
		numericLocation1, _ := strconv.Atoi(locations[0])
		numericLocation2, _ := strconv.Atoi(locations[1])
		firstLocationsList = append(firstLocationsList, numericLocation1)
		secondLocationsList = append(secondLocationsList, numericLocation2)
	}
	// fmt.Printf("first location list: %v, second: %v \n", firstLocationsList, secondLocationsList)
	sort.Ints(firstLocationsList)
	sort.Ints(secondLocationsList)
	// fmt.Printf("SORTED first location list: %v, second: %v \n", firstLocationsList, secondLocationsList)
	sum := 0
	for i := 0; i < len(firstLocationsList); i++ {
		diff := firstLocationsList[i] - secondLocationsList[i]
		if diff >= 0 {
			sum += diff
		} else {
			sum += -diff
		}
	}
	fmt.Printf("sum of differences is: %v \n", sum)
	similaritySum := 0
	for i := 0; i < len(firstLocationsList); i++ {
		occurrence := 0
		for j := 0; j < len(firstLocationsList); j++ {
			if firstLocationsList[i] == secondLocationsList[j] {
				occurrence += 1
			}
		}
		similaritySum += firstLocationsList[i] * occurrence
	}
	fmt.Printf("similarity sum: %v \n", similaritySum)
}
