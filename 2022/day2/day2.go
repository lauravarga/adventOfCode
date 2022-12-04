package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type rockPaperScissors struct {
	opponent string
	advice   string
	points   int
}

func main() {
	var theGuide []rockPaperScissors
	var instruction rockPaperScissors

	fmt.Println("welcome to day 2 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, "")
		instruction.opponent = moves[0]
		instruction.advice = moves[2]
		instruction.points = 0
		if instruction.advice == "X" {
			instruction.points += 1
			switch instruction.opponent {
			case "A":
				instruction.points += 3
			case "B":
				instruction.points += 0
			case "C":
				instruction.points += 6
			}
		}
		if instruction.advice == "Y" {
			instruction.points += 2
			switch instruction.opponent {
			case "A":
				instruction.points += 6
			case "B":
				instruction.points += 3
			case "C":
				instruction.points += 0
			}
		}
		if instruction.advice == "Z" {
			instruction.points += 3
			switch instruction.opponent {
			case "A":
				instruction.points += 0
			case "B":
				instruction.points += 6
			case "C":
				instruction.points += 3
			}
		}
		sum += instruction.points
		theGuide = append(theGuide, instruction)
	}
	fmt.Printf("the result is: %v\n", sum)
}
