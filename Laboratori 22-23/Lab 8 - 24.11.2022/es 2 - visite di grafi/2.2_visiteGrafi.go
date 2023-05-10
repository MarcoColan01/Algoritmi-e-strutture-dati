package main 
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

type vertice struct{
	val string		
}

type grafo struct{
	adiacenze map[string][]string
}

func graphNew() *grafo{
	g := new(grafo)
	g.adiacenze = make(map[string][]string)
	return g 
}


func addEdge(g *grafo, a string, b string){
	g.adiacenze[a] = append(g.adiacenze[a], b)
}

func bfs1 (g *grafo , v string) {
	aux := make(map[string]bool)
	coda := []string{v}
	aux [v] = true
	for len(coda) > 0 {
		v := coda [0]
		coda = coda [1:]
		fmt.Println(v)
		for _ , v2 := range g.adiacenze[v]{
			if ! aux [v2] {
				coda = append( coda , v2 )
				aux [v2] = true
			}
		}
	}
	fmt.Println()
}
func dfs1 (g *grafo , v string, aux map[string]bool) {
	fmt.Println(v)
	aux[v] = true
	for _ , v2 := range g.adiacenze[v] {
		if aux [ v2 ] != true {
			dfs1 (g, v2, aux )
		}
	}
}

func main(){
	g := graphNew()
	filename:= os.Args[1]
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan(){
		edge := strings.Split(scanner.Text(), " ")
		addEdge(g, edge[0], edge[1])
	}
	fmt.Println("bfs:")
	bfs1(g, "a")
	aux := make(map[string]bool)
	fmt.Println("dfs: ")
	dfs1(g, "a", aux)

}