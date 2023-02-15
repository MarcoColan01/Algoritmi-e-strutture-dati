package main

import "fmt"

type Bit_node struct {
	item, y int
	l, r    *Bit_node
}

func f(p *Bit_node, x int) {
	if p == nil {
		return
	}
	p.y = x + p.item
	f(p.l, p.y)
	f(p.r, p.y)
}

func dfs(r *Bit_node) {
	if r == nil {
		return
	}
	dfs(r.l)
	fmt.Printf("Valore: %d, y: %d\n", r.item, r.y)
	dfs(r.r)
}

func main() {
	tree := new(Bit_node)
	tree.r = new(Bit_node)
	tree.l = new(Bit_node)
	tree.r.r = new(Bit_node)
	tree.r.l = new(Bit_node)
	tree.l.l = new(Bit_node)
	tree.r.r.r = new(Bit_node)
	tree.r.r.l = new(Bit_node)

	tree.item = 4
	tree.r.item = 2
	tree.l.item = 5
	tree.r.r.item = 1
	tree.r.l.item = 13
	tree.l.l.item = 4
	tree.r.r.r.item = 7
	tree.r.r.l.item = 3
	f(tree, 3)
	dfs(tree)
}
