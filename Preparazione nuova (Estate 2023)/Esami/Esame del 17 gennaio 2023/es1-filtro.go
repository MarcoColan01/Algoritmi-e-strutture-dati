package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Definite un tipo strutturato node compatibile con la funzione f
*/
type node struct {
	val  int
	next *node
}

type listNode struct {
	head *node
}

/*
Definite un tipo circNode che serve a rappresentare i nodi di una lista circolare di interi.
*/
type circNode struct {
	val        int
	next, prev *circNode
}

type circList struct {
	head, tail *circNode
}

func insertCirc(l *circList, elem int) {
	node := &circNode{elem, nil, nil}
	if l.head == nil { //Inserimento in lista vuota
		node.next = l.head
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		node.prev = l.tail
		l.tail.next = node
		l.tail = l.tail.next
		l.head.prev = l.tail
	}
}

func stampaDaZero(p *circNode) {
	app := p
	continua := true
	for continua {
		if app.val == 0 {
			k := app
			for {
				fmt.Printf("%d ", k.val)
				k = k.next
				if k.val == 0 {
					continua = false
					break
				}
			}
		} else {
			app = app.next
		}
	}
	fmt.Println()
}

func f(p *node, k int) int {
	var a int
	if p == nil {
		return 0
	}

	a = 1 + f(p.next, k)
	if a == k {
		fmt.Println(p.val)
	}
	return a
}

func insert(l *listNode, elem int) {
	node := &node{elem, nil}
	node.next = l.head
	l.head = node
}

func printList(l *node) {
	app := l
	for app != nil {
		fmt.Printf("%d -> ", app.val)
		app = app.next
	}
	fmt.Println()
}

func main() {
	c := new(circList)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		insertCirc(c, n)
	}
	/*
		insertCirc(c, 3)
		insertCirc(c, -2)
		insertCirc(c, 0)
		insertCirc(c, 1)
		insertCirc(c, 7)
	*/
	stampaDaZero(c.head)
}

/*
Rispondere alle seguenti domande:
1. Cosa stampa f(p,k) con k = 1 ?
La funzione stampa 7 5

2. Cosa stampa f(p,k) con k = 5 ?
La funzione stampa 3 5

3. Cosa stampa f(p,k) con k = 10 ?
La funzione stampa 5

4. Completare la frase: "Ricevendo il puntatore alla testa di una lista e un intero k, la funzione f stampa ... e restituisce ...". Includere anche casi particolari/limite, se ce ne sono.
Ricevendo il puntatore alla testa di una lista e un intero k, la funzione stampa la lunghezza della lista e restituisce, se presente, l'elemento in posizione k, contando però le posizioni da 1 ad len(lista) a partire dal fondo della lista. Se k > len(lista), la funzione non stampa nulla ma restituisce solo la lunghezza della lista.

5. Complessità della funzione
La funzione dovrà scandire l'intera lista, anche quando ha già trovato l'elemento in posizione k. Per cui la complessità in tempo è proporzionale alla lunghezza n della lista O(n). Essendo una funzione ricorsiva bisogna anche tenere conto dell'altezza dello stack per la ricorsione: questo è proporzionale alla lunghezza della lista, quindi O(n).
*/
