/*
Scrivete un programma con l’implementazione di una lista concatenata semplice, seguendo i lucidi presentati a lezione. Definite i tipi listNode e linkedList e scrivete queste funzioni:
	• newNode, che crea un nuovo nodo di lista;
	• addNewNode, che inserisce un nuovo nodo in testa alla lista;
	• printList, che stampa una lista;
	• searchList, che cerca un elemento in una lista;
	• removeItem, che cancella un item da una lista.
Per testare le vostre funzioni scrivete una funzione main che gestisca un insieme dinamico (che variano nel tempo) di interi. Il main deve leggere da standard input una sequenza di istruzioni secondo il formato nella tabella qui sotto, dove n è un intero. I vari elementi sulla riga sono separati da uno o più spazi. Quando una riga è letta, viene eseguita l’operazione associata; le operazioni di stampa sono effettuate sullo standard output, e ogni operazione deve iniziare su una nuova riga.
*/

package main 
import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type listNode struct{
	val int 
	next *listNode
}

type linkedList struct{
	head *listNode
	card int		//cardinalità, numero di elementi in lista
}

func newNode(n int) *listNode{
	new := new(listNode)
	new.next = nil 
	new.val = n
	return new
}

func addNewNode(l *linkedList, n int){
	new := newNode(n)
	new.next = l.head 
	l.head = new
	l.card++
}

func printList(l *linkedList){
	app := l.head
	for{
		if app == nil{
			break
		}
		fmt.Print(app.val, " -> ")
		app = app.next
	}
	fmt.Println()
	return
}

/*func printInvList(l *linkedList){

}*/

func searchList(l *linkedList, n int) bool{
	trovato := false
	app := l.head 
	for{
		if app == nil{
			break 
		}
		if app.val == n{
			trovato = true 
			break
		}
		app = app.next
	}
	return trovato
}

func removeItem(l *linkedList, n int){
	var cur, prev *listNode 
	cur = l.head 
	if cur.val == n{		//il nodo da cancellare è in testa alla lista
		cur = cur.next
		l.head = cur 
	}else{
		for{
			if cur == nil{
				break
			}
			if cur.val == n{
				prev.next = cur.next 
				break
			}
			prev = cur 
			cur = cur.next
		}
	}
	l.card--
	return
}

func main(){
	filename := os.Args[1]
	file,err := os.Open(filename)

	if err != nil{
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	l := new(linkedList)
	l.head = nil

	var ch string
	var n int 
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan(){
		cmd := scanner.Text()
		if (cmd[0] == '+' || cmd[0] == '-' || cmd[0] == '?') && len(cmd) > 2{
			ch = string(cmd[0])
			n,_ = strconv.Atoi(string(cmd[2:]))
			if ch == "+"{
				if !searchList(l, n){
					addNewNode(l, n)
				}
			}else if ch == "-"{
				if searchList(l, n){
					removeItem(l, n)
				}
			}else if ch == "?"{
				if searchList(l, n){
					fmt.Println(n, "è presente nella lista")
				}else{
					fmt.Println(n, "NON è presente nella lista")
				}
			}

		}else if cmd[0] == 'p' ||cmd[0] == 'o' ||cmd[0] == 'f' ||cmd[0] == 'd' || cmd[0] == 'c'{
			ch = string(cmd[0])
			if ch == "p"{
				printList(l)
			}else if ch == "c"{
				fmt.Println("Numero elementi nella lista:", l.card)
			/*}else if ch == 'o'{
				printInvList(l)
			*/}else if ch == "d"{
				for i:=0; i < l.card; i++{
					removeItem(l, l.head.val)
				}
			}else if ch == "f"{
				break
			}
		}
	}
}