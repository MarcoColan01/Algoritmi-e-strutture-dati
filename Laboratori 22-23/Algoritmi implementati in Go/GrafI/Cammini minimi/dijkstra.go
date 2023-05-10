package main 
import(
	"fmt"
)

type edges struct{
	v int
	weigth int
}

type graph struct{
	g map[int][]edges
	vert int
}

func newGraph(n int) *graph{
	gr := new(graph)
	gr.g = make(map[int][]edges)
	gr.vert = n
	return gr
}

func edgeInsert(gr *graph, vert int, edge int, w int){
	if w >= 0{
		var e edges 
		e.v = edge 
		e.weigth = w
		gr.g[vert] = append(gr.g[vert], e)
	}
	return
}

func printGraph(gr *graph){
	for i,j := range gr.g{
		for k:=0; k < len(j);k++{
			fmt.Printf("Arco %d %d: costo %d\n", i, j[k].v, j[k].weigth)
		}
	}
	return 
}

//struttura dati: minHeap (quindi l'elemento con valore minore è quello con priorità maggiore e si trova nella radice, ossia il primo elemento dell'array)
type elems struct {
	key int 
	val int
}

type Pqueue struct{
	queue []elems
	size int
}

func newPriorityQueue() *Pqueue{
	pq := new(Pqueue)
	pq.queue = make([]elems, 0)
	return pq 
}

func printPriorityQueue(pq *Pqueue){
	for i:=0; i < pq.size; i++{
		fmt.Printf("%d: %d\n", pq.queue[i].key, pq.queue[i].val)
	}
	//fmt.Println()
}

func isEmpty(pq *Pqueue) bool{
	return pq.size == 0
}

func findMin(pq *Pqueue) (bool,int){
	find := false 
	i := 0
	if pq.size > 0{
		find = true
		i = pq.queue[0].val
	}else{
		i = -1
	}
	return find, i
}

func heapify(pq *Pqueue, size int, i int){
	if size > 1{	//trova l'elemento maggiore tra radice, figlio destro e sinistro
		largest := i
		l := 2*i +1
		r := 2*i +2
		if l < size && pq.queue[l].val < pq.queue[largest].val{
			largest = l 
		}
		if r < size && pq.queue[r].val < pq.queue[largest].val{
			largest = r
		}

		if largest != i{
			pq.queue[i], pq.queue[largest] = pq.queue[largest], pq.queue[i]
			heapify(pq, size, largest)
		}
	}
}

func deleteRoot(pq *Pqueue, num elems) elems{
	var min elems
	if pq.size > 0{
		i:=0
		for ; i < pq.size; i++{
			if num.key == pq.queue[i].key{
				break
			}
		}
		min = pq.queue[i]
	//fmt.Println("Elemento", pq.queue[i].key, "rimosso dalla testa della coda")

	pq.queue[i], pq.queue[pq.size -1] = pq.queue[pq.size -1], pq.queue[i]
	pq.size--
	for i = pq.size/2 -1; i >=0; i--{
		heapify(pq, pq.size, i)
	}
	}
	return min
}

func insert(pq *Pqueue, key int, val int) {
	var e elems
	e.key = key
	e.val = val
	if (pq.size == 0) {
	  pq.queue = append(pq.queue, e)
	  pq.size++
	} else {
	  pq.queue = append(pq.queue, e)
	  pq.size++
	  for i := pq.size/2 -1; i >= 0; i--{
		heapify(pq, pq.size, i)
	  }
	}
  }

func changeKey(pq *Pqueue, toChg int, newVal int){
	i :=0
	for ; i < pq.size; i++{
		if pq.queue[i].key == toChg{
			break
		}
	}

	if i == pq.size{
		return
	}else{
		pq.queue[i].val = newVal
		for i := pq.size/2 -1; i >= 0; i--{
			heapify(pq, pq.size, i)
		}
	}
}


func dijkstra(g *graph, s int){
	d := make([]int, g.vert)
	//pred := make([]int, g.vert)

	for i:=0; i < g.vert; i++{
		d[i] = 9999
	}
	d[s] = 0

	pq := newPriorityQueue()
	for i:=0; i < g.vert; i++{
		insert(pq, i, d[i])
	}

	//printPriorityQueue(pq)

	for	!isEmpty(pq){
		x := deleteRoot(pq, pq.queue[0])
			//fmt.Println(vertice)
			for j := 0; j < len(g.g[x.key]); j++{
				if d[x.key] + g.g[x.key][j].weigth < d[g.g[x.key][j].v]{
					d[g.g[x.key][j].v] = d[x.key] + g.g[x.key][j].weigth 
					changeKey(pq, g.g[x.key][j].v, d[g.g[x.key][j].v])
				}
			}
		}
	for i:=0; i < len(d); i++{
		fmt.Printf("Cammino minimo dal vertice %d al vertice %d: %d\n", s, i, d[i])
	}
	//fmt.Println(pred)
}

func main(){
	g := newGraph(7)
	edgeInsert(g,0,1,7)
	edgeInsert(g,0,2,6)
	edgeInsert(g,2,0,7)
	edgeInsert(g,1,3,7)
	edgeInsert(g,3,2,2)
	edgeInsert(g,2,4,4)
	edgeInsert(g,4,3,3)
	edgeInsert(g,3,6,2)
	edgeInsert(g,6,4,4)
	edgeInsert(g,5,6,4)
	edgeInsert(g,4,5,2)
	//printGraph(g)
	/*edgeInsert(g,0,1,2)
	edgeInsert(g,0,2,6)
	edgeInsert(g,1,3,5)
	edgeInsert(g,2,3,8)
	edgeInsert(g,3,5,15)
	edgeInsert(g,4,5,6)
	edgeInsert(g,4,6,2)
	edgeInsert(g,5,6,6)
	edgeInsert(g,3,4,10)*/
	//printGraph(g)
	dijkstra(g, 0)
}