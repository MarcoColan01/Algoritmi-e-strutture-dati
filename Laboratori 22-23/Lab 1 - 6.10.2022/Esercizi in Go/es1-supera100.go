/*
1 Supera 100
Scrivere un programma supera100.go che legge da standard input una sequenza di interi positivi terminata da -1 e stampa il primo numero che supera 100, se presente; altrimenti stampa “nessun numero maggiore di 100”.
*/

package main

import (
	"fmt";
)

func main(){
	var sequenza []int;
	var n int;
	sequenza = make([]int, 0);
	
	fmt.Println("Inserisci una sequenza terminata da -1");
	for{
		fmt.Scanf("%d", &n);
		if n == -1{
			break;
		}
		sequenza = append(sequenza, n);
	}

	trovato := false;
	fmt.Println("\n");
	for i:= 0; i < len(sequenza); i++{
		if sequenza[i] > 100{
			fmt.Println("Il primo elemento della seuqenza che supera 100 è: ", sequenza[i], "\n");
			trovato = true;
			break;
		}
	}
	if !trovato{
		fmt.Println("Nessun numero supera 100\n");
	}
}