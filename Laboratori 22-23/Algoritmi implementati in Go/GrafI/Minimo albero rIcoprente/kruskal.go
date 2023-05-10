package main 
import (
	"fmt"
	"sort"
)


type edges struct{
	vertex int
	edge int
}

type app struct{
	vert int 
	edge int 
	w int
}

type graph struct{
	g map[edges]int
	vert int
}

func newGraph(n int) *graph{
	gr := new(graph)
	gr.g = make(map[edges]int)
	gr.vert = n
	return gr
}

func edgeInsert(gr *graph, vert int, edge int, w int){
	if w >= 0{
		var e edges 
		e.vertex = vert 
		e.edge = edge 
		gr.g[e] = w
	}
	return
}

func printGraph(gr *graph){
	for i,j := range gr.g{
		fmt.Println("(",i.vertex,",",i.edge,"): costo", j)
	}
	return 
}

func find(i int, parent []int) int{
	for parent[i] != i{
		i = parent[i]
	}
	return i
}

func union(parent []int, i int, j int){
	a := find(i, parent)
	b := find(j, parent)
	parent[a] = b
}

func kruskalMTS(g *graph, parent []int){
	mincost := 0;

	//makeset per ogni vertice
	for i:=0; i < g.vert; i++{
		parent[i] = i 
	}

	weigths := make([]app, 0)
	for i,k := range g.g{
		var j app 
		j.vert = i.vertex 
		j.edge = i.edge 
		j.w = k
		weigths = append(weigths, j)
	}

	//ordino gli archi in modo non decrescente in base ai pesi
	sort.SliceStable(weigths, func (i, j int) bool {return weigths[i].w < weigths[j].w})


	edgeCount := 0
	min := -1
	for edgeCount < g.vert +1{
		a := find(weigths[edgeCount].vert, parent)
		b := find(weigths[edgeCount].edge, parent)
		if a != b {
			union(parent,a,b)
			min = weigths[edgeCount].w 
			fmt.Printf("Arco %d:(%d %d) costo: %d \n", edgeCount, weigths[edgeCount].vert, weigths[edgeCount].edge, min)
			mincost += min 
		}
		edgeCount++
	}
	fmt.Printf("\nCosto minimo: %d \n", mincost)
}


func main(){
	g := newGraph(6)
	parent := make([]int, g.vert)
	edgeInsert(g,2,5,1)
	edgeInsert(g,0,1,2)
	edgeInsert(g,0,2,3)
	edgeInsert(g,2,4,6)
	edgeInsert(g,2,3,3)
	edgeInsert(g,0,3,5)
	edgeInsert(g,1,5,7)
	edgeInsert(g,3,4,8)
	
	//printGraph(g)
	kruskalMTS(g, parent)
} 