/*
Si consideri un array A, non vuoto, di n elementi, contenente tutti gli interi appartenenti all’insieme
tranne uno. L’array A è ordinato in senso crescente.
Scrivere una funzione chiamata valoreMancante che, dato l’array A e la sua lunghezza n, determini il
valore mancante in A. Ad esempio, se A è il vettore [0;1;2;3;4;6;7], la funzione deve restituire il valore 5.
Oltre alla correttezza della funzione verrà valutata anche l’efficienza dell’algoritmo usato. Una soluzione
che impiega tempo lineare è poco interessante.
Note per la consegna. Si inserisca la funzione valoreMancante all’interno del file allegato di nome
es1-valoreMancante.c e si completi opportunamente la sua invocazione all’interno del main. Si noti che
il programma legge da standard input un intero n seguito da una sequenza di n interi ordinati (tutti gli interi
appartenenti all’insieme f0;1;2; : : : ;ng, tranne uno).*/

package main

import (
	"fmt"
	/*"strconv"
	"os"
	"bufio"	*/)

func valoreMancante(a []int, n int) int {
	val := 0
	if a[0] != 0 {
		val = 0
	}
	if a[n-1] != n {
		val = n
	}
	dx := n
	m := n
	sx := 0
	for sx < dx {
		m = (sx + dx) / 2
		if a[m] != m && a[m-1] == m-1 {
			val = m
			break
		} else if a[m] == m {
			sx = m + 1
		} else {
			dx = m
		}
	}
	return val
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4
	a[5] = 5
	a[6] = 6
	a[7] = 7
	a[8] = 8
	k := valoreMancante(a, n)
	fmt.Printf("Valore mancante: %d\n", k)

}
