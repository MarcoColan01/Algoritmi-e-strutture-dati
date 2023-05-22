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

func bfs1(g *grafo, v int, w int) bool {
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

func bfs2(g *grafo, v int, aux map[int]bool) {
	coda := []int{v}
	//aux := make(map[int]bool)
	aux[v] = true
	for len(coda) > 0 {
		app := coda[0]
		coda = coda[1:]

		for _, v2 := range g.adiacenze[app] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}

}

/*
path (v, w) testa l’esistenza di un cammino semplice che collega i vertici v e w. Si ricorda che un cammino si dice semplice quando attraversa ogni vertice al più una volta.
*/
func path(g *grafo, v int, w int) bool {
	return bfs1(g, v, w)
}

/*
ccc conta il numero di componenti connesse di un grafo (non orientato). Si ricorda che si chiama componente connessa di un grafo ogni insieme massimale di vertici connessi tra loro da un cammino
*/
func ccc(g *grafo) {
	cont := 0
	aux := make(map[int]bool)
	for i, _ := range g.adiacenze {
		if !aux[i] {
			bfs2(g, i, aux)
			fmt.Println(aux)
			cont++
		}
	}
	fmt.Printf("Numero componenti connesse nel grafo: %d\n", cont)

}

/*
5. cc (v) stampa l’elenco dei vertici della componente connessa contenente v;
*/

func cc(g *grafo, v int) {
	aux := make(map[int]bool)
	coda := []int{v}
	aux[v] = true

	for len(coda) > 0 {
		k := coda[0]
		coda = coda[1:]
		for _, y := range g.adiacenze[k] {
			if !aux[y] {
				aux[y] = true
				coda = append(coda, y)
				fmt.Printf("%d ", v)
			}
		}
	}

}

/*
6. span (v) calcola uno spanning tree con radice v e lo stampa nella rappresentazione "a sommario".
Per ottenere l'albero di altezza minima, eseguo una visita in ampiezza del grafo.
*/
func span(g *grafo, v int) {
	coda := []int{v}
	aux := make(map[int]bool)
	aux[v] = true

	for len(coda) > 0 {
		k := coda[0]
		coda = coda[1:]
		for _, w := range g.adiacenze[k] {
			if !aux[w] {
				aux[w] = true
				coda = append(coda, w)
				fmt.Printf("Arco %d %d\n", k, w)
			}
		}
	}
}

func dfs1(g *grafo, node int, color int, colori map[int]int) bool {
	colori[node] = color
	for _, w := range g.adiacenze[node] {
		if colori[w] == color {
			return false
		}
		if colori[w] == 0 && !dfs1(g, w, -color, colori) {
			return false
		}
	}
	return true
}

func twocolor(g *grafo) bool {
	//ok := true
	colori := make(map[int]int)
	for node, _ := range g.adiacenze {
		if colori[node] == 0 && !dfs1(g, node, 1, colori) {
			return false
		}
	}
	return true
}

func main() {
	g := gen(0.8)
	stampaGrafo(g)
	if path(g, 0, 4) {
		fmt.Println("Cammino esistente")
	} else {
		fmt.Println("Cammino NON esistente")
	}
	//ccc(g)
	//cc(g, 1)
	span(g, 1)
	if twocolor(g) {
		fmt.Println("Il grafo è bicolorabile")
	} else {
		fmt.Println("Il grafo NON è bicolorabile")
	}
}
