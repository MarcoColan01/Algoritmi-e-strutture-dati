/*
4 Elaborazione di serie di interi
Scrivete le tre funzioni specificate sotto.

1. Scrivere una funzione stranoProdotto(numeri []int) int che, data come parametro una slice di interi, trovi quelli che sono maggiori di 7 e multipli di 4 e ne restituisca il prodotto. Ad esempio, se la slice contiene i numeri 12, 3, 4, 8, 9, 2, la funzione dovrà restituire il numero 96 (pari al prodotto di 12 per 8).

2. Scrivere una funzione pariDispari(numeri []int) che, data come parametro una slice di interi, stampi, per ciascun numero, se è pari o dispari.

3. Scrivere una funzione minDispari(numeri []int) int che, data una slice di interi, restituisca il più piccolo numero dispari (la slice può contenere sia numeri positivi che negativi).
*/

package main 

import (
	"fmt";
)

func main(){
	var numeri []int;
	var n int;
	numeri = make([]int, 0);

	for{
		fmt.Scan(&n);
		if n == 0{
			break;
		}

		numeri = append(numeri, n);
	}

	fmt.Println("Risultato di stranoProdotto: ", stranoProdotto(numeri));
	pariDispari(numeri);
	fmt.Println("Risultato di minDispari:", minDispari(numeri));
}

func minDispari(numeri []int) int{
	min := 92233720368547758;

	for i:= 0; i < len(numeri); i++ {
		if numeri[i] %2 != 0 && numeri[i] < min{
			min = numeri[i];
		}
	}

	return min;
}

func pariDispari(numeri []int){
	for i:= 0; i < len(numeri); i++{
		if numeri[i] %2 == 0{
			fmt.Println(numeri[i], "è pari");
		}else{
			fmt.Println(numeri[i], "è dispari");
		}
	}
	return;
}

func stranoProdotto(numeri []int)int{
	prod := 1;
	for i:= 0; i < len(numeri); i++{
		if numeri[i] > 7 && numeri[i] %4 == 0{
			prod *= numeri[i];
		}
	}
	return prod;
}