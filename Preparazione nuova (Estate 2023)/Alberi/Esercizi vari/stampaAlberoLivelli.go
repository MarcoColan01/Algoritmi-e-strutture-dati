//Scrivi una funzione che accetti un albero binario e stampi i valori dei nodi per livelli, partendo dalla radice e continuando con i nodi dei livelli successivi.

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
	visitInorder(r.dx)
}

func PrintLevels(root *treeNode) {
	if root == nil {
		return
	}

	queue := []*treeNode{root} // Coda iniziale con la radice
	level := 1                 // Livello corrente

	for len(queue) > 0 {
		fmt.Printf("Livello %d: ", level)

		// Itera tutti i nodi nel livello corrente
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			fmt.Printf("%d ", node.val)

			// Aggiunge i figli del nodo corrente alla coda
			if node.sx != nil {
				queue = append(queue, node.sx)
			}
			if node.dx != nil {
				queue = append(queue, node.dx)
			}
		}

		fmt.Println() // Vai a capo dopo aver stampato i nodi del livello corrente

		// Rimuove i nodi del livello corrente dalla coda
		queue = queue[levelSize:]
		level++
	}
}

func main() {
	r := new(tree)
	r.root = &treeNode{10, nil, nil}
	r.root.dx = &treeNode{30, nil, nil}
	r.root.sx = &treeNode{9, nil, nil}
	r.root.dx.dx = &treeNode{12, nil, nil}
	r.root.dx.sx = &treeNode{59, nil, nil}
	r.root.sx.sx = &treeNode{23, nil, nil}
	r.root.sx.dx = &treeNode{19, nil, nil}
	r.root.dx.dx.dx = &treeNode{34, nil, nil}
	r.root.dx.dx.sx = &treeNode{7, nil, nil}
	//visitInorder(r.root)
	PrintLevels(r.root)
}
