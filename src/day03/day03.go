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
	absPath, _ := filepath.Abs("src/day03/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[rune]bool)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if i < len(line)/2 {
				m[c] = true
			} else {
				if m[c] {
					//fmt.Println(string(c), getPriority(c))
					sum += getPriority(c)
					m = make(map[rune]bool)
					break
				}
			}
		}
	}
	fmt.Println("one:", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day03/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := getMap()
	sum := 0
	i := 0 // current line of the group
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			m[i][c] = true
		}
		if i == 2 {
			for c := range m[0] {
				if m[0][c] && m[1][c] && m[2][c] {
					sum += getPriority(c)
					m = getMap()
					break
				}
			}
		}
		i = (i + 1) % 3
	}
	fmt.Println("two:", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getMap() map[int]map[rune]bool {
	m := make(map[int]map[rune]bool)
	m[0] = make(map[rune]bool)
	m[1] = make(map[rune]bool)
	m[2] = make(map[rune]bool)
	return m
}

func getPriority(c rune) int {
	if c >= 'a' {
		return int(c - 'a' + 1) // a - z => 1 - 26
	} else {
		return int(c - 'A' + 27) // A - Z => 27 - 52
	}
}
