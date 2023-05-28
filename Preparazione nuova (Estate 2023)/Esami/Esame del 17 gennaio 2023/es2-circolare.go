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

//func sposta(p *circNode){}

func main() {
	c := new(circList)
	insertCirc(c, 3)
	insertCirc(c, -2)
	insertCirc(c, 0)
	insertCirc(c, 1)
	insertCirc(c, 7)
	k := f2(c.head, 3)
	fmt.Println(k)
}
