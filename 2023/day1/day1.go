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
	fmt.Println("welcome 2023! day 1!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var firstDigit, lastDigit, sum int
	sum = 0
	numString := map[string]string{
		"one":   "on1e",
		"two":   "tw2o",
		"three": "thr3e",
		"four":  "fo4ur",
		"five":  "fi5ve",
		"six":   "si6x",
		"seven": "sev7en",
		"eight": "ei8ght",
		"nine":  "ni9ne",
	}

	for scanner.Scan() {
		line := scanner.Text()
		first := true
		fmt.Printf("original line: %s \n", line)
		replacing := true

		for replacing {
			toReplace := ""
			min := 10000
			replacing = false
			for strDigit := range numString {
				index := strings.Index(line, strDigit)

				if index != -1 && min > index {
					toReplace = strDigit
					min = index
				}
			}
			if toReplace != "" {
				line = strings.Replace(line, toReplace, numString[toReplace], 1)
				replacing = true
			}

		}
		fmt.Printf("replaced line: %s \n", line)
		for _, char := range line {
			if number, err := strconv.Atoi(string(char)); err == nil {
				if first {
					firstDigit = number
					first = false
				}
				lastDigit = number
			}
		}
		sum += firstDigit*10 + lastDigit
		fmt.Printf("found the digits: %d%d \n", firstDigit, lastDigit)
	}
	fmt.Printf("computed sum is: %d \n", sum)
}
