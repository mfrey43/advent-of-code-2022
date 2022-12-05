package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day00/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// yes, I'm lazy
	first := 0
	second := 0
	third := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > first {
				third = second
				second = first
				first = sum
			} else if sum > second {
				third = second
				second = sum
			} else if sum > third {
				third = sum
			}
			sum = 0
			continue
		}
		num, _ := strconv.Atoi(line)
		sum += num
	}
	if sum > first {
		third = second
		second = first
		first = sum
	} else if sum > second {
		third = second
		second = sum
	} else if sum > third {
		third = sum
	}
	fmt.Println(first + second + third)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
}
