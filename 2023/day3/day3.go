package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var engine [][]string

func checkNeighbour(i int, j int) int {
	enginePart := 0
	if unicode.IsDigit([]rune(engine[i][j])[0]) {
		// fmt.Printf("found a digit: %s\n", engine[i][j])
		stop := false
		stringNumber := ""
		for m := j; m >= 0 && !stop; m-- {
			if unicode.IsDigit([]rune(engine[i][m])[0]) {
				stringNumber = engine[i][m] + stringNumber
				engine[i][m] = "."
			} else {
				stop = true
			}
		}
		stop = false
		for m := j + 1; m < len(engine[i]) && !stop; m++ {
			if unicode.IsDigit([]rune(engine[i][m])[0]) {
				stringNumber = stringNumber + engine[i][m]
				engine[i][m] = "."
			} else {
				stop = true
			}
		}
		fmt.Printf("this looks like an engine part: %s\n", stringNumber)
		num, _ := strconv.Atoi(stringNumber)
		enginePart += num
	}
	return enginePart
}

func sumOfParts() {
	theSum := 0
	for i := 0; i < len(engine); i++ {
		for j := 0; j < len(engine[i]); j++ {
			// fmt.Printf("char is: %s\n", engine[i][j])
			if engine[i][j] != "." {
				_, err := strconv.Atoi(engine[i][j])
				if err != nil {
					fmt.Printf("found a symbol: %s, coordinates are i: %d, j: %d\n", engine[i][j], i, j)
					if i-1 >= 0 {
						theSum += checkNeighbour(i-1, j)
						if j+1 < len(engine[i]) {
							theSum += checkNeighbour(i-1, j+1)
						}
						if j-1 > 0 {
							theSum += checkNeighbour(i-1, j-1)
						}
					}
					if j-1 > 0 {
						theSum += checkNeighbour(i, j-1)
					}
					if j+1 < len(engine[i]) {
						theSum += checkNeighbour(i, j+1)
					}

					if i+1 < len(engine) {
						theSum += checkNeighbour(i+1, j)
						if j+1 < len(engine[i]) {
							theSum += checkNeighbour(i+1, j+1)
						}
						if j-1 > 0 {
							theSum += checkNeighbour(i+1, j-1)
						}

					}

				}
			}
		}
	}
	fmt.Printf("the sum of the parts is: %d\n", theSum)
}

func gearsRatio() {
	var sum int64
	sum = 0
	for i := 0; i < len(engine); i++ {
		for j := 0; j < len(engine[i]); j++ {

			// fmt.Printf("char is: %s\n", engine[i][j])
			var neighbours []int
			if engine[i][j] == "*" {
				fmt.Printf("found a possible gear: %s, coordinates are i: %d, j: %d\n", engine[i][j], i, j)
				neighbours = nil
				var tempNeighbour int
				if i-1 >= 0 {
					tempNeighbour = checkNeighbour(i-1, j)
					if tempNeighbour != 0 {
						neighbours = append(neighbours, tempNeighbour)
					}
					if j+1 < len(engine[i]) {
						tempNeighbour = checkNeighbour(i-1, j+1)
						if tempNeighbour != 0 {
							neighbours = append(neighbours, tempNeighbour)
						}
					}
					if j-1 > 0 {
						tempNeighbour = checkNeighbour(i-1, j-1)
						if tempNeighbour != 0 {
							neighbours = append(neighbours, tempNeighbour)
						}
					}
				}
				if j-1 > 0 {
					tempNeighbour = checkNeighbour(i, j-1)
					if tempNeighbour != 0 {
						neighbours = append(neighbours, tempNeighbour)
					}
				}
				if j+1 < len(engine[i]) {
					tempNeighbour = checkNeighbour(i, j+1)
					if tempNeighbour != 0 {
						neighbours = append(neighbours, tempNeighbour)
					}
				}
				if i+1 < len(engine) {
					tempNeighbour = checkNeighbour(i+1, j)
					if tempNeighbour != 0 {
						neighbours = append(neighbours, tempNeighbour)
					}
					if j+1 < len(engine[i]) {
						tempNeighbour = checkNeighbour(i+1, j+1)
						if tempNeighbour != 0 {
							neighbours = append(neighbours, tempNeighbour)
						}
					}
					if j-1 > 0 {
						tempNeighbour = checkNeighbour(i+1, j-1)
						if tempNeighbour != 0 {
							neighbours = append(neighbours, tempNeighbour)
						}
					}

				}
			}
			if len(neighbours) == 2 {
				sum += int64(neighbours[0] * neighbours[1])
			}
		}

	}
	fmt.Printf("sum of power for the gears: %d\n", sum)
}

func main() {
	fmt.Println("welcome 2023! day 3!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	length := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line: %v\n", line)
		var charLine []string
		for i := 0; i < len(line); i++ {
			fmt.Printf("char is: %v\n", string(line[i]))
			charLine = append(charLine, string(line[i]))
		}
		engine = append(engine, charLine)
		length += 1
	}

	// sumOfParts()
	gearsRatio()

}
