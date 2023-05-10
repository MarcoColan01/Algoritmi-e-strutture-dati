package main 

import (
	"fmt"
)

//Considerate ora una lista con concatenata di interi, dotata di due riferimenti al primo e all’ultimo elemento della lista, come descritto dal tipo linkedListwithTail nella porzione di codice qui sotto. Quando la lista è vuota, sia head che tail sono nil. La funzione newNode alloca lo spazio per un nuovo nodo e ne inizializza il valore con l’argomento passato.

type listNode struct{
	item int 
	next *listNode
}

type linkedListWithTail struct{
	head, tail *listNode
}

func newNode(val int) *listNode{
	return &listNode{val, nil}
}

//Considerate la seguente funzione (incompleta) che aggiunge un elemento e in fondo alla lista:
func addNewNodeAtEnd(l *linkedListWithTail, val int){
	if l.tail == nil {			//lista vuota
		l.tail = newNode(val)
		l.head = l.tail
		//Quale dei seguenti frammenti di codice completa la funzione addAtEnd? Spiegate quali problemi si riscontrano scegliendo ciascuna delle altre opzioni. Il codice corretto che completa la funzione addNewNodeAtEnd è:
		
		} else {				//lista non vuota
			//risposta D
			l.tail.next = newNode(val)
			l.tail = l.tail.next
			//altri codici e relativi problemi
			/* 
			Codice A:
			l.tail.next = val
			l.tail = val
			Problema: val è un valore intero che va inserito in un nodo, ma non è il nodo stesso, quindi non può essere assegnato direttamente come un elemento della lista, va prima creato un nodo con newNode(val) e poi inserito in lista

			Codice B:
			temp := newNode (val)
			l.tail.next = temp
			Problema: Dopo l'inserimento del primo nodo in lista, questo avrà sempre il riferimento come coda, quindi quando si inserisce un elemento questo è il next della coda, per cui all'inserimento del secondo elemento  questo sarà il next della coda, ma quando ne viene inserito un altro il nodo precedente viene "perso" e il next della coda sarà, quindi, l'ultimo elemento inserito in ordine, quindi la lista avrà sempre 2 elementi: il primo e l'ultimo aggiunto

			Codice C:
			temp := newNode (val)
			l.tail = temp
			Problema: Viene salvato solo il primo elemento inserito in lista

			Codice E:
			temp := l.head
			for temp.next != nil {
				temp = temp.next
			}	
			temp.next = newNode(val)
			Problema: Non viene assegnato il riferimento alla coda, che deve essere l'ultimo elemento inserito
			*/

			

		}
}

func printList(l *linkedListWithTail){
	app := l.head 
	for{
		if app == nil{
			break
		}
		fmt.Print(app.item , " -> ")
		app = app.next
	}
	fmt.Println()
	return
}

func main(){
	l := new(linkedListWithTail)
	addNewNodeAtEnd(l, 10)
	addNewNodeAtEnd(l, 23)
	addNewNodeAtEnd(l, 12)
	addNewNodeAtEnd(l, 34)
	addNewNodeAtEnd(l, 67)
	printList(l)

}


/*
3.2 Confronto tra lista semplice e lista con tail
Nella lista concatenata linkedListWithTail abbiamo i riferimenti all’inizio e alla fine della lista. Confrontate questa implementazione con quella di una lista semplice che non ha un riferimento alla fine della lista, cioè definita come segue:
	type linkedList struct {
		head *listNode
	}

	type  linkedListWithTail struct{
		head, tail *listNode
	}
Per quali delle seguenti operazioni si ha un tempo migliore con linkedListWithTail piuttosto che con linkedList? Scegliete tutte le opzioni corrette e giustificate la risposta.
	A) restituisci 1 se la lista contiene un dato elemento
	B) cancella l’ultimo elemento della lista
	C) aggiungi un elemento all’inizio della lista
	D) aggiungi un elemento alla fine della lista

	A) 	Se l'elemento si trova in fondo alla lista:
			linkedList: tempo O(n) devo scorrere tutta la lista
			linkedListWithTail: tempo O(1) ho accesso diretto all'ultimo elemento
		Se l'elemento si trova in testa alla lista: tempo O(1) per entrambe
		Se l'elemento è in una qualunque posizione della lista: tempo O(n) per entrambi
		Quindi: se l'elemento si trova in fondo alla lista è meglio llwt, altrimenti va bene una qualunque delle due liste
	
	B) 	linkedList: tempo O(n) per scorrere tutta la lista e cancellare l'ultimo elemento
		linkedListWithTail: tempo O(1) grazie all'accesso diretto all'ultimo elemento
		Tempo migliore con linkedListWithTail

	C) linkedList richiede di spostare il puntatore alla testa delal lista e aggiungere l'elemento, linkedListWithTail richiede di spostare il puntatore alla testa della lista e alla coda se si inserisce un secondo elemento, altrimenti va cambiato solo un puntatore. In entrambi i casi il tempo è O(1) ed è simile

	D) 	Se la lista è vuota: tempo O(1) per entrambe
		Altrimenti:
			linkedList: tempo O(n) per scorrere tutta la lista e inserire in fondo
			linkedListWithTail: tempo O(1) grazie all'accesso diretto alla coda
		Tempo migliore con llwt se la lista non è vuota




*/