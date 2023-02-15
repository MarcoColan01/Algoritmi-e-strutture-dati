/*

Esercizio dell'esame di luglio 2022
Scrivere una funzione in go che, data la lista di adiacenza di un generico grafo, costruisca l'equivalenza matrice d'adiacenza del grafo stesso.
Ad esempio, se la lista d'adiacenza per un grafo di 5 vertici è la seguente:
	A -> B -> C -> D
	B -> D -> E
	C -> A
	D -> B -> E
	E -> A -> C
Il programma genererà la seguente matrice d'adiacenza:

0 1 1 1 0
0 0 0 1 1
1 0 0 0 0
0 1 0 0 1
1 0 1 0 0
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type linkedList struct {
	val  int
	next *linkedList
}

type grafo struct {
	n         int //numero vertici
	adiacenti []*linkedList
}

func nuovoGrafo(n int) *grafo {
	g := new(grafo)
	g.n = n
	g.adiacenti = make([]*linkedList, n)
	for i := 0; i < n; i++ {
		g.adiacenti[i] = nil
	}
	return g
}

// aggiunta di un nuovo arco
func addEdge(g *grafo, x int, y int) {
	arco := new(linkedList)
	arco.val = y
	arco.next = g.adiacenti[x]
	g.adiacenti[x] = arco
}

// stampa il grafo
func stampaGrafo(g *grafo) {
	for i := 0; i < g.n; i++ {
		fmt.Printf("%d -> ", i)
		app := g.adiacenti[i]
		for app != nil {
			fmt.Printf("%d ", app.val)
			app = app.next
		}
		fmt.Println()
	}
	fmt.Println()
}

func listToMatrix(graph *grafo) [][]int {
	matrix := make([][]int, graph.n)
	for i := 0; i < graph.n; i++ {
		matrix[i] = make([]int, graph.n)
	}
	for i := 0; i < graph.n; i++ {
		app := graph.adiacenti[i]
		for app != nil {
			matrix[i][app.val] = 1
			app = app.next
		}
	}
	return matrix
}

func main() {
	filename := os.Args[1]
	n, _ := strconv.Atoi(os.Args[2]) //n vertici del grafo
	graph := nuovoGrafo(n)
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan() {
		k := strings.Split(scanner.Text(), ":")
		nodo, _ := strconv.Atoi(k[0])
		archi := strings.Split(k[1], " ")
		//fmt.Println(archi)
		for i := 0; i < len(archi); i++ {
			if archi[i] != " " {
				arco, _ := strconv.Atoi(string(archi[i]))
				addEdge(graph, nodo, arco)
			}
		}
	}
	stampaGrafo(graph)
	matrix := listToMatrix(graph)
	for i := 0; i < graph.n; i++ {
		for j := 0; j < graph.n; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
