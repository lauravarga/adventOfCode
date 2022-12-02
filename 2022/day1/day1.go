package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("welcome 2022! day 1!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	i := 1
	cals := 0
	maxCals := 0
	calories := make(map[int]int)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			fmt.Println("new elf")
			calories[i] = cals
			i += 1
			if maxCals < cals {
				maxCals = cals
			}
			cals = 0
		} else {
			intCals, err := strconv.Atoi(txt)
			if err != nil {
				log.Fatal(err)
			}
			cals += intCals
		}
	}
	fmt.Printf("we have %v elves \n", calories)
	fmt.Printf("max calories is: %v \n", maxCals)
}
