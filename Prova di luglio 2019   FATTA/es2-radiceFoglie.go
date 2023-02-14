package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Bit_node struct {
	val    int
	dx, sx *Bit_node
}

type Bit_tree struct {
	root *Bit_node
}

func newTreeNode(n int) *Bit_node {
	node := new(Bit_node)
	node.dx = nil
	node.sx = nil
	node.val = n
	return node
}
func createTree(a []int, n int) *Bit_node {
	if n >= len(a) {
		return nil
	}
	node := newTreeNode(a[n])
	if a[n] != 0 {
		node.sx = createTree(a, 2*n+1)
		node.dx = createTree(a, 2*n+2)
	}
	return node
}

func stampaSommario(t *Bit_node, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	if t == nil {
		fmt.Println()
	} else {
		fmt.Println(t.val)
		if t.dx != nil || t.sx != nil {
			stampaSommario(t.sx, spaces+1)
			stampaSommario(t.dx, spaces+1)
		}
	}
}

func costoMax(t *Bit_node) int {
	if t == nil {
		return 0
	}
	if t.val+costoMax(t.sx) > t.val+costoMax(t.dx) {
		return t.val + costoMax(t.sx)
	} else {
		return t.val + costoMax(t.dx)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	filename := os.Args[1]
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	i := 0
	for scanner.Scan() {
		k, _ := strconv.Atoi(scanner.Text())
		a[i] = k
		i++
	}
	t := new(Bit_tree)
	t.root = createTree(a, 0)
	stampaSommario(t.root, 0)
	max := costoMax(t.root)
	fmt.Printf("Percorso radice-foglia con costo maggiore: %d\n", max)
}
