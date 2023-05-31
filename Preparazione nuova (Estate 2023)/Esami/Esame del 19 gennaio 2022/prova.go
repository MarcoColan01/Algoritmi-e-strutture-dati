package main

import "fmt"

type Bit_node struct {
	item, y int
	dx, sx  *Bit_node
}

type Bit_tree struct {
	root *Bit_node
}

func f(p *Bit_node, x int) {
	if p == nil {
		return
	}
	p.y = x + p.item
	f(p.sx, p.y)
	f(p.dx, p.y)

}

func dfs(r *Bit_node) {
	if r == nil {
		return
	}
	dfs(r.sx)
	fmt.Printf("Nodo %d  y: %d\n", r.item, r.y)
	dfs(r.dx)
}

func main() {
	t := new(Bit_tree)
	t.root = &Bit_node{4, 0, nil, nil}

	t.root.sx = &Bit_node{5, 0, nil, nil}
	t.root.dx = &Bit_node{2, 0, nil, nil}

	t.root.sx.sx = &Bit_node{4, 0, nil, nil}
	t.root.dx.sx = &Bit_node{13, 0, nil, nil}
	t.root.dx.dx = &Bit_node{1, 0, nil, nil}

	t.root.dx.dx.sx = &Bit_node{3, 0, nil, nil}
	t.root.dx.dx.dx = &Bit_node{7, 0, nil, nil}

	f(t.root, 2)
	dfs(t.root)
}
