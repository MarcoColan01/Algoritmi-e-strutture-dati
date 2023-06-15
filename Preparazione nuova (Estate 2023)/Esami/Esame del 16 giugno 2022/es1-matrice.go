package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertiMatrice(mat [][]int, n int, m int) {
	somma := 0.0
	for i := 0; i < n; i++ {
		somma = 0
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				somma += math.Pow(2, float64(m-j-1))
			}
		}
		fmt.Println(somma)
	}
}

func main() {
	var mat [][]int
	cont := 0
	var n, m, j int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), " ")
		if k[0] == "" {
			break
		} else if len(k) == 2 {
			j = 0
			a, _ := strconv.Atoi(k[0])
			b, _ := strconv.Atoi(k[1])
			n = a
			m = b
			mat = make([][]int, n)
			for i := 0; i < n; i++ {
				mat[i] = make([]int, m)
			}
		} else {
			app := strings.Split(k[0], "")
			for i := 0; i < len(app); i++ {
				n, _ := strconv.Atoi(app[i])
				if n == 1 {
					cont++
				}
				mat[j][i] = n
			}
			j++
		}
	}

	fmt.Println(cont)
	convertiMatrice(mat, n, m)
}
