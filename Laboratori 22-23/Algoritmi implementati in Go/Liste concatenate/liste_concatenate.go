/*
Lista concatenata semplice con puntatore alla testa della lista. 
Operazioni da implementare:
	-printList(l): stampa la lista l
	-newList(): crea una nuova lista l
	-insertAtHead(elem): inserisce un elemento nella lista tramite la funzione newNode(elem)
	-find(elem): restituisce true se trova elem nella lista, false altrimenti
	-delete(elem): cancella elem dalla lista
	-newNode(n): crea un nuovo nodo con valore n
	isEmpty(l): restituisce true se la lista è vuota, false altrimenti
*/

package main

import(
	"fmt"
)

type listNode struct{
	val int 
	next *listNode
}

type linkedList struct{
	head *listNode
}

func newList() *linkedList{
	l := new(linkedList)
	l.head = nil
	return l
}

func newNode(n int) *listNode{
	new := new(listNode)
	new.val = n 
	new.next = nil
	return new
}

func isEmpty(l *linkedList) bool{
	return l.head == nil
}

func find(l *linkedList, n int) bool{
	app := l.head 
	trovato := false
	for{
		if app == nil{
			break
		}else if app.val == n{
			trovato = true
			break
		}
		app = app.next
	}
	return trovato
}

func delete(l *linkedList, n int){
	var cur, prev *listNode 
	cur = l.head 
	if cur.val == n{		//il nodo da cancellare è in testa alla lista
		cur = cur.next
		l.head = cur 
		return
	}else{
		for{
			if cur == nil{
				break
			}
			if cur.val == n{
				prev.next = cur.next 
				break
			}
			prev = cur 
			cur = cur.next
		}
	}
	return
}

func printList(l *linkedList){
	app := l.head 
	for{
		if app == nil{
			break
		}
		fmt.Print(app.val, " -> ")
		app = app.next
	}

	fmt.Println()
	return
}

func insertAtHead(l *linkedList, n int){
	new := newNode(n)
	new.next = l.head
	l.head = new
	return
}


func main(){
	var ch string
	var n int 
	var l *linkedList
	exit := false
	for !exit{
		fmt.Scanf("%s %d", &ch, &n)
		switch ch{
			case "n":
				l = newList()
			case "+":
				insertAtHead(l, n)
			case "-":
				delete(l, n)
			case "?":
				if find(l, n){
					fmt.Println("Elemento", n, "trovato")
				}else{
					fmt.Println("Elemento", n, "NON trovato")
				}
			case "p":
				printList(l)
			case "f":
				exit = true
				break
		}
	}
}

