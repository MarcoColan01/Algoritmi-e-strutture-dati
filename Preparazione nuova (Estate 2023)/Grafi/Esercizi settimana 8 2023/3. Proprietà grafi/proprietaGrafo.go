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

/*
gen (p) genera un grafo casuale, a partire dalla probabilità p compresa tra 0 e 1 (inclusi). Il modello matematico di riferimento è il seguente: si considerano tutti i possibili archi includendoli nel grafo con probabilità p. Più esplicitamente, per ogni possibile coppia di
vertici, si genera un numero reale compreso tra 0 e 1; se questo è minore di p si inserisce l’arco, altrimenti non lo si inserisce. */

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

/*
degree (v) calcola il grado del vertice v.
Si ricorda che il grado di un vertice è definito come il numero di vertici ad esso adiacenti.
*/
func degree(g *grafo, v int) {
	fmt.Printf("Grado del vertice %d: %d\n", v, len(g.adiacenze[v]))
}

func bfs(g *grafo, v int, w int) bool {
	trovato := false
	coda := []int{v}
	aux := make(map[int]bool)
	aux[v] = true
	for len(coda) > 0 && trovato == false {
		app := coda[0]
		coda = coda[1:]
		if app == w {
			trovato = true
			break
		}

		for _, v2 := range g.adiacenze[app] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
	return trovato

}

/*
path (v, w) testa l’esistenza di un cammino semplice che collega i vertici v e w. Si ricorda che un cammino si dice semplice quando attraversa ogni vertice al più una volta.
*/
func path(g *grafo, v int, w int) bool {
	return bfs(g, v, w)
}

func ccc(g *grafo) {
	cont := 0

}

func main() {
	g := gen(0.6)
	stampaGrafo(g)
	if path(g, 0, 4) {
		fmt.Println("Cammino esistente")
	} else {
		fmt.Println("Cammino NON esistente")
	}
}
