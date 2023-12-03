package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func gamePossible(maxRed int, maxGreen int, maxBlue int, games map[int][]map[string]int) int {
	possible := true
	sum := 0
	for gameID, game := range games {
		fmt.Printf("GameID: %v ,game: %v\n", gameID, game)
		for _, turn := range game {
			fmt.Printf("turn: %v\n", turn)
			if turn["red"] > maxRed || turn["blue"] > maxBlue || turn["green"] > maxGreen {
				possible = false
			}
		}
		if possible {
			sum += gameID
		} else {
			possible = true
		}
	}
	return sum
}

func power(games map[int][]map[string]int) {
	sum := 0
	// p = minGreen * minBlue * minRed
	for _, game := range games {
		minGreen := 0
		minBlue := 0
		minRed := 0
		for _, turn := range game {
			if turn["red"] > minRed {
				minRed = turn["red"]
			}
			if turn["blue"] > minBlue {
				minBlue = turn["blue"]
			}
			if turn["green"] > minGreen {
				minGreen = turn["green"]
			}
		}
		p := minGreen * minBlue * minRed
		fmt.Printf("The power of the minimum set of cubes in game: %d\n", p)
		sum += p
	}
	fmt.Printf("the sum of the power of these sets: %d\n", sum)
}

func main() {
	fmt.Println("welcome 2023! day 2!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	game := make(map[int][]map[string]int)
	// var games []map[int][]map[string]int

	for scanner.Scan() {
		line := scanner.Text()
		_, a, _ := strings.Cut(line, " ")
		id, turnsString, _ := strings.Cut(a, ":")
		idNum, _ := strconv.Atoi(id)
		turns := strings.Split(turnsString, ";")
		fmt.Printf("the line: %s\n", line)
		var gameTurns []map[string]int

		for _, turn := range turns {
			cubes := strings.Split(turn, ",")
			thisTurn := make(map[string]int)
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				cubeCount, cubeColor, _ := strings.Cut(cube, " ")
				cubeCountNum, _ := strconv.Atoi(cubeCount)
				thisTurn[cubeColor] = cubeCountNum
			}
			gameTurns = append(gameTurns, thisTurn)
		}
		game[idNum] = gameTurns
		// games = append(games, game)
		fmt.Printf("id: %d", idNum)

	}
	fmt.Printf("Games: %v\n", game)

	s := gamePossible(12, 13, 14, game)

	fmt.Printf("sum of possible games: %d\n", s)
	power(game)

}
