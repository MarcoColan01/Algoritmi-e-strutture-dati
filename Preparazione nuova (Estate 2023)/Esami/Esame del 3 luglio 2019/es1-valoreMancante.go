package main

import "fmt"

func valoreMancante(A []int, n int) int {
	dx := 0
	sx := n - 1
	for dx <= sx {
		m := (sx + dx) / 2
		if A[m]-A[dx] > m-dx {
			sx = m - 1
		} else {
			dx = m + 1
		}
	}
	return A[dx] - 1

}

func main() {
	A := []int{0, 1, 2, 3, 4, 6, 7}
	n := valoreMancante(A, len(A)+1)
	fmt.Println(n)
}
