package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name     string
	size     int
	children []*Node
	parent   *Node
}

var currentNode *Node
var root Node

func computeSizes(node *Node) int {

	if node.children == nil {
		return node.size
	} else {
		for i := 0; i < len(node.children); i += 1 {
			node.size = node.size + computeSizes(node.children[i])
		}
	}

	return node.size

}

func findSmallDirs(node *Node, sum int) int {

	if (node.size <= 100000) && (node.children != nil) {
		sum += node.size
	}
	if node.children != nil {
		for i := 0; i < len(node.children); i += 1 {
			sum = sum + findSmallDirs(node.children[i], 0)
		}
	}
	return sum
}

func main() {
	fmt.Println("welcome to day 7 2022!")
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("looking at: %v\n", line)
		prefix := line[0:4]
		switch prefix {
		case "$ cd":
			if line == "$ cd /" {
				root.name = "/"
				root.size = 0
				root.children = make([]*Node, 0)
				root.parent = nil
				currentNode = &root
			} else {
				cdDirName := strings.Split(line, " ")
				if cdDirName[2] == ".." {
					currentNode = currentNode.parent
				} else {
					for i := 0; i < len(currentNode.children); i += 1 {
						if currentNode.children[i].name == cdDirName[2] {
							currentNode = currentNode.children[i]
							fmt.Printf("current dir is now: %v\n", cdDirName[2])
						}
					}
				}
			}
		case "$ ls":
			fmt.Println("will process $ ls")
		case "dir ":
			dirName := strings.Split(line, " ")[1]
			fmt.Printf("dir name is: %v \n", dirName)
			dir := Node{
				name:     dirName,
				size:     0,
				children: nil,
				parent:   currentNode,
			}
			currentNode.children = append(currentNode.children, &dir)
		default:
			fmt.Println("will process file as the fallback")
			fileAttrs := strings.Split(line, " ")
			size, err := strconv.Atoi(fileAttrs[0])
			if err != nil {
				log.Fatal(err)
			}
			file := Node{
				name:     fileAttrs[1],
				size:     size,
				children: nil,
				parent:   currentNode,
			}
			currentNode.children = append(currentNode.children, &file)
		}
	}
	root.size = computeSizes(&root)
	fmt.Printf("computed size of /: %v\n", root)
	smallDirs := findSmallDirs(&root, 0)
	fmt.Printf("small dirs: %v \n", smallDirs)
}
