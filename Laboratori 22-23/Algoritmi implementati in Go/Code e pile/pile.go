/*
Funzioni da implementare:
	-creaPila(): crea una nuova pila
	-push(elem): mette elem in cima alla pila
	-top(): restituisce l'elemento in cima alla pila 
	-pop(): restituisce l'elemento in cima alla pila e lo rimuove
	-isEmpty(): restituisce true se la pila è vuota, false altrimenti 
	-stampaPila(): stampa la pila
*/

//struttura dati: lista concatenata con riferimento alla testa e inserimenti in testa

package main 

import(
	"fmt"
)

type listNode struct{
	val int 
	next *listNode
}

type pila struct{
	head *listNode
}

func creaPila() *pila{
	p := new(pila)
	p.head = nil 
	return p
}

func newNode(n int) *listNode{
	new := new(listNode)
	new.next = nil 
	new.val = n
	return new
}

func push(p *pila, n int){
	new := newNode(n)
	new.next = p.head 
	p.head = new
}

func top(p *pila) *listNode{
	return p.head
}

func pop(p *pila) *listNode{
	var head *listNode
	if p.head != nil{
		head = p.head 
		p.head = p.head.next
	}else{
		head = nil
	}
	return head 
}

func stampaPila(p *pila){
	app := p.head 
	for{
		if app == nil{
			return
			break
		}
		fmt.Println(app.val)
		app = app.next
	}
	fmt.Println()
	return
}

func isEmpty(p *pila) bool{
	return p.head == nil
}

func main(){
	var p *pila 
	var ch string 
	var n int 
	exit := false
	for !exit{
		fmt.Scanf("%s %d", &ch, &n)
		switch ch{
			case "n":
				p = creaPila()
			case "+":
				push(p, n)
			case "-":
				n := pop(p)
				if n != nil{
					fmt.Println("Elemento", n.val, "rimosso dalla pila")
				}else{
					fmt.Println("No push: pila vuota")
				}
			case "p":
				stampaPila(p)
			case "?":
				if isEmpty(p){
					fmt.Println("La pila è vuota")
				}else{
					fmt.Println("La pila NON è vuota")
				}
			case "1":
				n := top(p)
				if n != nil{
					fmt.Println("Primo elemento della pila: ", n.val)
				}else{
					fmt.Println("No top: pila vuota")
				}
			case "f":
				exit = true
				break
		}
	}
}