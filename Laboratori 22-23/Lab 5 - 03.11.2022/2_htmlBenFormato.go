
package main 

import(
	"fmt"
	//"strconv"
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
leggi un tag t;
	se t è un tag di apertura inserisci t nella pila ;
	se t è un tag di chiusura estrai il tag t2 dalla cima della pila ;
	se t e t2 non corrispondono il documento non ben formato ;

*/

func checkHtml(p *pila, cmd []string){
	for i:=0; i < len(cmd); i++{
		if cmd[i][1] != '/'{	//tag di apertura
			push(p, cmd[i])
		}else{					//tag di chiusura
			k := pop(p)
			if cmd[i][2] != k.val[1]{
				break
			}
		}
	}
	return
}

func main(){
	var c string 
	var cmd []string
	p := creaPila()
	cmd = make([]string, 0)

	for{
		fmt.Scan(&c)
		if c[0] != '<'{
			break
		}
		cmd = append(cmd, c)
	}
	checkHtml(p, cmd)
	if !isEmpty(p){
		fmt.Println("Documento NON ben formato. Sono rimasti i tag:")
		var app *listNode
		for{
			app = pop(p)
			if app == nil{
				break
			}
			fmt.Print(app.val, " ")
		}
		fmt.Println()
	}else{
		fmt.Println("Documento ben formato")
	}
}