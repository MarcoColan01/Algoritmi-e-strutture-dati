package main

import (
	"fmt"
	"math"
)

func algo(matrix [][]int, n int, h int, s int) int {
	curr_node := h
	c := 0
	var m, m_i int

	v := make([]int, n)
	for curr_node != s {
		if v[curr_node] > 0 {
			fmt.Printf("-1\n")
			return 0
		}
		v[curr_node] = 1
		m = math.MaxInt
		m_i = -1
		for i := 1; i < n; i++ {
			if m > matrix[curr_node][i] && matrix[curr_node][i] != 0 {
				m_i = i
				m = matrix[curr_node][i]
			}
		}
		if m == math.MaxInt {
			fmt.Printf("-1\n")
			return 0
		}
		curr_node = m_i
		c++
	}
	fmt.Printf("%d\n", c)
	return 0
}

func main() {
	a := make([][]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = make([]int, 5)
	}
	a[1][2] = 1
	a[2][3] = 3
	a[3][4] = 2
	a[4][1] = 3
	a[2][4] = 1
	a[3][0] = 4

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	n := algo(a, 5, 1, 4)
	fmt.Println(n)
}
