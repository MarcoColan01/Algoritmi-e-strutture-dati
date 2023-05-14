/*
Scrivete un programma che:
	1. generi a caso una sequenza di interi (di lunghezza massima fissata con una opportuna macro) e la memorizzi in una slice;
	2. costruisca un albero binario a partire dalla slice (come descritto in seguito);
	3. stampi l’albero nella reppresentazione a sommario (come descritta in seguito);
	4. stampi i nodi dell’albero negli ordini determinati rispettivamente dalle visite in preordine,
inordine e postordine.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type treeNode struct {
	val    int
	dx, sx *treeNode
}

type tree struct {
	root *treeNode
}

func visitPreorder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.val)
	visitPreorder(node.sx)
	visitPreorder(node.dx)
}

func visitInorder(node *treeNode) {
	if node == nil {
		return
	}
	visitInorder(node.sx)
	fmt.Printf("%d ", node.val)
	visitInorder(node.dx)
}

func visitPostorder(node *treeNode) {
	if node == nil {
		return
	}
	visitPostorder(node.sx)
	visitPostorder(node.dx)
	fmt.Printf("%d ", node.val)
}

func stampaAlberoASommario(node *treeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	if node == nil {
		fmt.Println()
	} else {
		fmt.Println(node.val)
		if node.dx != nil || node.sx != nil {
			stampaAlberoASommario(node.sx, spaces+1)
			stampaAlberoASommario(node.dx, spaces+1)
		}
	}
}

func array2tree(a []int, i int) (root *treeNode) {
	if i >= len(a) {
		return nil
	}
	root = &treeNode{a[i], nil, nil}
	root.sx = array2tree(a, 2*i+1)
	root.dx = array2tree(a, 2*i+2)
	return root
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 10)

	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Printf("Array: %d\n", a)
	r := array2tree(a, 0)
	visitInorder(r)
	stampaAlberoASommario(r, 0)

}
