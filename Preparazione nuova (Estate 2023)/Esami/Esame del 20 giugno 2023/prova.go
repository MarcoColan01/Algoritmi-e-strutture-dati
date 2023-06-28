package main

import (
	"fmt"
)

type node struct {
	left      *node
	right     *node
	etichetta rune
}

type tree = *node

func buildTree(s string) tree {
	if len(s) == 1 {
		return &node{etichetta: rune(s[0])}
	}

	if s[0] != '[' || s[len(s)-1] != ']' {
		panic("Stringa-albero non valida")
	}

	s = s[1 : len(s)-1]
	commaIndex := findCommaIndex(s)

	leftTree := buildTree(s[:commaIndex])
	rightTree := buildTree(s[commaIndex+1:])

	return &node{
		left:      leftTree,
		right:     rightTree,
		etichetta: '-',
	}
}

func findCommaIndex(s string) int {
	count := 0
	for i, char := range s {
		if char == '[' {
			count++
		} else if char == ']' {
			count--
		} else if char == ',' && count == 0 {
			return i
		}
	}

	panic("Stringa-albero non valida")
}

func main() {
	s := "[[[[1,2],[3,4]],[5,6],[7,8]],9]"
	root := buildTree(s)
	printTree(root)
}

func printTree(node tree) {
	if node == nil {
		return
	}

	printTree(node.left)
	fmt.Println(string(node.etichetta))
	printTree(node.right)
}
