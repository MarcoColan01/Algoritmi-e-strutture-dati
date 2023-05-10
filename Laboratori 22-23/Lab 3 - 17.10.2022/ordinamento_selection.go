/*
Scrivete una funzione che riceva una slice di interi e la ordini usando l’algoritmo SelectionSort:
alla fine dell’esecuzione, la slice originaria passata come argomento dovrà risultare ordinata. Può essere utile restituire la stessa slice (ordinata), ad esempio per poterla passare come argomento ad altre funzioni, come in fmt.Println(selectionSort(v)) .

Versione iterativa
La funzione selectionSortIter( int a[]) ripete la seguente operazioni tante volte quanto è lunga la slice da ordinare: per ogni prefisso di lunghezza n (con n inizialmente pari alla lunghezza della slice) cerca nel prefisso l’elemento massimo e lo scambia con quello nell’ultima posizione del prefisso.

Versione ricorsiva
Scrivete una versione ricorsiva dello stesso algoritmo: la funzione selectionSortRec
deve funzionare come segue:
• innanzitutto cerca nel vettore l’elemento massimo e lo sposta nell’ultima posizione;
• poi richiama se stessa ricorsivamente per ordinare i primi n − 1 elementi del vettore.
La base della ricorsione è data dai vettori di lunghezza 0 o 1, che sono sempre ordinati.
*/

package main 
import (
	"fmt"
)

func stampaSlice(a []int){
	for i :=0; i < len(a); i++{
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	return
}

func selectionSortIter(vet []int){
	for k:=0; k < len(vet); k++{
		m := k;
		for j := k+1; j < len(vet); j++{
			if vet[j] < vet[m]{
				m = j;
			}
		}
		vet[m], vet[k] = vet[k], vet[m]
	}
	return 
}

func selectionSortRic(vet []int){
	if len(vet) == 0 || len(vet) == 1{
		return
	}
	m := 0;
	for i:=0; i < len(vet); i++{
		if vet[i] > vet[m] {
			m = i
		}
		vet[m], vet[len(vet)-1] = vet[len(vet)-1], vet[m]
	}
	selectionSortRic(vet[:len(vet)-2])
}

func main(){
	a := []int {10, -5, 11, 23, 45, 9, 21, 2, 7, 18, 29, 1, 33, 100, 27, 31, 9, 87, 75, 65, 12, 99, 34, 67, 8, 4, 37, 60}

	fmt.Print("Prima: ") 
	stampaSlice(a)
	//selectionSortIter(a)
	selectionSortRic(a)
	fmt.Print("Dopo: ") 
	stampaSlice(a)

}