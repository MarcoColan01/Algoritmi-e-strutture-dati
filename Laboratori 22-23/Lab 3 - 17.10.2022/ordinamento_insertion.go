/*
Scrivete una funzione che legga da standard input una sequenza di interi distinti terminati da 0, memorizzandoli in un vettore ordinato (valutate se è più opportuno usare un array o una slice): ogni volta che viene letto un nuovo intero, il vettore viene scorso fino a trovare l’esatta collocazione del numero, quindi si crea lo spazio per il nuovo numero spostando in avanti i numeri successivi già memorizzati. Questo algoritmo è utile per riempire un vettore mantenendolo ordinato ad ogni passo. 
*/

package main 

import (
	"fmt"
)

func stampaSlice(a []int){
	for i:=0; i < len(a); i++{
		fmt.Print(a[i], " ")
	}

	fmt.Println();
	return;
}

func insertionSort(a []int, k int) []int{
	if len(a) == 0{
		a = append(a, k)
		return a;
	}

	if len(a) == 1{
		if a[0] > k{
			a = append(a, k)
			a[0],a[1] = a[1], a[0]
		}else{
			a = append(a,k)
		}
		return a
	}else{
		a = append(a, k);
		for i:=1; i < len(a); i++{
			x:= a[i];
			j:= i-1;
			for{
				if !(j >= 0 && a[j] > x){
					break
				}
				a[j+1] = a[j];
				j = j -1;
			}
			a[j+1] = x;
		}
	}
	return a;
}

func main(){
	var a []int 
	n := 0;
	a = make([]int, 0);

	for{
		fmt.Scanf("%d", &n);
		//fmt.Println(n);
		if n == 0{
			break;
		}
		a = insertionSort(a, n);
	}
	stampaSlice(a);
}