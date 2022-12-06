package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func twoWaySplitInt(line string, separator string) (int, int) {
	parts := strings.Split(line, separator)
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return start, end
}

func overlap(line string) bool {
	sectionsOverlap := false
	parts := strings.Split(line, ",")
	firstSection := parts[0]
	secondSection := parts[1]
	firstSectionStart, firstSectionEnd := twoWaySplitInt(firstSection, "-")
	secondSectionStart, secondSectionEnd := twoWaySplitInt(secondSection, "-")
	if (firstSectionStart <= secondSectionStart) && (firstSectionEnd >= secondSectionEnd) {
		sectionsOverlap = true
	}
	if secondSectionStart <= firstSectionStart && secondSectionEnd >= firstSectionEnd {
		sectionsOverlap = true
	}
	return sectionsOverlap
}

func overlapAtAll(line string) bool {
	sectionsOverlap := false
	parts := strings.Split(line, ",")
	firstSection := parts[0]
	secondSection := parts[1]
	firstSectionStart, firstSectionEnd := twoWaySplitInt(firstSection, "-")
	secondSectionStart, secondSectionEnd := twoWaySplitInt(secondSection, "-")
	if firstSectionEnd >= secondSectionStart && firstSectionStart <= secondSectionEnd {
		sectionsOverlap = true
	}

	return sectionsOverlap
}

func main() {
	fmt.Println("welcome to day 4!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	contains := 0
	anyOverlap := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if overlap(line) {
			contains += 1
		}
		if overlapAtAll(line) {
			anyOverlap += 1
		}
	}
	fmt.Printf("number of assinments that completely overlap: %v\n", contains)
	fmt.Printf("number of assinments that partially overlap: %v\n", anyOverlap)
}
