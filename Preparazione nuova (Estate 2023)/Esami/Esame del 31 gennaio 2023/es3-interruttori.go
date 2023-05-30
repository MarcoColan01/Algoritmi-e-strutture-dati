package main

import "fmt"

type conf struct {
	stato []bool
}

func premi(curr conf, x int) conf {
	if x == 1 {
		curr.stato[0] = !curr.stato[0]
		curr.stato[1] = !curr.stato[1]
		curr.stato[len(curr.stato)-1] = !curr.stato[len(curr.stato)-1]
	} else if x == len(curr.stato) {
		curr.stato[0] = !curr.stato[0]
		curr.stato[len(curr.stato)-1] = !curr.stato[len(curr.stato)-1]
		curr.stato[len(curr.stato)-2] = !curr.stato[len(curr.stato)-2]
	} else {
		curr.stato[x-1] = !curr.stato[x-1]
		curr.stato[x] = !curr.stato[x]
		curr.stato[x-2] = !curr.stato[x-2]
	}
	return curr
}

func sequenza(curr conf, x []int) conf {
	for i := 0; i < len(x); i++ {
		curr = premi(curr, x[i])
	}
	return curr
}

func sequenzaSpegniTutto(curr conf) []int {
	sequenza := make([]int, 0)
	coda := []int{0}
	visitati := make(map[int]bool)
	visitati[0] = true

	for len(coda) > 0 {
		app := coda[0]
		sequenza = append(sequenza, app)
		coda = coda[1:]
		curr.stato[app] = false
		visitati[app] = true

		if app == 0 {
			if curr.stato[app] == true && visitati[app] == false {
				coda = append(coda, app)
			}
			if curr.stato[len(curr.stato)-1] == true && visitati[len(curr.stato)-1] == false {
				coda = append(coda, len(curr.stato)-1)
			}
			if curr.stato[1] == false && visitati[1] == false {
				coda = append(coda, 1)
			}
		} else if app == len(curr.stato)-1 {
			if curr.stato[app] == true && visitati[app] == false {
				coda = append(coda, app)
			}
			if curr.stato[len(curr.stato)-2] == true && visitati[len(curr.stato)-2] == false {
				coda = append(coda, len(curr.stato)-2)
			}
			if curr.stato[0] == false && visitati[0] == false {
				coda = append(coda, 0)
			}
		} else {
			if curr.stato[app] == true && visitati[app] == false {
				coda = append(coda, app)
			}
			if curr.stato[app-1] == true && visitati[app-1] == false {
				coda = append(coda, app-1)
			}
			if curr.stato[app-2] == false && visitati[app-2] == false {
				coda = append(coda, app-2)
			}
		}
		fmt.Println(coda[0])
	}
	return sequenza
}

//func spegniTutto(curr conf) int {}

func main() {
	var rete conf
	rete.stato = make([]bool, 7)
	x := []int{1, 2}
	rete = sequenza(rete, x)
	fmt.Println(rete.stato)
	s := sequenzaSpegniTutto(rete)
	fmt.Println(rete.stato, s)
}
