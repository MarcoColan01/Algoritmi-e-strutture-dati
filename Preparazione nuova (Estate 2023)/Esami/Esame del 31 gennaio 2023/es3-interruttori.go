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

//func sequenzaSpegniTutto(curr conf) []int {}

//func spegniTutto(curr conf) int {}

func main() {
	var rete conf
	rete.stato = make([]bool, 7)
	rete = premi(rete, 1)
	fmt.Println(rete.stato)
}
