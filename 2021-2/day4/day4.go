package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var allGrids [][5][5]map[int]bool

// func lineToBLine(line string) map[int]bool {
// 	bLine := make(map[int]bool, 5)
// 	xLine := strings.Split(line, ",")
// 	for i := 0; i < len(xLine); i++ {
// 		num, err := strconv.Atoi(xLine[i])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		bLine[num] = false
// 	}
// 	return bLine
// }

func main() {
	fmt.Println("welcome to day4!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		extraction := strings.Split(scanner.Text(), ",")
		fmt.Printf("extractions: %v \n", extraction)
		n := len(extraction)
		// var extrNumbers []int
		extrNumbers := make([]int, 0)
		for i := 0; i < n; i++ {
			num, err := strconv.Atoi(extraction[i])
			if err != nil {
				log.Fatal(err)
			}
			extrNumbers = append(extrNumbers, num)
		}
	}
	var bingoGrid [5][5]map[int]bool
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fmt.Printf("full line: %v \n", line)
			// bLine := lineToBLine(line)
			// bingoGrid[i] = append(bingoGrid[i], bLine)
			xLine := strings.Fields(line)
			for j := 0; j < len(xLine); j++ {
				num, err := strconv.Atoi(xLine[j])
				if err != nil {
					log.Fatal(err)
				}
				bingoGrid[i][j] = make(map[int]bool, 0)
				bingoGrid[i][j][num] = false
			}
			i++
		} else {
			fmt.Println("EMPTY LINE")
			if i != 0 {
				allGrids = append(allGrids, bingoGrid)
			}
			i = 0

		}
	}
	allGrids = append(allGrids, bingoGrid)
	fmt.Println("it's done!")
}
