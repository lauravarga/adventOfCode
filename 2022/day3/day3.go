package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func priority(compartment1 string, compartment2 string) int {
	p := 0
	for i := 0; i < len(compartment1); i++ {
		itemType := compartment1[i]
		if strings.Contains(compartment2, string(itemType)) {
			fmt.Printf("found the problem: %v \n", string(itemType))
			if unicode.IsUpper(rune(itemType)) {
				p += int(itemType) - 38
			} else {
				p += int(itemType) - 96
			}
			break
		}

	}
	fmt.Printf("priority is: %v \n", p)
	return p
}

func main() {
	fmt.Println("welcome to day3 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sumPriorities := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		midpoint := len(rucksack) / 2
		// fmt.Printf("midpoint is: %v\n", midpoint)
		compartment1 := rucksack[0 : midpoint-1]
		compartment2 := rucksack[midpoint:]
		sumPriorities += priority(compartment1, compartment2)
	}
	fmt.Printf("priorities sum is: %v \n", sumPriorities)
}
