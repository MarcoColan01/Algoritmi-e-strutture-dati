package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func percorsi(mat [][]int, n int, m int) int {
	dyn := make([][]int, n)
	for i := 0; i < n; i++ {
		dyn[i] = make([]int, m)
	}

	//Casi base
	for j := 0; j < m; j++ {
		dyn[n-1][j] = 1
	}
	for i := 0; i < n; i++ {
		dyn[i][m-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			if mat[i][j] == 1 {
				dyn[i][j] = 0
			} else {
				dyn[i][j] = dyn[i+1][j] + dyn[i][j+1]
			}
		}
	}

	return dyn[0][0]
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
	k := percorsi(mat, n, m)
	fmt.Println(k)
}
