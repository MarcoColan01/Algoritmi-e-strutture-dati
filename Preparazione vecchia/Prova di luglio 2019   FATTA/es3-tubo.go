/*
Disponiamo di un tubo metallico di lunghezza L. Da questo tubo possono essere ricavati segmenti più corti,
segando il tubo a partire da un’estremità. Abbiamo bisogno di n segmenti di tubo di lunghezza rispettivamente
`1; `2; : : : `n, ma in generale la somma delle lunghezze dei segmenti supera la lunghezza L del tubo di partenza,
quindi non è possibile ricavare tutti i n segmenti voluti.
Ad esempio, se si ha un tubo di lunghezza 24 e servono 6 segmenti di lunghezza rispettivamente 7 5 2 6 2
3, si possono ottenere al massimo 5 segmenti (ad esempio quelli di lunghezza 7 5 2 6 2). Se invece si ha un
tubo di lunghezza 10 e servono 7 segmenti di lunghezza rispettivamente 7 4 3 5 1 2 6 si possono ottenere al
massimo 4 segmenti (quelli di lunghezza 4 3 1 2).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func segmentiMax(a []int, l int) int {
	sort.Ints(a)
	sum := 0
	cont := 0
	i := 0
	for sum <= l && i < len(a) {
		sum += a[i]
		i++
		cont++
	}
	return cont
}

func main() {
	filename := os.Args[1]
	l := 0
	n := 0
	fmt.Scanf("%d %d", &l, &n)
	a := make([]int, n)
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	i := 0
	for scanner.Scan() {
		k, _ := strconv.Atoi(scanner.Text())
		a[i] = k
		i++
	}
	segMax := segmentiMax(a, l)
	fmt.Printf("Il numero massimo di segmenti ricavabili con un tubo di lunghezza %d è: %d\n", l, segMax)
}
