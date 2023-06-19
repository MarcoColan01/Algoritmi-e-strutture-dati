package main

import "fmt"

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

func f2(p *circNode, k int) int {
	cont := 1
	app := p.prev
	for {
		if k == cont {
			fmt.Printf("%d ", app.val)
		}
		if app == p {
			break
		}
		app = app.prev
		cont++
	}
	return cont
}

func printList(c *circNode) {
	app := c
	for {
		fmt.Printf("%d -> ", app.val)
		app = app.next
		if app == c {
			break
		}
	}
}

func sposta(p *circNode) {
	if p == nil {
		return
	}

	numPosizioni := p.val
	if numPosizioni > 0 {
		// Sposta il nodo in avanti
		for i := 0; i < numPosizioni; i++ {
			p = p.next
		}
	} else if numPosizioni < 0 {
		// Sposta il nodo all'indietro
		numPosizioni = -numPosizioni
		for i := 0; i < numPosizioni; i++ {
			p = p.next
		}
	}
	prevNode := p
	for prevNode.next != p {
		prevNode = prevNode.next
	}

	prevNode.next = p.next
	for p.next != nil {
		p = p.next
	}
	p.next = prevNode

}

func main() {
	c := new(circList)
	insertCirc(c, 3)
	insertCirc(c, -2)
	insertCirc(c, 0)
	insertCirc(c, 1)
	insertCirc(c, 7)
	k := f2(c.head, 3)
	fmt.Println(k)
	printList(c.head)
	sposta(c.head)
	fmt.Println()
	printList(c.head)
}
