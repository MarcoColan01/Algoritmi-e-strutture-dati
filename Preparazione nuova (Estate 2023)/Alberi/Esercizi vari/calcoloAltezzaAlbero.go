/*
Esercizio proposto da ChatGPT: scrivere una funzione che calcola l'altezza di un albero binario

*/

package main

import "fmt"

type treeNode struct {
	val    int
	dx, sx *treeNode
}

type tree struct {
	root *treeNode
}

func visitInorder(r *treeNode) {
	if r == nil {
		return
	}
	visitInorder(r.sx)
	fmt.Printf("%d ", r.val)
}

func calcoloAltezza(r *treeNode) int {
	if r == nil {
		return 0
	}
	hsx := calcoloAltezza(r.sx)
	hdx := calcoloAltezza(r.dx)
	if hsx > hdx {
		return hsx + 1
	} else {
		return hdx + 1
	}
}

func main() {
	r := new(tree)
	r.root = &treeNode{10, nil, nil}
	r.root.dx = &treeNode{50, nil, nil}
	r.root.sx = &treeNode{20, nil, nil}
	r.root.sx.sx = &treeNode{3, nil, nil}
	r.root.sx.sx.sx = &treeNode{1, nil, nil}
	visitInorder(r.root)

	h := calcoloAltezza(r.root)
	fmt.Printf("Altezza albero: %d\n", h)

}
