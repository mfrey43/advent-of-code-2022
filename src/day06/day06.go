package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	absPath, _ := filepath.Abs("src/day06/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, "")
		for i := 3; i < len(l); i++ {
			// force it :D
			if l[i] != l[i-1] && l[i] != l[i-2] && l[i] != l[i-3] && l[i-1] != l[i-2] && l[i-1] != l[i-3] && l[i-2] != l[i-3] {
				result = i + 1
				break
			}
		}
	}
	fmt.Println("one:", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func two() {
	absPath, _ := filepath.Abs("src/day06/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, "")
		result = findMatch(l, 14)
	}
	fmt.Println("two:", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findMatch(l []string, length int) int {
	m := make(map[string]int)
	for i := 0; i < len(l); i++ {
		m[l[i]]++
		if i >= length {
			m[l[i-length]]--
			dupe := false
			for k := i; k > i-length; k-- {
				if m[l[k]] > 1 {
					dupe = true
					continue
				}
			}
			if !dupe {
				return i + 1
			}
		}
	}
	return 0
}
