package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func multiply2020(sir []int) int {
	for i := 0; i < len(sir); i++ {
		for j := 0; j < len(sir); j++ {
			if i != j {
				if sir[i]+sir[j] == 2020 {
					return sir[i] * sir[j]
				}
			}
		}
	}
	return 0
}

func multiplyThree2020(sir []int) int {
	for i := 0; i < len(sir); i++ {
		for j := 0; j < len(sir); j++ {
			if i != j {

				for z := 0; z < len(sir); z++ {
					if (z != j) && (z != i) {
						if sir[i]+sir[j]+sir[z] == 2020 {
							return sir[i] * sir[j] * sir[z]
						}
					}
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("hello")
	f, err := os.Open("in1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var report []int
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		i, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		} else {
			report = append(report, i)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("multiply2: %d \n", multiply2020(report))
	fmt.Printf("multiply3: %d \n", multiplyThree2020(report))
}
