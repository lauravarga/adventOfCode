package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	step      int
}

var instructions []instruction

func move(instructions []instruction) int {
	h := 0
	d := 0
	for i := 0; i < len(instructions); i++ {
		switch instructions[i].direction {
		case "forward":
			d += instructions[i].step
		case "down":
			h += instructions[i].step
		case "up":
			h -= instructions[i].step
		}
	}
	position := h * d
	return position
}

func moveWithAim(instructions []instruction) int {
	h := 0
	d := 0
	aim := 0
	for i := 0; i < len(instructions); i++ {
		switch instructions[i].direction {
		case "forward":
			h += instructions[i].step
			d += instructions[i].step * aim
		case "down":
			// d += instructions[i].step
			aim += instructions[i].step
		case "up":
			// d -= instructions[i].step
			aim -= instructions[i].step
		}
	}
	position := h * d
	return position
}

func main() {
	fmt.Println("day2 2021")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		s := strings.Split(str, " ")
		step, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		inst := instruction{
			direction: s[0],
			step:      step,
		}
		instructions = append(instructions, inst)
	}
	fmt.Printf("position is: %v \n", move(instructions))
	fmt.Printf("position of mving with aim is: %v \n", moveWithAim(instructions))
}
