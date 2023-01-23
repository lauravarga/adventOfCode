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
	fmt.Println("welcome to day10 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	i := 0
	var register []int
	register = append(register, 1)
	toAdd := 0
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		i += 1
		if line == "noop" {
			register = append(register, register[len(register)-1])
			continue
		}
		register = append(register, register[len(register)-1])
		toAdd, err = strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		register = append(register, register[len(register)-1]+toAdd)
	}
	fmt.Printf("the register is: %v \n", register)

	increment := 40
	stop := 220
	current := 20
	signalStrength := 0
	sumSignalStrength := 0
	for current <= stop {
		signalStrength = register[current-1]
		sumSignalStrength += signalStrength * current
		current += increment
	}
	fmt.Printf("the sum is: %v \n", sumSignalStrength)
	line = ""
	sprite := 0
	for i := 0; i < 6; i += 1 {
		line = ""
		for j := 0; j < 40; j += 1 {
			sprite = register[j+40*i]
			if (sprite-1 == j) || (sprite == j) || (sprite+1 == j) {
				line = line + "#"
			} else {
				line = line + "."
			}
		}
		fmt.Printf("%v\n", line)
	}
}
