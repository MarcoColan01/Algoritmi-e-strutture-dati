
package main 
import (
	"fmt"
)

//struttura dati: maxHeap (quindi l'elemento con valore maggiore è quello con priorità maggiore e si trova nella radice, ossia il primo elemento dell'array)

type Pqueue struct{
	queue []int
	size int
}

func newPriorityQueue() *Pqueue{
	pq := new(Pqueue)
	pq.queue = make([]int, 0)
	return pq 
}

func printPriorityQueue(pq *Pqueue){
	for i:=0; i < pq.size; i++{
		fmt.Print(pq.queue[i], " ")
	}
	fmt.Println()
}


func isEmpty(pq *Pqueue) bool{
	return pq.size == 0
}

func findMax(pq *Pqueue) (bool,int){
	find := false 
	i := 0
	if pq.size > 0{
		find = true
		i = pq.queue[0]
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
		if l < size && pq.queue[l] > pq.queue[largest]{
			largest = l 
		}
		if r < size && pq.queue[r] > pq.queue[largest]{
			largest = r
		}

		if largest != i{
			pq.queue[i], pq.queue[largest] = pq.queue[largest], pq.queue[i]
			heapify(pq, size, largest)
		}
	}
}

func deleteRoot(pq *Pqueue, num int){
	if pq.size > 0{
		i:=0
		for ; i < pq.size; i++{
			if num == pq.queue[i]{
				break
			}
		}
	fmt.Println("Elemento", pq.queue[i], "rimosso dalla testa della coda")

	pq.queue[i], pq.queue[pq.size -1] = pq.queue[pq.size -1], pq.queue[i]
	pq.size--
	for i = pq.size/2 -1; i >=0; i--{
		heapify(pq, pq.size, i)
	}
	}
}


func insert(pq *Pqueue, new int) {
	if (pq.size == 0) {
	  pq.queue = append(pq.queue, new)
	  pq.size++
	} else {
	  pq.queue = append(pq.queue, new)
	  pq.size++
	  for i := pq.size/2 -1; i >= 0; i--{
		heapify(pq, pq.size, i)
	  }
	}
  }

  func delete(pq *Pqueue, canc int){
	i:=0
	for ; i < pq.size; i++{
		if pq.queue[i] == canc{
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
		if pq.queue[i] == toChg{
			break
		}
	}

	if i == pq.size{
		return
	}else{
		pq.queue[i] = newVal
		for i := pq.size/2 -1; i >= 0; i--{
			heapify(pq, pq.size, i)
		}
	}
}



func main(){
	pq := newPriorityQueue()
	var ch string 
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
			case "chg":
				changeKey(pq, n, new)
			case "f":
				exit = true
				break
		}
	}
}