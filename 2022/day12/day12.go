// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math"
// 	"os"
// )

// type hill struct {
// 	height   int
// 	visited  bool
// 	distance int
// }

// var minMoves int

// func findPath(i int, j int, steps int, theHill [][]hill) [][]hill {
// 	if theHill[i][j].distance > steps {
// 		theHill[i][j].distance = steps
// 	}
// 	theHill[i][j].visited = true
// 	if j+1 < len(theHill[i]) {
// 		if !theHill[i][j+1].visited && ((theHill[i][j+1].height-theHill[i][j].height == 1) || (theHill[i][j+1].height-theHill[i][j].height == 0)) {
// 			findPath(i, j+1, steps+1, theHill)
// 		}
// 	}
// 	if i+1 < len(theHill) {
// 		if !theHill[i+1][j].visited && ((theHill[i+1][j].height-theHill[i][j].height == 1) || (theHill[i+1][j].height-theHill[i][j].height == 0)) {
// 			findPath(i+1, j, steps+1, theHill)
// 		}
// 	}
// 	if i-1 >= 0 {
// 		if !theHill[i-1][j].visited && ((theHill[i-1][j].height-theHill[i][j].height == 1) || (theHill[i-1][j].height-theHill[i][j].height == 0)) {
// 			findPath(i-1, j, steps+1, theHill)
// 		}
// 	}
// 	if j-1 >= 0 {
// 		if !theHill[i][j-1].visited && ((theHill[i][j-1].height-theHill[i][j].height == 1) || (theHill[i][j-1].height-theHill[i][j].height == 0)) {
// 			findPath(i, j-1, steps+1, theHill)
// 		}
// 	}
// 	// theHill[i][j].visited = true
// 	return theHill
// }

// func main() {
// 	var myHill [][]hill
// 	fmt.Println("Welcome to day12 2022!")
// 	f, err := os.Open("in.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(f)
// 	iStart := 0
// 	jStart := 0
// 	iEnd := 0
// 	jEnd := 0
// 	i := 0
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		hillLine := make([]hill, len(line))
// 		for j := 0; j < len(line); j += 1 {
// 			hillLine[j].height = int(line[j])
// 			hillLine[j].visited = false
// 			hillLine[j].distance = math.MaxInt
// 			if line[j] == 'S' {
// 				iStart = i
// 				jStart = j
// 			} else if line[j] == 'E' {
// 				iEnd = i
// 				jEnd = j
// 			}

// 		}
// 		myHill = append(myHill, hillLine)
// 		i += 1
// 	}

// 	myHill[iStart][jStart].height = 96
// 	myHill[iStart][jStart].visited = true
// 	myHill[iEnd][jEnd].height = 123
// 	minMoves = math.MaxInt
// 	finalHill := findPath(iStart, jStart, 0, myHill)
// 	fmt.Println("the hill looks like this:")
// 	for i := 0; i < len(finalHill); i += 1 {
// 		for j := 0; j < len(finalHill[i]); j += 1 {
// 			if finalHill[i][j].distance == 9223372036854775807 {
// 				fmt.Print("X ")
// 			} else {
// 				// fmt.Printf("%v ", finalHill[i][j].distance)
// 				fmt.Print("O ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Printf("Min moves: %v\n", minMoves)
// 	fmt.Printf("Steps to get to the transmission point: %v\n", finalHill[iEnd][jEnd].distance)
// }
