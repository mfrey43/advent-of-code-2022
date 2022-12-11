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
	absPath, _ := filepath.Abs("src/day04/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		split1 := strings.Split(ranges[0], "-")
		startA, _ := strconv.Atoi(split1[0])
		endA, _ := strconv.Atoi(split1[1])
		split2 := strings.Split(ranges[1], "-")
		startB, _ := strconv.Atoi(split2[0])
		endB, _ := strconv.Atoi(split2[1])
		if startA <= startB && endA >= endB || startA >= startB && endA <= endB {
			sum++
		}
	}
	fmt.Println("one:", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day04/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		split1 := strings.Split(ranges[0], "-")
		startA, _ := strconv.Atoi(split1[0])
		endA, _ := strconv.Atoi(split1[1])
		split2 := strings.Split(ranges[1], "-")
		startB, _ := strconv.Atoi(split2[0])
		endB, _ := strconv.Atoi(split2[1])
		if startA <= startB && endA >= endB || startA >= startB && endA <= endB || startA >= startB && startA <= endB || endA >= startB && endA <= endB {
			sum++
		}
	}
	fmt.Println("two:", sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
