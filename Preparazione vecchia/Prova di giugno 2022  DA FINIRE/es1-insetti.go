/*
Un entomologo sta studiando lo sviluppo di una piccola colonia di insetti. Periodicamente rileva quanti individui sono nella colonia e ora sta preparando una relazione sull’andamento della popolazione nella colonia. L’entomologo innanzitutto calcola quante volte si è osservato un aumento di individui rispetto alla misurazione precedente. Considera poi finestre di m rilevazioni consecutive e calcola quante volte osserva un aumento della somma degli individui nella finestra, rispetto alla finestra precedente.
Ad esempio, i dati delle ultime 8 rilevazioni sono le seguenti: 5689, 6258, 4923, 3926, 5916, 6101, 5804, 5707, Il numero di rilevazioni per cui si è osservato un aumento è 3; se consideriamo finestre di 3 rilevazioni consecutive, il numero di volte in cui si è osservato un aumento in una finestra di 3 rilevazioni è invece 2. Le somme degli individue nelle 6 finestre di ampiezza 3 sono infatti:
5689+6258+4923 = 16870
6258+4923+3926 = 15107
4923+3926+5916 = 14765
3926+5916+6101 = 15943
5916+6101+5804 = 17821
6101+5804+5707 = 17612

Scrivete un programma che
	• legge un intero n che indica il numero di rilevazioni;
	• legge un numero m che indica l’ampiezza della finestra (m < n);
	• stampa il numero di volte in cui si è osservato un aumento rispetto alla misurazione precedente;
	• stampa il numero di volte in cui si è osservato un aumento in una finestra di m rilevazioni.
Stimate la complessità in tempo degli algoritmi usati, in funzione di n e m; stimate inoltre quante somme vengono eseguite.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func rileva(misure []int, m int) (int, int) {
	k1 := 0
	k2 := 0
	var ok bool
	for i := 0; i < len(misure)-1; i++ {
		if misure[i] < misure[i+1] {
			k1++
		}
	}

	for i := 0; i < len(misure)-m+1; i++ {
		ok = true
		for k := i; k < i+m-1; k++ {
			fmt.Print(misure[k], " ")
			if misure[k+1] < misure[k] {
				ok = false
				break
			}
		}
		if ok {
			k2++
		}
		fmt.Println()
	}
	return k1, k2
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	misure := make([]int, n)
	filename := os.Args[1]
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	i := 0
	for scanner.Scan() {
		k, _ := strconv.Atoi(scanner.Text())
		misure[i] = k
		i++
	}
	aumenti, aumentiFinestra := rileva(misure, m)
	fmt.Printf("Aumenti rilevati: %d\nAumenti rilevati su finestra di %d rilevazioni: %d\n", aumenti, m, aumentiFinestra)
}
