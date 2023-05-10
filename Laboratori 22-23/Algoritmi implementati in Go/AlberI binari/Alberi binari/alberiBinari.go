/*
Funzioni da implementare:
	dfs: esegue una visita in ampiezza dell'albero  -> implementare code
	bfs: esegue una visita in profondità dell'albero (non serve la pila)
	pre: esegue una visita in preordine dell'albero
	post: esegue una visita in postordine dell'albero
	+ n: aggiunge un nodo con valore n all'albero (si assume di ottenere un albero quasi completo) chiamando newNode 
	? n: restituisce true se n è presente nell'albero e false altrimenti
	p: stampa un sommario grafico dell'albero (farla sia ricorsiva che iterativa)
	new: crea un nuovo albero FATTO
	c: restituisce il numero di nodi dell'albero
	d: cancella tutti i nodi dell'albero 
	per ora niente cancellazione 

*/

package main

import(
	"fmt"
	"os"
	"bufio"
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

//IMPLEMENTAZIONE CODA SEMPLICE PER LA BFS

type listNode struct{
	next *listNode
	tn *treenode
}
type queue struct{
	head *listNode
	tail *listNode
}

func createQueue() *queue{
	q := new(queue)
	q.head = nil 
	q.tail = nil 
	return q
}

func newNode(n *treenode) *listNode{
	new := new(listNode)
	new.next = nil 
	new.tn = n
	return new
}

func enqueue(q *queue, n *treenode){
	new := newNode(n)
	if q.head == nil{		//coda vuota
		q.head = new 
		q.tail = new
	}else{
		q.tail.next = new 
		q.tail = q.tail.next
	}
	return
}

func dequeue(q *queue) *listNode{
	var n *listNode
	if q.head != nil{
		n = q.head
		q.head = q.head.next
	}else{
		n = nil
	}
	return n 
}

func isEmpty(q *queue) bool{
	return q.head == nil
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

func printTreenode(t *treenode){
	fmt.Print(t.val, " ")
}

func dfs(t *treenode){		//inorder
	if t == nil{
		return
	}
	dfs(t.sx)
	printTreenode(t)
	dfs(t.dx)
}

func bfs(t *treenode){
	q := createQueue()
	enqueue(q, t)
	for {
		if isEmpty(q){
			break
		}
		n := dequeue(q)
		if n.tn != nil{
			fmt.Print(n.tn.val , " ")
			enqueue(q, n.tn.sx)
			enqueue(q, n.tn.dx)
		}
	}
	fmt.Println()
	return
}

func printPreorder(t *treenode){
	if t == nil{
		return
	}
	printTreenode(t)
	printPreorder(t.sx)
	printPreorder(t.dx)
}

func printPostorder(t *treenode){
	if t == nil{
		return
	}
	printPreorder(t.sx)
	printPreorder(t.dx)
	printTreenode(t)
}

func insert(t *tree, n int){
	var x *listNode
	if t.root == nil{	//albero vuoto
		t.root = newTreenode(n)
	}else{
		q := createQueue()
		enqueue(q, t.root)
		for {
			if isEmpty(q){
				break
			}
		x = dequeue(q)
		if x.tn.dx != nil && x.tn.sx != nil{
			enqueue(q, x.tn.sx)
			enqueue(q, x.tn.dx)
		}else{
			break
		}
		}
		app := x.tn 
		new := newTreenode(n)
		if app.sx == nil{
			app.sx = new 
		}else{
			app.dx = new
		}
	}
}

func printTreeSummary(t *treenode, spaces int){
	for i:=0; i < spaces; i++{
		fmt.Print(" ")
	}
	fmt.Print("*")
	if t == nil{
		fmt.Println()
	}else{
		fmt.Println(t.val)
		if t.dx != nil || t.sx != nil{
			printTreeSummary(t.sx, spaces+1)
			printTreeSummary(t.dx, spaces+1)
		}
	}
}

func nNodi(t *treenode) int{
	if t == nil{
		return 0
	}else{
		nsx := nNodi(t.sx)
		ndx := nNodi(t.dx)
		return 1+nsx+ndx
	}
}

func destroyTree(t *tree){
	t.root = nil 
	return 
}

func search(t *treenode, n int) bool{
	trovato := false
	if t.val == n{
		return true 
	}else{
		q := createQueue()
		enqueue(q, t)
		for !isEmpty(q){
			x := dequeue(q)
			if x.tn != nil{
				if x.tn.val == n{
					trovato = true
					break
				}else{
					enqueue(q, x.tn.sx)
					enqueue(q, x.tn.dx)
				}
			}
		}
	} 
	return trovato
}

func main(){
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(file)
	t := newTree()
	var cmd string
	var n int
	exit := false
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan() && exit == false{
		k := scanner.Text()
		fmt.Sscanf(k, "%s %d", &cmd, &n)
		switch cmd{
			case "+":
				insert(t, n)
			case "?":
				if search(t.root, n){
					fmt.Println("Nodo", n, "trovato nell'albero")
				}else{
					fmt.Println("Nodo", n, "NON trovato nell'albero")
				}
			case "dfs":
				fmt.Print("Vista in profondità: ")
				dfs(t.root)
				fmt.Println()
			case "bfs":
				fmt.Print("Vista in ampiezza: ")
				bfs(t.root)
				fmt.Println()
			case "pre":
				fmt.Print("Vista in preordine: ")
				printPreorder(t.root)
				fmt.Println()
			case "post":
				fmt.Print("Vista in postordine: ")
				printPostorder(t.root)
				fmt.Println()
			case "c":
				k:= nNodi(t.root)
				fmt.Println("Numero nodi nell'albero: ", k)
			case "f":
				exit = true 
				break
			case "d":
				destroyTree(t)
				fmt.Println("Albero distrutto")
			case "p":
				printTreeSummary(t.root, 0)
		}
	}
}