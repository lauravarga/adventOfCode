package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sliding3(depth []int) int {
	d13 := 0
	d23 := 0
	inc := 0
	for i := 0; i < len(depth)-2; i++ {
		d23 = depth[i] + depth[i+1] + depth[i+2]
		if d23 > d13 && d13 != 0 {
			inc++
		}
		d13 = d23
	}
	return inc
}

func main() {
	fmt.Println("day1 2021")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	d1 := 0
	d2 := 0
	inc := 0
	var depths []int
	for scanner.Scan() {
		str := scanner.Text()
		d2, err = strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, d2)
		if d2 > d1 && d1 != 0 {
			inc++
		}
		d1 = d2
	}
	fmt.Printf("Depth increase: %v \n", inc)
	fmt.Printf("Depth increase sliding3: %v \n", sliding3(depths))
}
