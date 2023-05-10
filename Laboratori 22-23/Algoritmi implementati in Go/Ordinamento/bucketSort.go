package main 

import (
	"fmt"
)

//il record va adattato in base al tipo di valore che si vuole memorizzare, l'unico vincolo è avere chiavi intere
type record struct{
	key int 
	val string
}

type listNode struct{
	n record
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

//adattare in base al valore
func printNode(l *listNode){
	fmt.Print(l.n.key, ": ", l.n.val, " ")
}

//adattare in base al tipo di valore
func newNode(i int, str string) *listNode{
	new := new(listNode)
	new.next = nil 
	new.n.key = i 
	new.n.val = str
	return new
}

//adattare in base al tipo di valore
func enqueue(q *queue, i int, str string){
	new := newNode(i,str)
	if q.head == nil{		//coda vuota
		q.head = new 
		q.tail = new
	}else{
		q.tail.next = new 
		q.tail = q.tail.next
	}
	return
}

func dequeue(q *queue) record{
	var n record
	if q.head != nil{
		n = q.head.n
		q.head = q.head.next
	}
	return n 
}

func isEmpty(q *queue) bool{
	return q.head == nil
}

func main(){
	//supponiamo di ordinare le macchine del gruppo di lavoro di un'azienda
	a := []record {record{key:4, val:"bmw"}, record{key:5, val:"mercedes"},record{key:8, val:"audi"},record{key:3, val:"ferrari"},record{key:3, val:"mclaren"},record{key:2, val:"lamborghini"},record{key:5, val:"porsche"},record{key:10, val:"volkswagen"},record{key:1, val:"renault"},record{key:8, val:"jeep"},record{key:4, val:"ford"},record{key:1, val:"honda"},record{key:9, val:"toyota"},}
	max := 10

	fmt.Println("Prima:")
	for i:=0; i < len(a); i++{
		fmt.Println(a[i].key, a[i].val)
	}
	fmt.Println()

	y := make([]*queue, 0)

	for i:=0; i < max+1; i++{
		y = append(y, createQueue())
	}

	for i:=0; i < len(a); i++{
		x := a[i].key
		enqueue(y[x], x, a[i].val)
	}

	j :=0
	for i:=0; i <= max; i++{
		for{
			if isEmpty(y[i]){
				break
			}
			a[j] = dequeue(y[i])
			j++
		}
	}
	fmt.Println("Dopo:")
	for i:=0; i < len(a); i++{
		fmt.Println(a[i].key, a[i].val)
	}
	
}