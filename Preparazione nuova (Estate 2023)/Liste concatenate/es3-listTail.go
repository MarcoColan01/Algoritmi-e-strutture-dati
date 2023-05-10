package main 
import(
	"fmt"
)

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

func addNewNodeAtEnd(l* linkedListWithTail, val int){
	if l.tail == nil{
		l.tail = newNode(val)
		l.head = l.tail 
	}else{
		/*Codice A. Problema: non viene allocata la memoria necessaria per il nuovo nodo con newNode()
		l.tail.next = val 
		l.tail = val 
		*/

		/*Codice B. Problema: l.tail non viene aggiornato 
		temp := newNode(val)
		l.tail.next = temp 
		*/

		/*Codice C. Problema: si perde il riferimento al nodo che prima era tail 
		temp := newNode(val)
		l.tail = temp 
		*/

		// codice D corretto 
		l.tail.next = newNode(val) 
		l.tail = l.tail.next

		/*Codice E. Problema:

		*/
	}
}

func printList(l *linkedListWithTail){
	node := l.head 
	for node != nil{
		fmt.Printf("%d -> ", node.item)
		node = node.next 
	}
	fmt.Println()
}

func main(){
	l := new(linkedListWithTail)
	addNewNodeAtEnd(l, 50)
	addNewNodeAtEnd(l, 10)
	addNewNodeAtEnd(l, 5)
	printList(l)

}