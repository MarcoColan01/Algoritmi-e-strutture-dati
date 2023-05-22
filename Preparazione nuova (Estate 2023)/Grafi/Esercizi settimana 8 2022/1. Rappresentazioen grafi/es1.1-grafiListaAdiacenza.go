package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	val  int
	next *listNode
}

type linkedList struct {
	head *listNode
}

type grafo struct {
	n         int
	adiacenti []*linkedList
}

func nuovoGrafo(n int) *grafo {
	g := new(grafo)
	g.n = n
	g.adiacenti = make([]*linkedList, n)
	for i := 0; i < n; i++ {
		g.adiacenti[i] = new(linkedList)
		g.adiacenti[i].head = nil
	}
	return g
}

func addEdge(g *grafo, u int, w int) {
	arco := &listNode{w, nil}
	arco.next = g.adiacenti[u].head
	g.adiacenti[u].head = arco
}

func leggiGrafo(n int) *grafo {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}
	scanner := bufio.NewScanner(file)
	g := nuovoGrafo(n)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(k[0])
		v, _ := strconv.Atoi(k[1])
		addEdge(g, u, v)
	}
	return g
}

func stampaGrafo(g *grafo) {
	for i := 0; i < g.n; i++ {
		fmt.Printf("Nodo %d: ", i)
		app := g.adiacenti[i].head
		for app != nil {
			fmt.Printf("%d -> ", app.val)
			app = app.next
		}
		fmt.Println()
	}
}

func checkEdge(g *grafo, u int, v int) bool {
	app := g.adiacenti[u].head
	for app != nil && app.val != v {
		app = app.next
	}
	return !(app == nil)
}

func main() {
	var n int
	fmt.Scan(&n)
	g := leggiGrafo(n)
	stampaGrafo(g)
	u, v := 3, 5
	if checkEdge(g, u, v) {
		fmt.Printf("Arco %d-%d presente\n", u, v)
	} else {
		fmt.Printf("Arco %d-%d NON presente\n", u, v)
	}
}
