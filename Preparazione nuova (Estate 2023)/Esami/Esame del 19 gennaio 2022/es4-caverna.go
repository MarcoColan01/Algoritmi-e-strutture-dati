package main

import "fmt"

func min(n1 int, n2 int, n3 int, n4 int) int {
	vet := make([]int, 4)
	vet[0] = n1
	vet[1] = n2
	vet[2] = n3
	vet[3] = n4

	min := vet[0]
	for i := 1; i < len(vet); i++ {
		if vet[i] < min && vet[i] != 0 {
			min = vet[i]
		}
	}
	fmt.Println(min)
	return min

}

func m1(M [][]int, n int) int {
	//var C [][]int
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	C[0] = M[0]
	C[0][0] = 0

	fmt.Println(C)

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				C[i][j] = C[i-1][j] + M[i][j]
			} else if j == n-1 || i == n-1 {
				if C[i-1][j] < C[i][j-1] {
					C[i][j] = C[i-1][j] + M[i][j]
				} else {
					C[i][j] = C[i][j-1] + M[i][j]
				}
			} else {
				n := min(C[i-1][j], C[i+1][j], C[i][j-1], C[i][j+1])
				C[i][j] = M[i][j] + n
			}
		}
	}
	fmt.Println(C)

	return C[n-1][n-1]
}

func m2(M [][]int, n int) int {
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	C[0] = M[0]
	C[0][0] = 0

	fmt.Println(C)

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				C[i][j] = M[i][j] + C[i-1][j]
			} else {
				if C[i][j-1] < C[i-1][j] {
					C[i][j] = M[i][j] + C[i][j-1]
				} else {
					C[i][j] = M[i][j] + C[i-1][j]
				}
			}
		}
	}
	fmt.Println(C)

	return C[n-1][n-1]
}

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	M := [][]int{{7, 1, 8, 4, 4}, {9, 1, 8, 4, 4}, {1, 1, 8, 4, 4}, {1, 9, 8, 4, 4}, {1, 1, 1, 1, 1}}
	/*r1 := m1(M, n)
	fmt.Println(r1)*/
	r2 := m2(M, n)
	fmt.Println(r2)

}
