package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type scratchcardType struct {
	cardId  int
	matches int
	copies  int
}

var allScratchcards []scratchcardType

func scratch() {
	fmt.Printf("Processing scratchcards \n")
	for i := 0; i < len(allScratchcards); i++ {
		fmt.Printf("processing: %v\n", allScratchcards[i])
		for k := i + 1; k <= i+allScratchcards[i].matches; k++ {
			allScratchcards[k].copies += allScratchcards[i].copies
		}
	}
	sum := 0
	for i := 0; i < len(allScratchcards); i++ {
		sum += allScratchcards[i].copies
	}
	fmt.Printf("this many scratchcards: %d\n", sum)
}

func main() {
	fmt.Println("welcome 2023! day 4!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%v\n", line)
		cardIDString, numbersStrings, _ := strings.Cut(line, ":")
		var sc scratchcardType
		_, tempCardID, _ := strings.Cut(cardIDString, " ")
		tempCardID = strings.TrimSpace(tempCardID)
		sc.cardId, _ = strconv.Atoi(tempCardID)
		scratchcard := strings.Split(numbersStrings, "|")

		scratchcardWinning := strings.Split(strings.TrimSpace(scratchcard[0]), " ")
		myScratchcardNumbers := strings.Split(strings.TrimSpace(scratchcard[1]), " ")
		// fmt.Printf("scratchcardwinning strings are: %v, mynumbers: %v\n", scratchcardWinning, myScratchcardNumbers)
		power := 0
		first := true
		matches := 0
		for i := 0; i < len(myScratchcardNumbers); i++ {

			for j := 0; j < len(scratchcardWinning); j++ {
				if myScratchcardNumbers[i] != "" && myScratchcardNumbers[i] == scratchcardWinning[j] {
					if first {
						power = 1
						first = false
						matches++
					} else {
						power *= 2
						matches++
					}
				}
			}

		}
		sc.matches = matches
		// fmt.Printf("scratchcards value is: %d\n", power)
		sc.copies = 1
		sum += power
		allScratchcards = append(allScratchcards, sc)
	}
	fmt.Printf("scratchcards sum is: %d\n", sum)
	scratch()
}
