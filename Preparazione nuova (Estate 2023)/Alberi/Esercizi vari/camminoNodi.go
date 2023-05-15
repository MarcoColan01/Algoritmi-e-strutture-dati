//Ricerca del cammino: Scrivi una funzione che accetti un albero binario e due valori di nodi e restituisca true se esiste un cammino che collega i due nodi, altrimenti restituisca false.

package main

import "fmt"

type treeNode struct {
	val    int
	dx, sx *treeNode
}

type tree struct {
	root *treeNode
}

func searchNode(t *treeNode, n int) *treeNode {
	if t == nil {
		return nil
	}

	if t.val == n {
		return t
	}

	foundNode := searchNode(t.sx, n)
	if foundNode != nil {
		return foundNode
	}

	return searchNode(t.dx, n)
}

func main() {
	r := new(tree)
	var n1, n2 int
	r.root = &treeNode{10, nil, nil}
	r.root.dx = &treeNode{30, nil, nil}
	r.root.sx = &treeNode{9, nil, nil}
	r.root.dx.dx = &treeNode{12, nil, nil}
	r.root.dx.sx = &treeNode{59, nil, nil}
	r.root.sx.sx = &treeNode{23, nil, nil}
	r.root.sx.dx = &treeNode{19, nil, nil}
	r.root.dx.dx.dx = &treeNode{34, nil, nil}
	r.root.dx.dx.sx = &treeNode{7, nil, nil}

	fmt.Println("Inserisci il nodo da cui parte il cammino e quello da raggiungere: ")
	fmt.Scanf("%d %d", &n1, &n2)
	t := searchNode(r.root, n1)
	if t == nil {
		fmt.Printf("Il nodo di partenza con valore %d non è presente nell'albero\n", n1)
	} else {
		k := searchNode(t, n2)
		if k != nil {
			fmt.Printf("Percorso %d-%d esistente\n", n1, n2)
		} else {
			fmt.Printf("Percorso %d-%d NON esistente\n", n1, n2)
		}
	}
}
