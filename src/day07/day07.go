package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

type FileType int64

const (
	Directory FileType = iota
	File
)

type Node struct {
	name     string
	size     int
	fileType FileType
	parent   *Node
	children []*Node
}

func one() {
	absPath, _ := filepath.Abs("src/day07/input.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	root := &Node{
		name:     "/",
		size:     0,
		fileType: Directory,
		parent:   nil,
		children: []*Node{},
	}
	//m := make(map[string]Node)
	//m["/"] = root

	var current *Node
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, " ")
		if l[0] == "$" {
			if l[1] == "cd" {
				if l[2] == "/" {
					current = root
				} else if l[2] == ".." {
					current = current.parent
				} else {
					current = findChild(current.children, l[2])
				}
			}
		} else {
			if l[0] == "dir" {
				getOrCreate(current, l[1], 0)
			} else {
				size, _ := strconv.Atoi(l[0])
				getOrCreate(current, l[1], size)
			}
		}
	}
	calcSizes(root)
	printTree(root, 0)
	fmt.Println("one:", sumUpLessThan(root, 100000))

	diskSize := 70000000
	diskRequired := 30000000
	toFree := root.size - (diskSize - diskRequired)
	fmt.Println("two:", findSmallest(root, toFree))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findSmallest(node *Node, minSize int) int {
	if node.size < minSize {
		return math.MaxInt32
	}
	min := node.size
	for _, v := range node.children {
		smallestChild := findSmallest(v, minSize)
		if smallestChild < min {
			min = smallestChild
		}
	}
	return min
}

func printTree(node *Node, depth int) {
	fmt.Println(strings.Repeat("-", depth), node.name, node.size)
	for _, v := range node.children {
		printTree(v, depth+1)
	}
}

func two() {
}

func getOrCreate(parent *Node, name string, size int) {
	var fileType FileType
	if size > 0 {
		fileType = File
	} else {
		fileType = Directory
	}
	parent.children = append(parent.children, &Node{
		name:     name,
		size:     size,
		fileType: fileType,
		parent:   parent,
		children: []*Node{},
	})
}

func findChild(children []*Node, name string) *Node {
	for _, v := range children {
		if v.name == name && v.fileType == Directory {
			return v
		}
	}
	return nil
}

func calcSizes(node *Node) int {
	sum := 0
	for _, v := range node.children {
		if v.fileType == Directory {
			sum += calcSizes(v)
		} else {
			sum += v.size
		}
	}
	node.size = sum
	return sum
}

func sumUpLessThan(node *Node, max int) int {
	sum := 0
	for _, v := range node.children {
		if v.fileType == Directory {
			sum += sumUpLessThan(v, max)
		}
	}
	if node.size < max {
		return sum + node.size
	} else {
		return sum
	}
}
