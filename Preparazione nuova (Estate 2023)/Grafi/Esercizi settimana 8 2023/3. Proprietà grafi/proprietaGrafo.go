package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type grafo struct {
	adiacenze map[int][]int
	n         int
}

func leggiGrafo(g *grafo) {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(k[0])
		for i := 1; i < len(k); i++ {
			y, _ := strconv.Atoi(k[i])
			g.adiacenze[x] = append(g.adiacenze[x], y)
		}
	}
}

func stampaGrafo(g *grafo) {
	for i, k := range g.adiacenze {
		fmt.Printf("Nodo %d: %d\n", i, k)
	}
}

func addEdge(g *grafo, x int, y int) {
	g.adiacenze[x] = append(g.adiacenze[x], y)
}

func gen(p float64) *grafo {
	rand.Seed(time.Now().UnixNano())
	g := new(grafo)
	g.n = 7
	g.adiacenze = make(map[int][]int)
	for i := 0; i < g.n; i++ {
		for j := i + 1; j < g.n; j++ {
			k := rand.Float64()
			if k < p {
				addEdge(g, i, j)
			}
		}
	}
	return g
}

func main() {
	g := gen(0.6)
	stampaGrafo(g)

}
