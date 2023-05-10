package main 
import(
	"fmt"
)

type graph struct{
	g [][]int 
	vertici int 
}

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

func newNode(n int) *listNode{
	new := new(listNode)
	new.next = nil 
	new.val = n
	return new
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

func isEmpty(q *queue) bool{
	return q.head == nil
}

func newGraph(n int) *graph{
	gph := new(graph)
	gph.g = make([][]int, n)
	gph.vertici = n
	return gph
}

func printGraph(gr *graph){
	for i:=0; i < len(gr.g); i++{
		fmt.Print(i, " -> ")
		for j :=0; j < len(gr.g[i]); j++{
			fmt.Print(gr.g[i][j], " -> ")
		}
		fmt.Println()
	}
	return 
}

func edgeInsert(gr *graph, vert int, edge int){
	gr.g[vert] = append(gr.g[vert], edge)
	return
}

func bfs_graph(g *graph, s int){
	q := createQueue()
	raggiunti := make([]int, g.vertici)
	raggiunti[s] = 1
	enqueue(q, s)
	for !isEmpty(q){
		u := dequeue(q)
		fmt.Printf("%d ", u.val)
		for i:=0; i < len(g.g[u.val]); i++{
			if raggiunti[g.g[u.val][i]] == 0{
				raggiunti[g.g[u.val][i]] = 1
				enqueue(q, g.g[u.val][i])
			}
		}
	}
	fmt.Println()
}

func main(){
	gr := newGraph(11)
	edgeInsert(gr,1,2)
	edgeInsert(gr,1,4)
	edgeInsert(gr,1,6)
	edgeInsert(gr,2,3)
	edgeInsert(gr,2,5)
	edgeInsert(gr,3,5)
	edgeInsert(gr,4,5)
	edgeInsert(gr,5,7)
	edgeInsert(gr,5,9)
	edgeInsert(gr,6,7)
	edgeInsert(gr,6,8)
	edgeInsert(gr,6,7)
	edgeInsert(gr,7,8)
	edgeInsert(gr,9,10)
	bfs_graph(gr, 1)
	//printGraph(gr)
}