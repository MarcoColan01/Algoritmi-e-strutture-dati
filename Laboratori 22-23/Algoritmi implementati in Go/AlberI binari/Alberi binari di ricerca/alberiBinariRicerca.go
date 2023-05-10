package main 

import (
	"fmt"
	"os"
	"bufio"
)


/*
Proprietà fondamentale alberi binari di ricerca: per ogni nodo, la chiave contenuta nel figlio sx è minore/uguale della sua chiave e quella contenuta nel figlio dx è maggiore della sua chiave.

Funzioni da implementare:
	dfs: esegue una visita in ampiezza dell'albero  -> implementare code
	bfs: esegue una visita in profondità dell'albero (non serve la pila)
	pre: esegue una visita in preordine dell'albero
	post: esegue una visita in postordine dell'albero
	+ n: aggiunge un nodo con valore n all'albero in base al criterio di inserimento per gli alberi binari di ricerca (si assume di ottenere un albero quasi completo) chiamando newNode 
	? n: restituisce true se n è presente nell'albero e false altrimenti
	p: stampa un sommario grafico dell'albero (farla sia ricorsiva che iterativa)
	new: crea un nuovo albero binario di ricerca FATTO
	c: restituisce il numero di nodi dell'albero
	d: cancella tutti i nodi dell'albero 
	per ora niente cancellazione 
*/

//STRUTTURA DATI DI RIFERIMENTO
type bistTreenode struct{
	val int 
	dx *bistTreenode
	sx *bistTreenode
}

type bistree struct{
	root *bistTreenode
}

//IMPLEMENTAZIONE CODA SEMPLICE PER LA BFS
type listNode struct{
	next *listNode 
	tn *bistTreenode
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

func newNode(n *bistTreenode) *listNode{
	new := new(listNode)
	new.next = nil 
	new.tn = n
	return new
}

func enqueue(q *queue, n *bistTreenode){
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

//FINE IMPLEMENTAZIONE CODA

func newBistree() *bistree{
	tree := new(bistree)
	tree.root = nil 
	return tree
}


func newBistreenode(n int) *bistTreenode{
	new := new(bistTreenode)
	new.dx = nil 
	new.sx = nil 
	new.val = n
	return new
}

func printTreenode(t *bistTreenode){
	fmt.Print(t.val, " ")
}

func dfs(t *bistTreenode){		//inorder
	if t == nil{
		return
	}
	dfs(t.sx)
	printTreenode(t)
	dfs(t.dx)
}

func bfs(t *bistTreenode){
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

func printPreorder(t *bistTreenode){
	if t == nil{
		return
	}
	printTreenode(t)
	printPreorder(t.sx)
	printPreorder(t.dx)
}

func printPostorder(t *bistTreenode){
	if t == nil{
		return
	}
	printPreorder(t.sx)
	printPreorder(t.dx)
	printTreenode(t)
}

func printTreeSummary(t *bistTreenode, spaces int){
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

func nNodi(t *bistTreenode) int{
	if t == nil{
		return 0
	}else{
		nsx := nNodi(t.sx)
		ndx := nNodi(t.dx)
		return 1+nsx+ndx
	}
}

func destroyTree(t *bistree){
	t.root = nil 
	return 
}

func insert(t *bistTreenode, n int) *bistTreenode{
	if t == nil{
		t = newBistreenode(n)
	}else if n < t.val{
		t.sx = insert(t.sx, n) 
	}else{
		t.dx = insert(t.dx, n)
	}
	return t
}

func search(t *bistTreenode, n int) bool{
	for t != nil && t.val != n{
		if n < t.val{
			t = t.sx
		}else{
			t = t.dx
		}
	}
	trovato := false
	if t == nil{
		trovato = false
	}else if t.val == n{
		trovato = true
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
	t := newBistree()
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
				t.root = insert(t.root, n)
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