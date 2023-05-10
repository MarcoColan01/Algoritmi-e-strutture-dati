

package main 
import (
	"fmt"
)

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

func deleteRoot(pq *Pqueue, num elems){
	if pq.size > 0{
		i:=0
		for ; i < pq.size; i++{
			if num.key == pq.queue[i].key{
				break
			}
		}
	fmt.Println("Elemento", pq.queue[i].key, "rimosso dalla testa della coda")

	pq.queue[i], pq.queue[pq.size -1] = pq.queue[pq.size -1], pq.queue[i]
	pq.size--
	for i = pq.size/2 -1; i >=0; i--{
		heapify(pq, pq.size, i)
	}
	}
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

  func delete(pq *Pqueue, canc int){
	i:=0
	for ; i < pq.size; i++{
		if pq.queue[i].key == canc{
			break
		}
	}

	if i == pq.size{
		return
	}else{
		pq.queue[i] = pq.queue[pq.size-1]
		pq.queue = append(pq.queue[:pq.size-1], pq.queue[pq.size:]...)
		pq.size--
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

func main(){
	pq := newPriorityQueue()
	insert(pq, 3, 5)
	insert(pq, 4, 51)
	insert(pq, 7, 2)
	insert(pq, 30, 1)
	insert(pq, 1, 16)
	insert(pq, 6, 23)
	insert(pq, 9, 31)
	insert(pq, 10, 7)
	printPriorityQueue(pq)
	changeKey(pq, 9, 12)
	changeKey(pq, 30, 19)
	printPriorityQueue(pq)
	deleteRoot(pq, pq.queue[0])
	printPriorityQueue(pq)
	delete(pq, 4)
	/*var ch string 
	var n int
	var new int

	exit := false
	for !exit{
		fmt.Scanf("%s %d %d", &ch, &n, &new)
		switch ch{
			case "+":
				insert(pq, n)
			case "-":
				delete(pq, n)
			case "p":
				printPriorityQueue(pq)
			case "?":
				if isEmpty(pq){
					fmt.Println("La coda è vuota")
				}else{
					fmt.Println("La coda NON è vuota")
				}
			case "1":
				deleteRoot(pq, pq.queue[0])
			case "first":
				find,i := findMin(pq)
				if find{
					fmt.Println("Elemento con priorità maggiore:", i)
				}
			case "chg":
				changeKey(pq, n, new)
			case "f":
				exit = true
				break
		}
	}*/
}
