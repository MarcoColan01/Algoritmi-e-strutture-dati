package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grafo struct {
	adiacenze map[string][]string
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
		g.adiacenze[k[0]] = k[1:]
	}
}

func stampaGrafo(g *grafo) {
	for i, k := range g.adiacenze {
		fmt.Printf("Nodo %s: %s\n", i, k)
	}
}

func bfs(g *grafo, v string, aux map[string]bool) {
	coda := []string{v}
	aux[v] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Printf("%s ", v)

		for _, v2 := range g.adiacenze[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
}

func main() {
	g := new(grafo)
	g.adiacenze = make(map[string][]string)
	leggiGrafo(g)
	//stampaGrafo(g)
	aux := make(map[string]bool)
	bfs(g, "A", aux)
}
