package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("welcome 2022! day 1!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			fmt.Println("new elf")
			i += 1
		}
	}
	fmt.Printf("we have %v elves \n", i)
}
