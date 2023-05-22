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

func dfs(g *grafo, v string, aux map[string]bool) {
	fmt.Printf("%s ", v)
	aux[v] = true
	for _, v2 := range g.adiacenze[v] {
		if !aux[v2] {
			dfs(g, v2, aux)
		}
	}
}

func main() {
	g := new(grafo)
	g.adiacenze = make(map[string][]string)
	leggiGrafo(g)
	//stampaGrafo(g)
	aux := make(map[string]bool)
	dfs(g, "A", aux)
}
