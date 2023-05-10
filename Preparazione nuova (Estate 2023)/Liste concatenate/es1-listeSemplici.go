package main 

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type listNode struct{
	val int 
	next *listNode
}

type linkedList struct{
	head *listNode 
}

func newNode(n int) *listNode{
	node := new (listNode)
	node.next = nil
	node.val = n
	return node
}

func addNewNode(l *linkedList, n int){
	node := newNode(n)
	node.next = l.head 
	l.head = node  
}

func printList(l *linkedList){
	node := l.head
	for(node != nil){
		fmt.Printf("%d -> ", node.val)
		node = node.next 
	}
	fmt.Println()
}

func searchList(l *linkedList, n int) *listNode{
	node := l.head
	for(node != nil){
		if node.val == n{break}  else {node = node.next } 
	}
	return node 
}

func removeItem(l *linkedList, n int) bool{
	var curr, prev *listNode = l.head, nil 
	for curr != nil{
		if curr.val == n{
			if prev == nil{
				l.head = l.head.next 
			}else{
				prev.next = curr.next 
			}
			return true 
		}
		prev = curr 
		curr = curr.next 
	}
	return false 
}

func deleteList(l *linkedList){
	l.head = nil
}

func listCard(l *linkedList) int{
	var n int = 0 
	node := l.head 
	for node != nil{
		n++
		node = node.next 
	}
	return n 
}

func main(){
	l := new (linkedList)
	exit := false;
	filename := os.Args[1]
	file,err := os.Open(filename)
	if err != nil{
		os.Exit(-4)
	}
	scanner := bufio.NewScanner(file) 
	scanner.Split(bufio.ScanLines)
	defer file.Close()
	for scanner.Scan(){
		r := strings.Split(scanner.Text() , " ")
		if len(r) == 1{
			switch(r[0]){
				case "c":
					n := listCard(l) 
					fmt.Printf("Cardinalità lista: %d\n", n) 	
				break

				case "f":
				break

				case "p":
					printList(l)
				break 

				case "d":
					deleteList(l) 
				break 

				/*case 'o': 
				break*/

				default:
					fmt.Println("Comando non riconosciuto")
			}
		}else{
			n,_ := strconv.Atoi(r[1])
			switch(r[0]){
				case "+":
					addNewNode(l, n)
				break 

				case "-":
					if removeItem(l, n){
						fmt.Printf("Elemento %d rimosso\n", n)
					}else{
						fmt.Printf("Elemento %d NON rimosso\n", n)
					}
				break 

				case "?": 
					if searchList(l, n) != nil{
						fmt.Printf("Elemento %d appartiene alla lista\n", n)
					}else{
						fmt.Printf("Elemento %d NON appartiene alla lista\n", n)
					}
				break 

				default:
					fmt.Println("Comando non riconosciuto")
				break
			}
		}
		if exit == true{
			break
		}
	}

}


