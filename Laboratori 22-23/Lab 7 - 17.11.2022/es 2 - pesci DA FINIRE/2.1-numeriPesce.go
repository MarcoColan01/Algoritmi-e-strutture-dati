/*
• Definite un tipo di dato di Go per rappresentare i numeri-pesci come alberi.

• Scrivete una funzione readTree che, data una stringa con parentesi che indica un numeropesce, costruisca l’albero corrispondente.

• Scrivete una funzione printTree che, dato un albero che rappresenta un numero-pesce, stampi il numero-pesce (come stringa con parentesi).

• Scrivete una funzione treeToSlice che, dato un albero che rappresenta un numero-pesce, produce una slice con le sue foglie, visitate da sinistra a destra. Ad esempio, sul numeropesce rappresentato graficamente qui sopra, la slice conterrà nell’ordine i tre nodi 1,2,3.
*/

package main 
import(
	"fmt"
	//"strings"
	"strconv"
	//"math"
)

type fishNode struct{
	val int 
	dx,sx *fishNode
}

type fishTree struct{
	root *fishNode
}

func newFishNode(n int) *fishNode{
	new := new(fishNode)
	new.dx = nil 
	new.sx = nil
	new.val = n 
	return new
}

func newFishTree()*fishTree{
	ft := new(fishTree)
	ft.root = nil 
	return ft
}

func readTree(fish string) *fishNode{
	if len(fish) == 0{
		return nil 
	}
	var fn *fishNode
	if fish[0] == ']' || fish[0] == '['{
		fn = newFishNode(1000000)
	}else{
		k,_:= strconv.Atoi(string(fish[0]))
		fn = newFishNode(k)
	}
	fn.dx = readTree(fish[1:])
	return fn
}

func printTree(ft *fishNode){
	if ft == nil{
		return
	}
	printTree(ft.sx)
	if ft.val != 1000000{
		fmt.Print(ft.val," ")
	}else{
		fmt.Print("[ ")
	}
	printTree(ft.dx)
}

func main(){
	fish := "[[1 2] 3]"
	ft := newFishTree()
	ft.root = readTree(fish)
	printTree(ft.root)
	fmt.Println()
}