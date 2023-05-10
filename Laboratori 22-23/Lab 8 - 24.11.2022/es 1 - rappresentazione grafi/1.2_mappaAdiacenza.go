package main 
import(
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)

type vertice struct{
	età int 		//informazioni del vertice
	key string 		//chiave del vertice (nome)
	//hobby []string
}

type grafo struct{
	vertici map[string] *vertice 			//ad ogni key string associa il corrispondente vertice
	adiacenti map[vertice] []*vertice		//lista dei vertici adiacenti ad un vertice
	hobby map[*vertice] []string
}

func graphNew() *grafo{
	g := new(grafo)
	g.vertici = make(map[string]*vertice)
	g.adiacenti = make(map[vertice] []*vertice)
	g.hobby = make(map[*vertice] []string)
	return g 
}

func newVertex(name string, age int) *vertice{
	var v vertice 
	v.key = name 
	v.età = age 
	return &v
}

func addUser(g *grafo, name string, age int, hobbies []string){
	vertex := newVertex(name, age)
	g.vertici[name] = vertex 
	g.hobby[vertex] = hobbies
}

func readGraph(g *grafo, filename string){
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan(){
		k := strings.Split(scanner.Text(), " ")
		if len(k) != 2{
			n,_ := strconv.Atoi(k[1])
			addUser(g, k[0], n, k[2:])
		}else{
			addEdge(g, k[0], k[1])
		}
	}
}

func addEdge(g *grafo, a string, b string){
	follower := g.vertici[b]
	seguito := *g.vertici[a]
	g.adiacenti[seguito] = append(g.adiacenti[seguito], follower)
}

func printGraph(g *grafo){
	for i,k := range g.vertici{
		fmt.Printf("Utente %s:\nEtà %d\nHobby:%v\n\n", i, k.età, g.hobby[k])
	}

	/*for i,k := range g.adiacenti{
		fmt.Printf("Utente %s segue: ", i.key)
		for j :=0; j < len(k); j++{
			fmt.Printf("%s ", k[j].key)
		}
		fmt.Println()
	} */
	
}

func printHobbies1(g *grafo, user string){
	app := g.vertici[user]
	fmt.Printf("Hobby dell'utente %s: %v\n", user, g.hobby[app])
	for i,k := range g.adiacenti{
		for j:=0; j < len(k); j++{
			if k[j] == app{
				fmt.Printf("Hobby dell'utente %s che segue %s: %v\n", i.key, user, g.hobby[k[j]])
			} 
		}
	}
}

func printHobbies2(g *grafo, user string){
	app := g.vertici[user]
	fmt.Printf("Hobby dell'utente %s: %v\n", user, g.hobby[app])
	followers := g.adiacenti[*app]
	for i:=0; i < len(followers); i++{
		fmt.Printf("Hobby dell'utente %s seguito da %s: %v\n", followers[i].key, user, g.hobby[followers[i]])
	}
}

func main(){
	filename := os.Args[1]
	g := graphNew()
	readGraph(g, filename)
	printGraph(g)
	printHobbies1(g, "simone")
	printHobbies2(g, "marco")
}