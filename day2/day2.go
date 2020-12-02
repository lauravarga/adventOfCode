package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	minOccurence int
	maxOccurence int
	letter       string
	password     string
}

func validate(l line) bool {
	occ := strings.Count(l.password, l.letter)
	if (l.maxOccurence >= occ) && (l.minOccurence <= occ) {
		return true
	}
	return false
}

func validate2(l line) bool {
	occ := 0
	if string(l.password[l.minOccurence-1]) == l.letter {
		occ++
	}
	if string(l.password[l.maxOccurence-1]) == l.letter {
		occ++
	}
	if occ == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println("hello")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var lines []line
	var l line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		l.password = strings.TrimSpace(s[len(s)-1])
		l.letter = strings.TrimSpace(strings.Trim(s[1], ":"))
		l.minOccurence, err = strconv.Atoi(string(strings.Split(s[0], "-")[0]))
		if err != nil {
			log.Fatal(err)
		}
		l.maxOccurence, err = strconv.Atoi(string(strings.Split(s[0], "-")[1]))
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, l)
	}
	validPasswords := 0
	for _, v := range lines {
		if validate2(v) {
			validPasswords++
			fmt.Printf("v is valid: %v\n", v)
		}

	}
	fmt.Printf("there are %d valid passwords\n", validPasswords)
}
