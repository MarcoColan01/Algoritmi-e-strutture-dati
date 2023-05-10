//data una slice di interi, viene costrutito il corrispondente albero binario quasi completo
package main

import(
	"fmt"
)

//STRUTTURA DATI DI RIFERIMENTO
type treenode struct{
	val int
	dx *treenode
	sx *treenode
}

type tree struct{
	root *treenode
}

func newTree() *tree{
	tree := new(tree)
	tree.root = nil 
	return tree
}

func newTreenode(n int) *treenode{
	new := new(treenode)
	new.dx = nil 
	new.sx = nil 
	new.val = n
	return new
}

func printNode(n *treenode){
	fmt.Print(n.val, " ")
}

func printTree(t *treenode){		//stampa l'albero con una DFS
	if t == nil{
		return
	}
	printTree(t.sx)
	printNode(t)
	printTree(t.dx)
}

func array2tree(a []int, size int, i int) *treenode{
	if i >= size{
		return nil
	}
	t := newTreenode(a[i])
	t.sx = array2tree(a, size, 2*i +1)
	t.dx = array2tree(a, size, 2*i +2)
	return t
}

func main(){
	t := newTree()
	vet := []int{3,4,5,1,23,12,9,8,14,25,63,32,30,20,2}
	t.root = array2tree(vet, len(vet), 0)
	printTree(t.root)
	fmt.Println()
}