package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	w, peso int
}

type grafo struct {
	adiacenze map[int][]*edge
}

func addEdge(g *grafo, x int, y int, peso int) {
	app := &edge{y, peso}
	g.adiacenze[x] = append(g.adiacenze[x], app)
}

func stampaGrafo(g *grafo) {
	for i, k := range g.adiacenze {
		fmt.Printf("Vertice %d: ", i)
		for j := 0; j < len(k); j++ {
			fmt.Printf("%d peso %d  ", k[j].w, k[j].peso)
		}
		fmt.Println()
	}
}

func raggiungiSarah(g *grafo, h int, s int) int {
	cont := 0
	trovato := false
	coda := []int{h}
	aux := make(map[int]bool)
	aux[h] = true

	for len(coda) > 0 {
		k := coda[0]
		if k == s {
			trovato = true
			break
		}
		coda = coda[1:]
		minPeso, minV := math.MaxInt, 0
		for _, k2 := range g.adiacenze[k] {
			if !aux[k2.w] {
				aux[k2.w] = true
			}
			if k2.peso < minPeso {
				minPeso = k2.peso
				minV = k2.w
			}
		}
		if minV != 0 {
			coda = append(coda, minV)
			cont++
		}
	}
	if !trovato {
		return -1
	} else {
		return cont
	}
}

func main() {
	g := new(grafo)
	var h, s int
	g.adiacenze = make(map[int][]*edge)

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
		if len(k) == 4 {
			h, _ = strconv.Atoi(k[2])
			s, _ = strconv.Atoi(k[3])
		} else if len(k) == 3 {
			x, _ := strconv.Atoi(k[0])
			y, _ := strconv.Atoi(k[1])
			peso, _ := strconv.Atoi(k[2])
			addEdge(g, x, y, peso)
		}
	}
	//stampaGrafo(g)
	k := raggiungiSarah(g, h, s)
	fmt.Println(k)
}
