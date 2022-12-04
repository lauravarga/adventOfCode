package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Day03Part1(input string) {
	file, err := os.Open(input)
	Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		in := scanner.Text()

		var lookup = map[int]int{}

		mid := len(in) / 2
		for idx := 0; idx < mid; idx++ {
			lookup[int(in[idx])] += 1
		}

		for idx := mid; idx < len(in); idx++ {
			key := int(in[idx])
			_, exists := lookup[key]
			if exists {
				switch {
				case in[idx] >= 97 && in[idx] <= 122:
					sum += int(in[idx]) - 97 + 1

				case in[idx] >= 65 && in[idx] <= 90:
					sum += int(in[idx]) - 65 + 26 + 1
				}
				break
			}
		}
	}
	fmt.Println(sum)

}

func Day03Part2(input string) {
	file, err := os.Open(input)
	Check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		row1 := scanner.Text()

		scanner.Scan()
		row2 := scanner.Text()

		scanner.Scan()
		row3 := scanner.Text()

		saw := map[int32]bool{}
		for _, v1 := range row1 {

			foundIn2 := false
			foundIn3 := false

			if saw[v1] {
				continue
			}

			saw[v1] = true

			for _, v2 := range row2 {
				if v2 == v1 {
					foundIn2 = true
					break
				}
			}

			for _, v3 := range row3 {
				if v3 == v1 {
					foundIn3 = true
					break
				}
			}

			if foundIn2 && foundIn3 {
				switch {
				case v1 >= 97 && v1 <= 122:
					sum += int(v1) - 97 + 1

				case v1 >= 65 && v1 <= 90:
					sum += int(v1) - 65 + 26 + 1
				}
				break
			}
		}
	}
	fmt.Println(sum)
}

func priority(compartment1 string, compartment2 string) int {

	p := 0
	const priority = "0abcdefghijklmnopqrstuvwzyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len(compartment1); i++ {
		itemType := string(compartment1[i])
		if strings.Contains(compartment2, itemType) {
			p += strings.Index(priority, itemType)
			break
		}
	}
	return p
}

func main() {
	Day03Part1("in.txt")
	Day03Part2("in.txt")
	fmt.Println("welcome to day3 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sumPriorities := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		midpoint := len(rucksack) / 2
		// fmt.Printf("midpoint is: %v\n", midpoint)
		compartment1 := rucksack[0:midpoint]
		compartment2 := rucksack[midpoint:]
		sumPriorities += priority(compartment1, compartment2)
	}
	fmt.Printf("priorities sum is: %v \n", sumPriorities)
}
