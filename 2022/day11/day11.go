package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	startingItems []int
	operation     string
	divBy         int
	ifTrue        int
	ifFalse       int
	inspections   int
}

var denominator int

func round(monkeys []monkey) {

	for i := 0; i < len(monkeys); i += 1 {
		currentMonkey := monkeys[i]
		monkeys[i].inspections += len(monkeys[i].startingItems)
		for j := 0; j < len(currentMonkey.startingItems); j += 1 {
			ops := strings.Split(currentMonkey.operation, " ")
			var operand int
			if strings.Contains(ops[3], "old") {
				operand = currentMonkey.startingItems[j]
			} else {
				op, err := strconv.Atoi(ops[3])
				if err != nil {
					log.Fatal(err)
				}
				operand = op
			}
			var nw int
			if strings.Contains(ops[2], "+") {
				nw = currentMonkey.startingItems[j] + operand
			}
			if strings.Contains(ops[2], "*") {
				nw = currentMonkey.startingItems[j] * operand
			}
			if nw > denominator {
				nw = nw % denominator
			}
			currentMonkey.startingItems[j] = nw // 3
			d := currentMonkey.startingItems[j] % currentMonkey.divBy
			if d == 0 {
				monkeys[currentMonkey.ifTrue].startingItems = append(monkeys[currentMonkey.ifTrue].startingItems, currentMonkey.startingItems[j])
			} else {
				monkeys[currentMonkey.ifFalse].startingItems = append(monkeys[currentMonkey.ifFalse].startingItems, currentMonkey.startingItems[j])
			}

		}
		monkeys[i].startingItems = nil
	}
}

func main() {
	fmt.Println("Welcome to day 11 2022!")
	var monkeys []monkey
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	var line string
	var monkey monkey
	var numbers []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			monkeys = append(monkeys, monkey)
			monkey.startingItems = nil
			monkey.operation = ""
			monkey.divBy = 0
			monkey.ifFalse = 0
			monkey.ifTrue = 0
			monkey.inspections = 0
		}
		if strings.Contains(line, "Monkey") {
			continue
		}
		if strings.Contains(line, "Starting items:") {
			num := strings.Split(line, ":")[1]
			numbers = strings.Split(num, ", ")
			for i := 0; i < len(numbers); i += 1 {
				x, err := strconv.Atoi(strings.TrimSpace(numbers[i]))
				if err != nil {
					log.Fatal(err)
				}
				monkey.startingItems = append(monkey.startingItems, x)
			}
		}
		if strings.Contains(line, "Operation") {
			op := strings.Split(line, "=")[1]
			monkey.operation = op
		}
		if strings.Contains(line, "Test") {
			div := strings.Split(line, " ")
			monkey.divBy, err = strconv.Atoi(div[len(div)-1])
			if err != nil {
				log.Fatal(err)
			}
		}
		if strings.Contains(line, "If true") {
			iftrue := strings.Split(line, " ")
			monkey.ifTrue, err = strconv.Atoi(iftrue[len(iftrue)-1])
			if err != nil {
				log.Fatal(err)
			}
		}
		if strings.Contains(line, "If false") {
			iffalse := strings.Split(line, " ")
			monkey.ifFalse, err = strconv.Atoi(iffalse[len(iffalse)-1])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	monkeys = append(monkeys, monkey)

	denominator = 1
	for i := 0; i < len(monkeys); i += 1 {
		denominator *= monkeys[i].divBy
	}
	for i := 0; i < 10000; i += 1 {
		fmt.Printf("Round %v\n", i)
		round(monkeys)
		for j := 0; j < len(monkeys); j += 1 {
			fmt.Printf("monkey %v, inspections: %v\n", j, monkeys[j].inspections)
		}
		fmt.Println("")
	}

	max := monkeys[0].inspections
	secondMax := max
	for i := 0; i < len(monkeys); i += 1 {
		if max < monkeys[i].inspections {
			secondMax = max
			max = monkeys[i].inspections
		} else if secondMax < monkeys[i].inspections {
			secondMax = monkeys[i].inspections
		}
	}
	monkeyBusiness := max * secondMax
	fmt.Printf("Monkey business: %v\n", monkeyBusiness)
}
