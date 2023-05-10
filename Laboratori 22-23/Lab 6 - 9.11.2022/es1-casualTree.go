/*
Scrivete un programma che:
	1. generi a caso una sequenza di interi (di lunghezza massima fissata con una opportuna
	macro) e la memorizzi in una slice;
	2. costruisca un albero binario a partire dalla slice (come descritto in seguito);
	3. stampi l’albero nella reppresentazione a sommario (come descritta in seguito);
	4. stampi i nodi dell’albero negli ordini determinati rispettivamente dalle visite in preordine, inordine e postordine.
*/

package main
import(
	"fmt"
	"math/rand"
	"time"
)
var n int = 10

type bitreeNode struct{
	val int 
	dx, sx *bitreeNode
}

type bitree struct{
	root *bitreeNode
}

func createBitree() *bitree{
	t := new(bitree)
	t.root = nil 
	return t
}

func newBitreeNode(n int) *bitreeNode{
	node := new(bitreeNode)
	node.val = n 
	node.dx = nil 
	node.sx = nil
	return node
}

func casualInts(a []int, n int) []int{
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < n; i++{
		k := rand.Intn(100)
		a = append(a, k)
	}
	return a
}

func array2tree(a[]int, i int) *bitreeNode{
	if i >= len(a){
		return nil
	}
	node := newBitreeNode(a[i])
	node.sx = array2tree(a, 2*i +1)
	node.dx = array2tree(a, 2*i +2)
	return node 
	
}

func printBitreeNode(n *bitreeNode){
	fmt.Print(n.val, " ")
}

func inorder(root *bitreeNode){
	if root == nil{
		return
	}
	inorder(root.sx)
	printBitreeNode(root)
	inorder(root.dx)
}

func preorder(root *bitreeNode){
	if root == nil{
		return
	}
	printBitreeNode(root)
	preorder(root.sx)
	preorder(root.dx)
}

func postorder(root *bitreeNode){
	if root == nil{
		return
	}
	postorder(root.sx)
	postorder(root.dx)
	printBitreeNode(root)
}

func stampaAlberoSommario(root *bitreeNode, spaces int){
	for i:=0; i < spaces; i++{
		fmt.Print(" ")
	}
	fmt.Print("*")
	if root == nil{
		fmt.Println()
	}else{
		fmt.Println(root.val)
		if root.dx != nil || root.sx != nil{
			stampaAlberoSommario(root.sx, spaces+1)
			stampaAlberoSommario(root.dx, spaces+1)
		}
	}
}

func main(){
	t := createBitree()
	a := make([]int, 0)
	a = casualInts(a, n)
	fmt.Println("Array:", a)
	t.root = array2tree(a, 0)
	fmt.Print("Visita in ordine (dfs): ")
	inorder(t.root)		//in ordine
	fmt.Println()
	fmt.Print("Visita in preordine: ")
	preorder(t.root)		//in ordine
	fmt.Println()
	fmt.Print("Visita in postordine: ")
	postorder(t.root)		//in ordine
	fmt.Println("\nSommario:")
	stampaAlberoSommario(t.root, 0)
}