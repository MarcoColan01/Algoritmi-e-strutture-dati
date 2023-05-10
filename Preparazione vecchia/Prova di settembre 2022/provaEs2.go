package main

import (
	"fmt"
	"math"
)

type graph struct {
	matrix  [][]int
	vertici int
}

func newGraph(n int) *graph {
	g := new(graph)
	g.vertici = n
	g.matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g.matrix[i] = append(g.matrix[i], 0)
		}
	}
	return g
}

func algo(matrix [][]int, n int, h int, s int) int {
	current_node := h
	c := 0
	m := 0
	m_i := 0
	v := make([]int, n)

	for current_node != s {
		if v[current_node] > 0 {
			fmt.Printf("-1\n")
			return 0
		}
		v[current_node] = 1
		m = math.MaxInt
		m_i = -1
		for i := 0; i < n; i++ {
			if m > matrix[current_node][i] && matrix[current_node][i] != 0 {
				m_i = i
				m = matrix[current_node][i]
			}
		}
		if m == math.MaxInt {
			fmt.Printf("-1\n")
			return 0
		}
		current_node = m_i
		c++
	}
	fmt.Printf("%d\n", c)
	return 0
}

func printGraph(g *graph) {
	for i := 0; i < g.vertici; i++ {
		fmt.Println(g.matrix[i])
	}
}

func main() {
	g := newGraph(5)
	g.matrix[1][4] = 5
	g.matrix[2][4] = 3
	g.matrix[0][1] = 7
	g.matrix[4][3] = 10
	g.matrix[4][1] = 1
	g.matrix[3][2] = 12
	g.matrix[1][3] = 4
	g.matrix[1][2] = 8
	printGraph(g)

	k := algo(g.matrix, g.vertici, 2, 4)
	fmt.Println(k, "diocane")

}
