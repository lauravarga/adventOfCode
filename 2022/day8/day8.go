package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isVisible(i int, j int, forrest [][]int) bool {
	visible := true
	for x := 0; x < i; x += 1 {
		if forrest[x][j] >= forrest[i][j] {
			visible = false

			break
		}
	}
	if !visible {
		visible = true
		for x := i + 1; x < len(forrest[i]); x += 1 {
			if forrest[x][j] >= forrest[i][j] {
				visible = false

				break
			}
		}
	}
	if !visible {
		visible = true
		for y := 0; y < j; y += 1 {
			if forrest[i][y] >= forrest[i][j] {
				visible = false

				break
			}
		}
	}
	if !visible {
		visible = true
		for y := j + 1; y < len(forrest[i]); y += 1 {
			if forrest[i][y] >= forrest[i][j] {
				visible = false

				break
			}
		}
	}

	return visible
}

func visibleTrees(forrest [][]int, count int) int {
	for i := 0; i < len(forrest); i += 1 {
		for j := 0; j < len(forrest[i]); j += 1 {
			if (i == 0) || (i == len(forrest)-1) {
				count += 1
			} else if (j == 0) || (j == len(forrest[i])-1) {
				count += 1
			} else if isVisible(i, j, forrest) {
				count += 1
			}

		}
	}
	return count
}

func scenicScore(forrest [][]int, i int, j int) int {
	score := 1
	count := 0
	for x := i - 1; x >= 0; x -= 1 {
		if forrest[x][j] >= forrest[i][j] {
			count += 1
			break
		}
		count += 1
	}
	score = score * count
	count = 0
	for x := i + 1; x < len(forrest[i]); x += 1 {
		if forrest[x][j] >= forrest[i][j] {
			count += 1
			break
		}
		count += 1
	}
	score = score * count
	count = 0
	for y := j - 1; y >= 0; y -= 1 {
		if forrest[i][y] >= forrest[i][j] {
			count += 1
			break
		}
		count += 1
	}
	score = score * count
	count = 0
	for y := j + 1; y < len(forrest[i]); y += 1 {
		if forrest[i][y] >= forrest[i][j] {
			count += 1
			break
		}
		count += 1
	}
	score = score * count
	return score
}

func maxScenicScore(forrest [][]int) int {
	max := 0
	sscenicScore := 0
	for i := 0; i < len(forrest); i += 1 {
		for j := 0; j < len(forrest[i]); j += 1 {
			if (i == 0) || (i == len(forrest)-1) {
				sscenicScore = 0
			} else if (j == 0) || (j == len(forrest[i])-1) {
				sscenicScore = 0
			} else {
				sscenicScore = scenicScore(forrest, i, j)
			}

			if max < sscenicScore {
				max = sscenicScore
			}
		}
	}
	return max
}

func main() {
	fmt.Println("welcome to day8 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var forrest [][]int
	forrest = make([][]int, 0)

	i := 0
	for scanner.Scan() {
		row := make([]int, 0)
		forrest = append(forrest, row)
		line := scanner.Text()
		for j := 0; j < len(line); j++ {
			tree, err := strconv.Atoi(string(line[j]))
			if err != nil {
				log.Fatal(err)
			}
			forrest[i] = append(forrest[i], tree)
		}
		i += 1
	}
	visibleTrees := visibleTrees(forrest, 0)
	fmt.Printf("visible tree count: %v\n", visibleTrees)
	maxScore := maxScenicScore(forrest)
	fmt.Printf("max scenic score: %v\n", maxScore)
	fmt.Println("it's done")
}
