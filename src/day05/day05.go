package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day05/moves.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stacks = getStacks()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])
		for i := 0; i < amount; i++ {
			n := len(stacks[from-1]) - 1
			stacks[to-1] = append(stacks[to-1], stacks[from-1][n])
			stacks[from-1] = stacks[from-1][:n]
		}
	}
	fmt.Print("one: ")
	for _, c := range stacks {
		fmt.Print(c[len(c)-1])
	}
	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day05/moves.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var stacks = getStacks()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])
		n := len(stacks[from-1]) - amount
		stacks[to-1] = append(stacks[to-1], stacks[from-1][n:]...)
		stacks[from-1] = stacks[from-1][:n]
	}
	fmt.Print("two: ")
	for _, c := range stacks {
		fmt.Print(c[len(c)-1])
	}
	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getStacks() [][]string {
	absPath, _ := filepath.Abs("src/day05/stacks.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stacks [][]string
	for scanner.Scan() {
		line := scanner.Text()
		stacks = append(stacks, strings.Split(line, ""))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stacks
}
