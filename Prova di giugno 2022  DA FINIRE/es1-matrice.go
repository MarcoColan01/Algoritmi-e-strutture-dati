/*
Scrivere un programma che:
	-legge da input due interi n e m (rispettivamente, righe e colonne della matrice) e una matrice di bit n x m
	-memorizza la matrice
	-calcola e stampa il numero di bit pari a 1 nella matrice
	-interpreta ogni riga della matrice come numero binario, lo converte e lo stampa
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func conversioni(matrix [][]int, riga int) {
	cont := 0
	//exp := 0
	valori := make([]int, 0)
	app := 0
	for i := 0; i < riga; i++ {
		//fmt.Println(matrix[i])
		for j := 0; j < len(matrix[i]); j++ {
			//fmt.Println(len(matrix[i]))
			if matrix[i][j] == 1 {
				cont++
				//fmt.Print(math.Exp2(float64(j)), " ")
				app = app + int(math.Exp2(float64(len(matrix[i])-j-1)))
			}
		}
		//fmt.Println()
		valori = append(valori, app)
		app = 0
	}
	fmt.Printf("%d %d\n", cont, valori)
}

func main() {
	var n, m int
	i := 0
	var riga string
	fmt.Scanf("%d %d", &n, &m)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	//fmt.Println(matrix)
	for {
		_, err := fmt.Scan(&riga)
		if err != nil {
			break
		}
		for j := 0; j < len(riga); j++ {
			k, _ := strconv.Atoi(string(riga[j]))
			//fmt.Println(k)
			matrix[i][j] = k
		}
		i++
	}
	conversioni(matrix, n)
}
