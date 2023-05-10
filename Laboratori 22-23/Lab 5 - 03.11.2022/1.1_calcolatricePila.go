package main 

import(
	"fmt"
	"strconv"
)


type listNode struct{
	val int 
	next *listNode
}

type pila struct{
	head *listNode
	card int
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
	p.card++
}

func top(p *pila) *listNode{
	return p.head
}

func pop(p *pila) *listNode{
	var head *listNode
	if p.head != nil{
		head = p.head 
		p.head = p.head.next
		p.card--
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

/*
1.	leggi un token ( operatore o numero );
2.	se il token è un numero inseriscilo nella pila ;
3.	se il token è un operatore 
	3a. estrai due valori dalla pila ;
	3b. applica ad essi l’operatore;
	3c. inserisci il risultato nella pila;
*/

func valuta(p *pila, cmd string) int{
	var app *listNode
	for i :=0; i < len(cmd); i++{
		if cmd[i] != ' '{
			n,err := strconv.Atoi(string(cmd[i]))
			if err == nil{		//numero
				push(p, n)
			}else{				//operazione
				app = pop(p)
				n1 := app.val
				app = pop(p)
				n2 := app.val
				if cmd[i] == '+'{
					n1 += n2
					push(p, n1)
				}else if cmd[i] == '*'{
					n1 *= n2
					push(p, n1)
				}else if cmd[i] == '/'{
					n2 /= n1
					push(p, n2)
				}else if cmd[i] == '-'{
					n2 -= n1
					push(p, n2)
				}
			}
		}
	}
	app = pop(p)
	return app.val
}

func main(){
	p := creaPila()
	cmd := "5 3 + 2 *"
	n := valuta(p, cmd)
	fmt.Println("Risultato:", n)
	
}