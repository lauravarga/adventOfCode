package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("day3 2021")
	zeroes := [5]int{0, 0, 0, 0, 0}
	ones := [5]int{0, 0, 0, 0, 0}
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bline := scanner.Text()
		for pos, char := range bline {
			if string(char) == "0" {
				zeroes[pos] += 1
			} else {
				ones[pos] += 1
			}
		}
	}
	// gamma, epsilon := 0, 0
	// for i := 4; i >= 0; i-- {
	// 	if zeroes[i] > ones[i] {
	// 		epsilon += int(math.Pow(2, float64(i)))
	// 	} else {
	// 		gamma += int(math.Pow(2, float64(i)))
	// 	}
	// }
	// r := epsilon * gamma
	// fmt.Printf("result is: %v \n", r)
}
