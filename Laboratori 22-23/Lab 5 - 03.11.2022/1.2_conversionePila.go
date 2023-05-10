
package main 

import(
	"fmt"
	"strconv"
)

type listNode struct{
	val string 
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

func newNode(n string) *listNode{
	new := new(listNode)
	new.next = nil 
	new.val = n
	return new
}

func push(p *pila, n string){
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

/*
1. leggi un token ;
2a. se il token è un numero stampa il token ;
2b. se il token è un operatore inserisci il token in cima alla pila ;
2c. se il token è una parentesi aperta ignora il token
2d. se il token è una parentesi chiusa estrai l’operatore in cima alla pila; stampalo;
*/
func valutaConversione(p *pila, cmd string) string{
	var res string
	for i:=0; i < len(cmd); i++{
		_,err := strconv.Atoi(string(cmd[i]))
		if err == nil{		//numero
			res = res + string(cmd[i]) + " "
		}else{
			if cmd[i] == '+' || cmd[i] == '-' || cmd[i] == '*' || cmd[i] == '/'{
				push(p, string(cmd[i]))
			}else if cmd[i] == ')'{
				k := pop(p)
				res = res + k.val + " "
			}
		}
	}
	return res
}

func main(){
	p := creaPila()
	cmd := "( ( ( 5 + 3 ) * 2 ) / 2 )"
	res := valutaConversione(p, cmd)
	fmt.Println("Prima:", cmd, "\nDopo:", res)
}