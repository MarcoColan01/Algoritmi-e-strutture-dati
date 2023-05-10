package main 
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type linkedList struct{
	val int 
	next *linkedList 
}

type grafo struct{
	n int 		//numero vertici
	adiacenti []*linkedList
}

func nuovoGrafo(n int) *grafo{
	g := new(grafo)
	g.n = n
	g.adiacenti = make([]*linkedList, n)
	for i:=0; i < n; i++{
		g.adiacenti[i] = nil 
	}
	return g
}

func addEdge(g *grafo, x int, y int){
	arco := new(linkedList)
	arco.val = y 
	arco.next = g.adiacenti[x]
	g.adiacenti[x] = arco 
}

func leggiGrafo(g *grafo, filename string){
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan(){
		k := strings.Split(scanner.Text(), " ")
		x,_ := strconv.Atoi(k[0])
		y,_ := strconv.Atoi(k[1])
		addEdge(g, x, y)
	}
}

func stampaGrafo(g *grafo){
	for i:=0; i < g.n; i++{
		fmt.Printf("%d -> ", i)
		app := g.adiacenti[i]
		for app != nil{
			fmt.Printf("%d ", app.val)
			app = app.next
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkEdge(g *grafo, x int, y int){
	var app *linkedList
	if x < 0 || x >= g.n{
		fmt.Printf("Arco %d %d inesistente\n", x, y)
		return
	}else{
		app = g.adiacenti[x]
		for app != nil{
			if app.val == y{
				break 
			}else{
				app = app.next
			}
		}
	}
	if app == nil{
		fmt.Printf("Arco %d %d inesistente\n", x, y)
	}else{
		fmt.Printf("Arco %d %d prsente\n", x, y)
	}
}

func main(){
	n,err:= strconv.Atoi(os.Args[1])
	if err != nil{
		os.Exit(-1)
	}
	filename := os.Args[2]
	g := nuovoGrafo(n)
	leggiGrafo(g, filename)
	stampaGrafo(g)
	checkEdge(g, 1, 2)
	checkEdge(g, 1, 3)
	checkEdge(g, 6, 3)
}