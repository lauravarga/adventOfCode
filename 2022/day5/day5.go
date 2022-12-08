package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stacks []string

func formStacks(line string, index int) {
	token := ""
	if len(line) >= 3 {
		if len(stacks) <= index {
			newStack := ""
			stacks = append(stacks, newStack)
			if token != "" {
				stacks[index] = token + stacks[index]
			}
		}
		if len(line) == 3 {
			token = line
			token = strings.TrimSpace(token)
			if token != "" {
				stacks[index] = token + stacks[index]
			}
		} else {
			token = line[0:4]
			token = strings.TrimSpace(token)
			if token != "" {
				stacks[index] = token + stacks[index]
			}
			index += 1
			formStacks(line[4:], index)
		}
	}
}

func move(stacks []string, line string) {

	movementInstructions := strings.Split(line, " ")
	moves, err := strconv.Atoi(movementInstructions[1])
	if err != nil {
		log.Fatal(err)
	}
	from, err := strconv.Atoi(movementInstructions[3])
	if err != nil {
		log.Fatal(err)
	}
	from = from - 1
	to, err := strconv.Atoi(movementInstructions[5])
	if err != nil {
		log.Fatal(err)
	}
	to = to - 1
	container := ""
	for i := 1; i <= moves; i += 1 {
		fromStack := stacks[from]
		container = fromStack[len(fromStack)-3:]
		stacks[to] = stacks[to] + container
		stacks[from] = stacks[from][0 : len(fromStack)-3]
	}
}

func main() {
	fmt.Println("welcome to day5 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	readingStacks := true
	for scanner.Scan() {
		line := scanner.Text()
		if readingStacks && strings.Contains(line, "[") {
			formStacks(line, 0)
			continue
		}
		if line == "" {
			readingStacks = false
			continue
		}
		if strings.Contains(line, "move") {
			move(stacks, line)
		}
	}
	result := ""
	for i := 0; i < len(stacks); i += 1 {
		container := stacks[i][len(stacks[i])-3:]
		result += string(container[1])
	}
	fmt.Printf("And the result: %v \n", result)
}
