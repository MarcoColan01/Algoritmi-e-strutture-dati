
/*
Funzioni da implementare:
    -createQueue(): crea una nuova coda 
    -isEmpty(): restituisce true se la coda è vuota, false altrimenti
    -printQueue(): stampa la coda
    -enqueue(elem): aggiunge l'elemento alla coda (ultimo elemento)
    -dequeue(): rimuove il primo elemento della coda e lo restituisce
    -first(): restituisce il primo elemento della coda senza rimuoverlo
*/

package main

import(
	"fmt"
)

//struttura dati: lista concatenata con puntatore alla testa e alla coda

type listNode struct{
	val int 
	next *listNode 
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

func printNode(l *listNode){
	fmt.Print(l.val, " ")
}

func newNode(n int) *listNode{
	new := new(listNode)
	new.next = nil 
	new.val = n
	return new
}

func printQueue(q *queue){
	app := q.head 
	for{
		if app == nil{
			break
		}
		printNode(app)
		app = app.next
	}
	fmt.Println()
	return 
}

func enqueue(q *queue, n int){
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

func first(q *queue) *listNode{
	return q.head
}

func isEmpty(q *queue) bool{
	return q.head == nil
}

func main(){
	var q *queue
	var ch string
	var n int 
	exit := false
	for !exit{
		fmt.Scanf("%s %d", &ch, &n)
		switch ch{
			case "n":
				q = createQueue()
			case "+":
				enqueue(q, n)
			case "-":
				n := dequeue(q)
				if n != nil{
					fmt.Println("Elemento", n.val, "rimosso dalla coda")
				}else{
					fmt.Println("No dequeue: coda vuota")
				}
			case "p":
				printQueue(q)
			case "?":
				if isEmpty(q){
					fmt.Println("La coda è vuota")
				}else{
					fmt.Println("La coda NON è vuota")
				}
			case "1":
				n := first(q)
				if n != nil{
					fmt.Println("Primo elemento della coda: ", n.val)
				}else{
					fmt.Println("No first: coda vuota")
				}
			case "f":
				exit = true
				break
		}
	}
}