/*
AlgoFlix è un sistema di streaming online di film. Il sistema memorizza ogni giorno quante persone hanno visto un certo film, usando un vettore P di numeri interi: P[i] è il numero di persone che hanno visto il film nel giorno i. Vogliamo calcolare il massimo numero di giorni consecutivi in cui il numero di persone che hanno visto il film risulta non decrescente. In altre parole, vogliamo calcolare la lunghezza massima tra tutti i sottovettori composti da elementi contigui di P che risultino ordinati in senso non decrescente. Si scriva un algoritmo per risolvere questo problema, se ne determini il costo computazionale e lo si implementi in Go.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sottoVetMax(p []int, n int) int {
	dyn := make([]int, n)
	max := 1
	dyn[0] = 1
	for i := 1; i < len(p); i++ {
		if p[i] > p[i-1] {
			dyn[i] = dyn[i-1] + 1
		} else {
			dyn[i] = 1
		}
		if dyn[i] > max {
			max = dyn[i]
		}
	}
	return max
}

func main() {
	var n int
	var a []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		k := scanner.Text()
		if k == "" {
			break
		}
		if len(k) == 1 {
			j, _ := strconv.Atoi(k)
			n = j
			a = make([]int, n)
		} else {
			nums := strings.Split(k, " ")
			for i := 0; i < len(nums); i++ {
				n, _ := strconv.Atoi(string(nums[i]))
				a[i] = n
			}
		}
	}
	max := sottoVetMax(a, n)
	fmt.Println(max)
}
