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

func option1(opponent string, advice string) int {
	points := 0
	if advice == "X" {
		points += 1
		switch opponent {
		case "A":
			points += 3
		case "B":
			points += 0
		case "C":
			points += 6
		}
	}
	if advice == "Y" {
		points += 2
		switch opponent {
		case "A":
			points += 6
		case "B":
			points += 3
		case "C":
			points += 0
		}
	}
	if advice == "Z" {
		points += 3
		switch opponent {
		case "A":
			points += 0
		case "B":
			points += 6
		case "C":
			points += 3
		}
	}
	return points
}

func option2(opponent string, advice string) int {
	points := 0
	if advice == "X" {
		points += 0
		switch opponent {
		case "A":
			points += 3
		case "B":
			points += 1
		case "C":
			points += 2
		}
	}
	if advice == "Y" {
		points += 3
		switch opponent {
		case "A":
			points += 1
		case "B":
			points += 2
		case "C":
			points += 3
		}
	}
	if advice == "Z" {
		points += 6
		switch opponent {
		case "A":
			points += 2
		case "B":
			points += 3
		case "C":
			points += 1
		}
	}
	return points
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
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, "")
		instruction.opponent = moves[0]
		instruction.advice = moves[2]
		instruction.points = option1(instruction.opponent, instruction.advice)

		sum1 += instruction.points
		sum2 += option2(instruction.opponent, instruction.advice)
		theGuide = append(theGuide, instruction)
	}
	fmt.Printf("the first result is: %v\n", sum1)
	fmt.Printf("the second result is: %v\n", sum2)
}
