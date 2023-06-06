package main

import "fmt"

type node struct {
	val    int
	dx, sx *node
}

type tree struct {
	root *node
}

func buildTree(a []int, size int, i int) *node {
	if i >= size {
		return nil
	}
	t := &node{a[i], nil, nil}
	if a[i] != 0 {
		t.sx = buildTree(a, size, 2*i+1)
		t.dx = buildTree(a, size, 2*i+2)
		return t
	} else {
		return nil
	}

}

func dfs(t *node) {
	if t == nil {
		return
	}
	dfs(t.sx)
	fmt.Println(t.val)
	dfs(t.dx)
}

func costoMax(t *node, somma int) int {
	if t == nil {
		return 0
	}
	sommaSx := costoMax(t.sx, somma)
	sommaDx := costoMax(t.dx, somma)

	tot := t.val + sommaSx + sommaDx
	if tot > somma {
		somma = tot
	}
	if t.val+sommaSx > t.val+sommaDx {
		return t.val + sommaSx
	} else {
		return t.val + sommaDx
	}
}

func main() {
	a := []int{10, 4, 7, 2, 0, 3, 6, 8, 9, 0, 0, 1, 2, 4, 5}
	n := 15
	t := new(tree)
	t.root = buildTree(a, n, 0)
	//dfs(t.root)
	k := costoMax(t.root, 0)
	fmt.Println(k)

}
