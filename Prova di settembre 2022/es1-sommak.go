/*
Progettate e implementate un algoritmo che, dato un intero k e un albero binario t, restituisce true solo se nell'albero è presente una coppia di nodi la cui somma è pari a k

*/

package main

import "fmt"

type treeNode struct {
	val int
	dx  *treeNode
	sx  *treeNode
}

type bitree struct {
	root *treeNode
}

func newTree() *bitree {
	t := new(bitree)
	t.root = nil
	return t
}

func newTreeNode(n int) *treeNode {
	node := new(treeNode)
	node.val = n
	node.dx = nil
	node.sx = nil
	return node
}

func array2tree(a []int, size int, i int) *treeNode {
	if i >= size {
		return nil
	}
	t := newTreeNode(a[i])
	t.sx = array2tree(a, size, 2*i+1)
	t.dx = array2tree(a, size, 2*i+2)
	return t
}

func dfs(t *treeNode, b *[]int) {
	if t == nil {
		return
	}
	dfs(t.sx, b)
	*b = append(*b, t.val)
	dfs(t.dx, b)
}

func f(t *treeNode, k int) bool {
	trovata := false
	b := make([]int, 0)
	//b = bfs(t)
	dfs(t, &b)
	fmt.Println(b)
	for i := 0; i < len(b); i++ {
		n := b[i]
		for j := i + 1; j < len(b); j++ {
			if n+b[j] == k {
				trovata = true
				break
			}
		}
	}
	return trovata
}

func main() {
	k := 22
	a := []int{3, 4, 9, 12, 34, 5, 8, 9}

	t := newTree()
	t.root = array2tree(a, len(a), 0)

	if f(t.root, k) {
		fmt.Printf("Nell'albero esistono due nodi di somma %d\n", k)
	} else {
		fmt.Printf("Nell'albero NON esistono due nodi di somma %d\n", k)
	}
}
