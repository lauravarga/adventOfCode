package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type hill struct {
	height   int
	visited  bool
	distance int
}

func findPath(i int, j int, theHill [][]hill, count int) [][]hill {
	distance := theHill[i][j].distance + 1
	if j+1 < len(theHill[i]) {
		if !theHill[i][j+1].visited && (theHill[i][j+1].height-theHill[i][j].height <= 1) {

			if theHill[i][j+1].distance > distance {
				theHill[i][j+1].distance = distance
			}
		}
	}
	if j-1 >= 0 {
		if !theHill[i][j-1].visited && (theHill[i][j-1].height-theHill[i][j].height <= 1) {
			if theHill[i][j-1].distance > distance {
				theHill[i][j-1].distance = distance
			}
		}
	}
	if i+1 < len(theHill) {
		if !theHill[i+1][j].visited && (theHill[i+1][j].height-theHill[i][j].height <= 1) {
			if theHill[i+1][j].distance > distance {
				theHill[i+1][j].distance = distance
			}
		}
	}
	if i-1 >= 0 {
		if !theHill[i-1][j].visited && (theHill[i-1][j].height-theHill[i][j].height <= 1) {
			if theHill[i-1][j].distance > distance {
				theHill[i-1][j].distance = distance
			}
		}
	}

	theHill[i][j].visited = true
	// fmt.Printf("visited: %v, %v, %v, executions: %v\n", i, j, theHill[i][j].height, count)
	nextI, nextJ, moreToVisit := min(theHill)
	count += 1
	if moreToVisit {
		findPath(nextI, nextJ, theHill, count)
	}
	return theHill
}

func min(aHill [][]hill) (int, int, bool) {
	min := math.MaxInt
	minI, minJ := 0, 0
	moreToVisit := false
	for i := 0; i < len(aHill); i += 1 {
		for j := 0; j < len(aHill[i]); j += 1 {
			if !aHill[i][j].visited {
				moreToVisit = true
				if min >= aHill[i][j].distance {
					minI = i
					minJ = j
					min = aHill[i][j].distance
				}
			}
		}
	}
	if moreToVisit && min == math.MaxInt {
		fmt.Println("the other nodes are isolated, can't reach them. I hope we made it ot the End point")
		return minI, minJ, false
	}
	return minI, minJ, moreToVisit
}

func main() {
	var myHill [][]hill
	fmt.Println("Welcome to day12 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	iStart := 0
	jStart := 0
	iEnd := 0
	jEnd := 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		hillLine := make([]hill, len(line))
		for j := 0; j < len(line); j += 1 {
			hillLine[j].height = int(line[j])
			hillLine[j].visited = false
			hillLine[j].distance = math.MaxInt
			if line[j] == 'S' {
				iStart = i
				jStart = j
			} else if line[j] == 'E' {
				iEnd = i
				jEnd = j
			}

		}
		myHill = append(myHill, hillLine)
		i += 1
	}

	myHill[iStart][jStart].height = 96
	myHill[iStart][jStart].visited = true
	myHill[iStart][jStart].distance = 0
	myHill[iEnd][jEnd].height = 123
	finalHill := findPath(iStart, jStart, myHill, 0)
	// fmt.Println("the hill looks like this:")
	// for i := 0; i < len(finalHill); i += 1 {
	// 	for j := 0; j < len(finalHill[i]); j += 1 {
	// 		if finalHill[i][j].visited {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Print("X")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	fmt.Printf("Steps to get to the transmission point: %v\n", finalHill[iEnd][jEnd].distance)
}
