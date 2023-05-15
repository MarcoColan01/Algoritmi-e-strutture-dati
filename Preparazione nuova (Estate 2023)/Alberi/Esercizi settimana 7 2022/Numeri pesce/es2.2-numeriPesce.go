package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func readTree(s string) *TreeNode {
	stack := []*TreeNode{}
	var root *TreeNode
	i := 0

	for i < len(s) {
		if s[i] == '[' {
			stack = append(stack, root)
			root = nil
		} else if s[i] == ']' {
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if parent.Left == nil {
					parent.Left = root
				} else {
					parent.Right = root
				}
				root = parent
			}
		} else {
			j := i
			for j < len(s) && s[j] != ',' && s[j] != ']' {
				j++
			}
			value := s[i:j]
			if strings.Contains(value, "[") {
				root = &TreeNode{
					Val:   readTree(value),
					Left:  nil,
					Right: nil,
				}
			} else {
				// Parse integer value
				// Assuming input contains valid integers
				// Error handling can be added if needed
				root = &TreeNode{
					Val:   value,
					Left:  nil,
					Right: nil,
				}
			}
			i = j - 1
		}
		i++
	}

	return root
}

func main() {
	s := "[[1,2],[3,[4,5]]]"
	tree := readTree(s)
	fmt.Println(tree)
}
