/*
Funzioni da implementare:
	-newOlist(): crea una nuova lista 
	-find(l, elem): restituisce true se trova elem nella lista l, false altrimenti. Ricerca dicotomica 
	-insert(l, elem): inserisce elem nella lista ordinata l
	-printList(): stampa la lista 
	-delete(l, elem): cancella elem dalla lista
	-newNode
*/

package main

import (
	"fmt"
)

type listNode struct{
	val int 
	next *listNode
}

type OlinkedList struct{
	head *listNode
}

func newNode(n int) *listNode{
	new := new(listNode)
	new.next = nil
	new.val = n
	return new
}

func newOlist() *OlinkedList{
	l := new(OlinkedList)
	l.head = nil 
	return l
}

func printOlist(l *OlinkedList){
	app := l.head 
	for{
		if app == nil{
			fmt.Println()
			return 
			break
		}
		fmt.Print(app.val, " -> ")
		app = app.next 
	}
	return 
}

func findOlist(l *OlinkedList, n int) bool{
	trovato := false
	app := l.head
	for{
		if !(app != nil && app.val <= n){
			return trovato
			break
		}
		if app.val == n{
			trovato = true
			break
		}
		app = app.next
	}
	return trovato
}


func delete(l *OlinkedList, n int){
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

func insertOlist(l *OlinkedList, n int){		//inserimenti dalla testa
	new := newNode(n)
	cur := l.head
	var prev *listNode 
	prev = nil 
	if l.head == nil || new.val < l.head.val{		//lista con un solo elemento
		new.next = l.head
		l.head = new
		}else{
			for{
				if cur != nil && cur.val < n{
					prev = cur
            		cur = cur.next
				}else{
					prev.next = new
        			new.next = cur
					break
				}
			}
		}
	return
}

func main(){
	var ch string
	var n int 
	var l *OlinkedList
	exit := false
	for !exit{
		fmt.Scanf("%s %d", &ch, &n)
		switch ch{
			case "n":
				l = newOlist()
			case "+":
				insertOlist(l, n)
			case "-":
				delete(l, n)
			case "?":
				if findOlist(l, n){
					fmt.Println("Elemento", n, "trovato")
				}else{
					fmt.Println("Elemento", n, "NON trovato")
				}
			case "p":
				printOlist(l)
			case "f":
				exit = true
				break
		}
	}
}