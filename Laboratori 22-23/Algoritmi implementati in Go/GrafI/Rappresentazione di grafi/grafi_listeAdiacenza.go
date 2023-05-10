/*
Funzioni da implementare:
	-newGraph()
	-edgeInsert()
	-printGraph()
	-destroyGraph()

Struttura dati: array di liste d'adiacenza implementato come array di slice 
*/

package main 

import(
	"fmt"
	/*"os"
	"bufio"
	"strconv"*/
)

type graph struct{
	g [][]int 
	vertici int 
}

func newGraph(n int) *graph{
	gph := new(graph)
	gph.g = make([][]int, n)
	gph.vertici = n
	return gph
}

func printGraph(gr *graph){
	for i:=0; i < len(gr.g); i++{
		fmt.Print(i, " -> ")
		//app := gr.g[i]
		for j :=0; j < len(gr.g[i]); j++{
			fmt.Print(gr.g[i][j], " -> ")
		}
		fmt.Println()
	}
	return 
}

func edgeInsert(gr *graph, vert int, edge int){
	gr.g[vert] = append(gr.g[vert], edge)
	return
}

func destroyGraph(gr *graph){
	gr.g = make([][]int, gr.vertici)
}

func main(){
	g := newGraph(5)
	edgeInsert(g, 0, 1)
	edgeInsert(g, 2, 1)
	edgeInsert(g, 4, 2)
	edgeInsert(g, 4, 3)
	printGraph(g)
	destroyGraph(g)
	printGraph(g)
}

