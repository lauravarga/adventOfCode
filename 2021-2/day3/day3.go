package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func comps(readings []string, cursor int) (int, int) {
	zs := 0
	ons := 0
	for i := 0; i < len(readings); i++ {
		if string(readings[i][cursor]) == "0" {
			zs++
		} else {
			ons++
		}
	}
	return zs, ons
}

func o2gen(readings []string, cursor int) string {
	fmt.Printf("recursive: %v \n", len(readings))
	if len(readings) == 1 {
		return readings[0]
	}
	if cursor > len(readings[0]) {
		return ""
	}
	numbers := make([]string, 0)
	zs, ons := comps(readings, cursor)
	if zs > ons {
		for i := 0; i < len(readings); i++ {
			if string(readings[i][cursor]) == "0" {
				numbers = append(numbers, readings[i])
			}
		}
		cursor += 1
		return o2gen(numbers, cursor)
	} else {
		for i := 0; i < len(readings); i++ {
			if string(readings[i][cursor]) == "1" {
				numbers = append(numbers, readings[i])
			}
		}
		cursor += 1
		return o2gen(numbers, cursor)
	}
}

func co2gen(readings []string, cursor int) string {
	fmt.Printf("recursive: %v \n", len(readings))
	if len(readings) == 1 {
		return readings[0]
	}
	if cursor > len(readings[0]) {
		return ""
	}
	numbers := make([]string, 0)
	zs, ons := comps(readings, cursor)
	if zs <= ons {
		for i := 0; i < len(readings); i++ {
			if string(readings[i][cursor]) == "0" {
				numbers = append(numbers, readings[i])
			}
		}
		cursor += 1
		return co2gen(numbers, cursor)
	} else {
		for i := 0; i < len(readings); i++ {
			if string(readings[i][cursor]) == "1" {
				numbers = append(numbers, readings[i])
			}
		}
		cursor += 1
		return co2gen(numbers, cursor)
	}
}

func main() {
	fmt.Println("day3 2021")
	// TODO: rework to not have static array
	zeroes := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ones := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var readings []string
	for scanner.Scan() {
		bline := scanner.Text()
		readings = append(readings, bline)
		for pos, char := range bline {
			if string(char) == "0" {
				zeroes[pos] += 1
			} else {
				ones[pos] += 1
			}
		}
	}
	gamma, epsilon := 0, 0
	for i := len(zeroes) - 1; i >= 0; i-- {
		if zeroes[i] > ones[i] {
			epsilon += int(math.Pow(2, float64(11-i)))
		} else {
			gamma += int(math.Pow(2, float64(11-i)))
		}
	}
	r := epsilon * gamma
	fmt.Printf("power consumption is: %v \n", r)
	o2g := o2gen(readings, 0)
	co2g := co2gen(readings, 0)
	o2gdec, err := strconv.ParseInt(o2g, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	co2gdec, err := strconv.ParseInt(co2g, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("o2 generator is: %v \n", o2gdec)
	fmt.Printf("Co2 generator is: %v \n", co2gdec)
	fmt.Printf("co2 scrubber value: %v \n", o2gdec*co2gdec)

}
