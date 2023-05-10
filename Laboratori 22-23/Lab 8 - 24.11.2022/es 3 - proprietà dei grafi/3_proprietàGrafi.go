package main 
import(
	"fmt"
	"math/rand"
	"time"
	/*"os"
	"bufio"
	"strings"
	*/
)
var color int  = 0

//rand.Seed(time.Now().UnixNano())

type grafo struct{
	adiacenze map[string][]string 
}

/*
1. gen (p) genera un grafo casuale, a partire dalla probabilità p compresa tra 0 e 1 (inclusi).
Il modello matematico di riferimento è il seguente: si considerano tutti i possibili archi includendoli nel grafo con probabilità p. Più esplicitamente, per ogni possibile coppia di vertici, si genera un numero reale compreso tra 0 e 1; se questo è minore di p si inserisce l’arco, altrimenti non lo si inserisce
*/
func gen(n int) *grafo{
	var x,y byte
	x = 97
	y = 97
	rand.Seed(time.Now().UnixNano())
	g := new(grafo)
	g.adiacenze = make(map[string][]string, n)
	for i:=0; i < n; i++{
		y = 97
		for j := i; j < n; j++{
			p := rand.Intn(2)
			if p == 1 && x != y{
				g.adiacenze[string(x)] = append(g.adiacenze[string(x)], string(y))
				//g.adiacenze[string(y)] = append(g.adiacenze[string(y)], string(x))
			}
			y++
		}
		x++
	}
	return g
}

/*
2. degree (v) calcola il grado del vertice v.
Si ricorda che il grado di un vertice è definito come il numero di vertici ad esso adiacenti
*/
func degree(g *grafo, v string) int{
	return len(g.adiacenze[v])
}

/*
3. path (v, w) testa l’esistenza di un cammino semplice che collega i vertici v e w.
Si ricorda che un cammino si dice semplice quando attraversa ogni vertice al più una volta.
il numero di vertici ad esso adiacenti.
*/
func path(g *grafo, x,y string) bool{
	path := false
	aux := make(map[string]bool)
	coda := []string{x}
	aux [x] = true
	for len(coda) > 0 {
		v := coda [0]
		coda = coda [1:]
		if v == y{
			path = true
			break
		}
		for _ , v2 := range g.adiacenze[v]{
			if ! aux [v2] {
				coda = append( coda , v2 )
				aux [v2] = true
			}
		}
	}
	return path 
}

/*
4.ccc conta il numero di componenti connesse di un grafo (non orientato).
Si ricorda che si chiama componente connessa di un grafo ogni insieme massimale di
vertici connessi tra loro da un cammino. */
func bfs1 (g *grafo , v string, visitati map[string]bool) {
	aux := make(map[string]bool)
	coda := []string{v}
	aux [v] = true
	for len(coda) > 0 {
		v := coda [0]
		coda = coda [1:]
		for _ , v2 := range g.adiacenze[v]{
			if ! aux [v2] {
				coda = append( coda , v2 )
				aux [v2] = true
				visitati[v2] = true
			}
		}
	}
	return
}
func ccc(g *grafo){
	nCC := 0
	visitati := make(map[string]bool)
	for i,_ := range g.adiacenze{
		if !visitati[i]{
			bfs1(g, i, visitati)
			nCC++
		}
	}
	fmt.Printf("Numero componenti connesse nel grafo: %d\n", nCC)
}

//5. cc (v) stampa l’elenco dei vertici della componente connessa contenente v
func cc(g *grafo, v string){
	fmt.Printf("Vertici della componente connessa contenente %s: ", v)
	for i:= 0; i < len(g.adiacenze[v]); i++{
		fmt.Printf("%s ", g.adiacenze[v][i])
	}
	fmt.Println()
}


/*
6.  span (v) calcola uno spanning tree con radice v e lo stampa nella rappresentazione "a
sommario".
Si ricorda che si definisce spanning tree (in italiano, albero di copertura) un albero che
ha per nodi tutti e soli i vertici del grafo. Osservate che per ottenere uno spanning tree
con radice v è sufficiente eseguire una visita della componente connessa contenente v
stampando ad ogni passo l’arco attraversato. Che tipo di visita si deve eseguire per avere
la garanzia di ottenere uno spanning tree di altezza minimale? in ampiezza
*/
func span(g *grafo, v string){
	aux := make(map[string]bool)
	coda := []string{v}
	aux [v] = true
	for len(coda) > 0 {
		x := coda [0]
		coda = coda [1:]
		for _ , v2 := range g.adiacenze[x]{
			if ! aux [v2] {
				fmt.Printf("Arco %s %s\n", x, v2)
				coda = append( coda , v2 )
				aux [v2] = true
			}
		}
	}
}

/*
7. twocolor testa se il grafo è bicolorabile.
Un grafo si dice bicolorabile quando è possibile assegnare ad ogni vertice del grafo uno
dei due colori bianco o nero in modo che due nodi vicini abbiano sempre colori diversi. Quando un grafo è bicolorabile, si dice anche che è bipartito */
func twocolor(g *grafo, aux map[string]string, v string, ok bool) bool{
	color++
	if color%2 == 0{
		aux[v] = "black"
		
	}else{
		aux[v] = "white"
		
	}
	for _ , v2 := range g.adiacenze[v] {
		if color%2 == 0 && aux[v2] == "white" {
			fmt.Println("orcodio bianco")
			twocolor(g, aux, v2, ok)
		}else if color%2 != 0 && aux[v2] == "black" { 
			fmt.Println("orcodio nero")
			twocolor(g, aux, v2, ok)
		}else{
			ok = false
			break
		}
	}
	return ok
}



func printGraph(g *grafo){
	for i,k := range g.adiacenze{
		fmt.Printf("Nodo %s: ", i)
		for j :=0; j < len(k); j++{
			fmt.Printf("%s ", k[j])
		}
		fmt.Println()
	}
}



func main(){
	//g := gen(6)
	g := new(grafo)
	g.adiacenze = make(map[string][]string)
	g.adiacenze["a"] = append(g.adiacenze["a"], "f")
	g.adiacenze["a"] = append(g.adiacenze["a"], "g")
	g.adiacenze["b"] = append(g.adiacenze["b"], "g")
	g.adiacenze["b"] = append(g.adiacenze["b"], "i")
	g.adiacenze["c"] = append(g.adiacenze["c"], "g")
	g.adiacenze["c"] = append(g.adiacenze["c"], "h")
	g.adiacenze["c"] = append(g.adiacenze["c"], "i")
	g.adiacenze["d"] = append(g.adiacenze["d"], "f")
	g.adiacenze["e"] = append(g.adiacenze["e"], "h")
	g.adiacenze["e"] = append(g.adiacenze["e"], "h")
	/*g.adiacenze["f"] = append(g.adiacenze["f"], "a")
	g.adiacenze["f"] = append(g.adiacenze["f"], "d")
	g.adiacenze["g"] = append(g.adiacenze["g"], "a")
	g.adiacenze["g"] = append(g.adiacenze["g"], "b")
	g.adiacenze["g"] = append(g.adiacenze["g"], "c")
	g.adiacenze["h"] = append(g.adiacenze["h"], "c")
	g.adiacenze["h"] = append(g.adiacenze["h"], "e")
	g.adiacenze["i"] = append(g.adiacenze["i"], "b")
	g.adiacenze["i"] = append(g.adiacenze["i"], "c")
	g.adiacenze["i"] = append(g.adiacenze["i"], "e")*/
	printGraph(g)
	//span(g, "a")
	/*deg := degree(g, "d")
	fmt.Printf("Grado del vertice d: %d\n", deg)*/
	/*
	x := "b"
	y := "d"
	if path(g, x,y){
		fmt.Printf("Percorso %s verso %s esiste\n", x, y)
	}else{
		fmt.Printf("Percorso %s verso %s NON esiste\n", x, y)
	}
	app := make(map[string]bool)
	dfs1(g, "a", app)
	*/
	aux := make(map[string]string)
	if twocolor(g, aux, "a", true){
		fmt.Println("Il grafo è bicolorabile")
	}else{
		fmt.Println("Il grafo NON è bicolorabile")
	}
	//ccc(g)
	//cc(g, "a")
}