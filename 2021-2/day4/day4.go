package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var allGrids [][5]map[int]bool

func mark(extr int) {
	for i := 0; i < 5; i++ {
		for k, _ := range grid[i] {
			if k == extr {
				grid[i][k] = true
			}
		}
	}
}

func main() {
	fmt.Println("welcome to day4!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	extrNumbers := make([]int, 0)
	if scanner.Scan() {
		extraction := strings.Split(scanner.Text(), ",")
		fmt.Printf("extractions: %v \n", extraction)
		n := len(extraction)
		for i := 0; i < n; i++ {
			num, err := strconv.Atoi(extraction[i])
			if err != nil {
				log.Fatal(err)
			}
			extrNumbers = append(extrNumbers, num)
		}
	}
	var bingoGrid [5]map[int]bool
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fmt.Printf("full line: %v \n", line)
			xLine := strings.Fields(line)
			bingoGrid[i] = make(map[int]bool, 5)
			for j := 0; j < len(xLine); j++ {
				num, err := strconv.Atoi(xLine[j])
				if err != nil {
					log.Fatal(err)
				}
				bingoGrid[i][num] = false
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
	fmt.Println("let's play!")
	// extracting number
	for i := 0; i < len(extrNumbers); i++ {
		//marking each grid & searching for winner (iterating through the grids)
		for j := 0; j < len(allGrids); j++ {
			if mark(allGrids[j], extrNumbers(i)) {
				fmt.Printf("grid %v wins!", j)
			}
		}
	}
}
