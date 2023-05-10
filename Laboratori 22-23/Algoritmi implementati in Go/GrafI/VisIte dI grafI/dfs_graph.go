package main 

import(
	"fmt"
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

func dfs_graph(g *graph, s int, visitati []int){
	visitati[s] = 1
	fmt.Printf("%d ", s)
	for i:=0; i < len(g.g[s]); i++{
		if visitati[g.g[s][i]] == 0{
			dfs_graph(g, g.g[s][i], visitati)
		}
	}
}

func main(){
	gr := newGraph(11)
	visitati := make([]int, gr.vertici)
	edgeInsert(gr,1,2)
	edgeInsert(gr,1,4)
	edgeInsert(gr,1,6)
	edgeInsert(gr,2,3)
	edgeInsert(gr,2,5)
	edgeInsert(gr,4,5)
	edgeInsert(gr,3,5)
	edgeInsert(gr,5,7)
	edgeInsert(gr,5,9)
	edgeInsert(gr,6,7)
	edgeInsert(gr,6,8)
	edgeInsert(gr,6,7)
	edgeInsert(gr,7,8)
	edgeInsert(gr,9,10)
	printGraph(gr)
	dfs_graph(gr, 1, visitati)
}