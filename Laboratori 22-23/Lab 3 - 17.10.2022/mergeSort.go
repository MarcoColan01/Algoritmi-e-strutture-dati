/*
Scrivete una funzione che implementa l’algoritmo ricorsivo mergeSort. La funzione:
1. divide il vettore in due sotto-vettori di dimensione circa uguale;
2. ordina il sotto-vettore di sinistra richiamando se stessa;
3. ordina il sotto-vettore di destra richiamando se stessa;
4. integra (merge) i due vettori in un vettore ordinato.
La base della ricorsione è, anche qui, data dai vattori di lunghezza 0 o 1, che sono sempre
ordinati.
La parte di integrazione (merge) di due vettore ordinati a1 e a2 funziona con un vettore di
supporto: si scorrono entrambi i vettori da sinistra a destra usando due indicatori i1 e i2 rispetti-
vamente; ad ogni passo si confronta a1[i1] con a2[i2] e si sceglie l’elemento più piccolo, lo si
copia nel vettore di supporto (nella prima posizione libera) e si incrementa l’indicatore relativo
ad esso. Quando i1 esce da a1 oppure i2 esce da a2 , la parte rimanente dell’altro vettore viene
copiata nel vettore di supporto. Alla fine si copia il contenuto del vettore di supporto nel vettore
originale.
*/
package main 
import(
	"fmt"
)

func merge(vet []int, i int, m int, f int, x []int) []int{
	i1 := i;
	i2 := m;
	k := 0;

	for{
		if !(i1 < m && i2 < f){
			break;
		}

		if vet[i1] <= vet[i2]{
			x[k] = vet[i1];
			i1++;
		}else{
			x[k] = vet[i2];
			i2++;
		}
		k++;
	}

	if i1 < m{
		for j:=i1; j < m; j++{
			x[k] = vet[j];
			k++;
		}
	}else{
		for j:=i2; j < f; j++{
			x[k] = vet[j];
			k++;
		}
	}

	for k:=0; k < f-i; k++{
		vet[i+k] = x[k];
	}

	return vet;
}

func stampaSlice(a []int){
	for i :=0; i < len(a); i++{
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	return
}

func mergeSort(vet []int, i int, f int, x []int) []int{
	if f-i > 1{
		m:= (i+f)/2;
		mergeSort(vet, i, m, x);
		mergeSort(vet, m, f, x);
		vet = merge(vet, i, m, f, x);
	}

	return vet;
}

func main(){
	vet := []int {10, -5, 11, 23, 45, 9, 21, 2, 7, 18, 29, 1, 33, 100, 27, 31, 9, 87, 75, 65, 12, 99, 34, 67, 8, 4, 37, 60}
	fmt.Print("Prima: ")
	stampaSlice(vet);

	var x []int;
	x = make([]int, len(vet));
	vet = mergeSort(vet, 0, len(vet), x);

	fmt.Print("Dopo: ")
	stampaSlice(vet);
}