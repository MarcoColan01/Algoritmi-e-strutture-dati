/*
Considerate nuovamente la situazione e il problema descritti nel riquadro, e l’algoritmo che
avete progetato per l’esercizio precedente.
Implementate l’algoritmo con un programma Go. Il formato di input deve essere il seguente:
	• La prima riga dell’input è composta dai quattro numeri interi N, M, H e S definiti come
	nel riquadro;
	• Ognuna delle successive M righe descrive una galleria e contiene tre numeri interi A, B
	e L separati da uno spazio: i primi due rappresentano gli svincoli collegati dalla galleria
	mentre il terzo rappresenta il suo grado di luminosità.
L’output dovrà contenere un solo numero: il numero di gallerie che Harmony percorrerà per
raggiungere la casa di Sarah, oppure -1.
*/


package main 
import(
	"fmt"
	"os"
	"strconv"
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


func isEmpty(pq *Pqueue) bool{
	return pq.size == 0
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

func printPriorityQueue(pq *Pqueue){
	for i:=0; i < pq.size; i++{
		fmt.Printf("%d: %d\n", pq.queue[i].key, pq.queue[i].val)
	}
}


func dijkstra(g *graph, h int, s int) int{
	d := make([]int, g.vert)

	for i:=0; i < g.vert; i++{
		d[i] = 9999
	}
	d[h] = 0

	pq := newPriorityQueue()
	for i:=0; i < g.vert; i++{
		insert(pq, i, d[i])
	}
	printPriorityQueue(pq)

	cont := 0
	find := false
	for	!isEmpty(pq){
		x := deleteRoot(pq, pq.queue[0])
		fmt.Println("orcodio")
		cont++
			for j := 0; j < len(g.g[x.key]); j++{
				if g.g[x.key][j].v == s{
					find = true
					break
				}
				if d[x.key] + g.g[x.key][j].weigth < d[g.g[x.key][j].v]{
					d[g.g[x.key][j].v] = d[x.key] + g.g[x.key][j].weigth 
					changeKey(pq, g.g[x.key][j].v, d[g.g[x.key][j].v])
				}
			}
			if find{
				break
			}
		}
		if find{
			return cont 
		}else{
			return -1
		}
}

func reachSarah(g *graph, h int, s int) int{
	k := dijkstra(g, h, s)
	return k
}


func main(){
	n,_ := strconv.Atoi(os.Args[1])
	h,_ := strconv.Atoi(os.Args[2])
	s,_ := strconv.Atoi(os.Args[3])
	fmt.Printf("n %d h %d s %d\n", n, h, s)
	g := newGraph(n)
	//esempio che non funziona
	edgeInsert(g, 2,0,2)
	edgeInsert(g, 1,2,1)
	//esempio che funziona
	/*edgeInsert(g, 0,1,5)
	edgeInsert(g, 1,2,1)
	edgeInsert(g, 2,3,3)
	edgeInsert(g, 3,4,2)
	edgeInsert(g, 4,0,6)
	edgeInsert(g, 0,3,4)*/
	k := reachSarah(g,h,s)
	if k != -1{
		fmt.Printf("harmony in %d può raggiungere sarah in %d percorrendo %d gallerie\n", h,s,k)
	}else{
		fmt.Printf("harmony in %d non può raggiungere sarah in %d\n", h, s)
	}

}

