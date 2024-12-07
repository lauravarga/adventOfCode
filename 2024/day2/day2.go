package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome 2024! day 1!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var reports [][]int
	var reportInt []int
	for scanner.Scan() {
		reportString := scanner.Text()
		reportStringList := strings.Split(reportString, " ")
		reportInt = nil
		for i := 0; i < len(reportStringList); i++ {
			num, _ := strconv.Atoi(reportStringList[i])
			reportInt = append(reportInt, num)
		}
		reports = append(reports, reportInt)
	}
	fmt.Printf("reports: %v \n", reports)

	safe := 0
	for i := 0; i < len(reports); i++ {
		isSafe := true
		increase := false
		dampner := 1
		if reports[i][0]-reports[i][1] <= 0 {
			increase = true
		}
		for j := 0; j < len(reports[i])-1; j++ {
			diff := reports[i][j] - reports[i][j+1]
			if diff == 0 {
				isSafe = false
				dampner -= 1
			} else {
				if (diff < 0) && !increase {
					isSafe = false
					dampner -= 1
				}
				if (diff > 0) && increase {
					isSafe = false
					dampner -= 1
				}
				if (diff > 3) || (diff < -3) {
					isSafe = false
					dampner -= 1
				}

			}

		}
		if isSafe {
			safe += 1
		} else {
			if dampner >= 0 {
				safe += 1
			}
		}

	}

	fmt.Printf("safe reports: %v \n", safe)
}
