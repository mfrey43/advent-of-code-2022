package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day02/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	sumTwo := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += getMatchScore(line)
		sumTwo += getMatchScorePredict(line)
	}
	fmt.Println("one: ", sum)
	fmt.Println("two: ", sumTwo)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
}

func getMatchScore(match string) int {
	switch match {
	case "A X": // 3 + 1
		return 4
	case "A Y": // 6 + 2
		return 8
	case "A Z": // 0 + 3
		return 3
	case "B X": // 0 + 1
		return 1
	case "B Y": // 3 + 2
		return 5
	case "B Z": // 6 + 3
		return 9
	case "C X": // 6 + 1
		return 7
	case "C Y": // 0 + 2
		return 2
	case "C Z": // 3 + 3
		return 6
	}
	return 0
}

func getMatchScorePredict(match string) int {
	switch match {
	case "A X": // 0 + 3
		return 3
	case "A Y": // 3 + 1
		return 4
	case "A Z": // 6 + 2
		return 8
	case "B X": // 0 + 1
		return 1
	case "B Y": // 3 + 2
		return 5
	case "B Z": // 6 + 3
		return 9
	case "C X": // 0 + 2
		return 2
	case "C Y": // 3 + 3
		return 6
	case "C Z": // 6 + 1
		return 7
	}
	return 0
}
