package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func hasDuplicates(datadatastream string) bool {
	for i := 0; i < len(datadatastream); i += 1 {
		if strings.Count(datadatastream, string(datadatastream[i])) > 1 {
			return true
		}
	}
	return false
}

func findMarker(datastream string) int {
	// fmt.Printf("finding marker in %v \n", datastream)
	for i := 0; i < len(datastream)-14; i += 1 {
		end := i + 14
		if !hasDuplicates(datastream[i:end]) {
			return end
		}
	}
	return 0
}

func main() {
	fmt.Println("welcome to day 6 2022!")
	f, err := os.Open("in.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		datastream := scanner.Text()
		result := findMarker(datastream)
		fmt.Printf("found marker at char: %v\n", result)
	}

}
