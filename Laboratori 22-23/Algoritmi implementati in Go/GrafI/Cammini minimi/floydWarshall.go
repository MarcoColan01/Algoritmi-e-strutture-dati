package main 
import(
	"fmt"
)

type graph struct{
	mtr [][]int 
	vert int 
}

func newGraph(n int) *graph{
	gr := new(graph)
	gr.vert = n 
	gr.mtr = make([][]int, n)
	for i :=0; i < n; i++{
		for j:=0; j < n; j++{
			gr.mtr[i] = append(gr.mtr[i], 0)
		}
	}
	return gr
}

func printGraph(g *graph){
	for i:=0; i < g.vert; i++{
		for j :=0; j < len(g.mtr[i]); j++{
			fmt.Printf("%d ", g.mtr[i][j])
		}
		fmt.Println()
	}
}

func edgeInsert(g *graph, i int, j int, w int){
	g.mtr[i][j] = w 
}

func printMatr(m [][]int){
	for i:=0; i < len(m); i++{
		for j :=0; j < len(m[i]); j++{
			if m[i][j] == 9999{
				fmt.Printf("+inf " ,)
			}else{
				fmt.Printf("%d ", m[i][j])
			}
		}
		fmt.Println()
	}
}

func floydWarshall(g *graph)[][]int{
	d := make([][]int, g.vert)
	p := make([][]int, g.vert)

	for i:=0; i < g.vert; i++{
		for j:=0; j < g.vert; j++{
			d[i] = append(d[i], 0)
			p[i] = append(p[i], 0)
		}
	}

	
	for i:=0; i < g.vert; i++{
		for j:=0; j < g.vert; j++{
			if i == j{
				d[i][j] = 0
			}else if g.mtr[i][j] == 0{
				d[i][j] = 9999
			}else{
				d[i][j] = g.mtr[i][j]
			}
			p[i][j] = 0
		}
	}

	for k :=0; k < g.vert; k++{
		for i :=0; i < g.vert; i++{
			for j :=0; j < g.vert; j++{
				if d[i][k] + d[k][j] < d[i][j]{
					d[i][j] = d[i][k] + d[k][j]
					p[i][j] = k
				}
			}
		}
	}
	return d
	
}

/*func ricostruzioneCammino(p [][]int, part int, dest int){
	i := part 
	j := dest
	fmt.Printf("Cammino %d-%d: %d-", part, dest, part)
	for {
		if 
	}
}*/

func main(){
	g := newGraph(4)
	edgeInsert(g,0,1,5)
	edgeInsert(g,1,2,7)
	edgeInsert(g,1,3,4)
	edgeInsert(g,2,1,1)
	edgeInsert(g,3,0,2)
	edgeInsert(g,3,2,2)
	pesi := floydWarshall(g)
	fmt.Printf("Cammino minimo da %d a %d: costo %d\n", 0,2, pesi[0][2])
	//ricostruzioneCammino(cammini, 0,2)
}