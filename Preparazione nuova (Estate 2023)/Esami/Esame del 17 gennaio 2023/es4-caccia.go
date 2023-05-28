package main

import (
	"fmt"
	"sort"
)

type foglietto struct {
	nome          string
	val           int
	op            string
	padre, sx, dx *foglietto
}

type gioco struct {
	root *foglietto
}

func dfs(r *foglietto) {
	if r == nil {
		return
	}
	dfs(r.sx)
	fmt.Printf("Oggetto %s:\nValore %d\nPadre %s\n", r.nome, r.val, r.padre.nome)
	dfs(r.dx)
}

func calcolaNumero(r *foglietto) {
	if r == nil {
		return
	}
	calcolaNumero(r.sx)
	calcolaNumero(r.dx)
	if r.val == 10000 {
		switch r.op {
		case "*":
			r.val = r.sx.val * r.dx.val
			break
		case "/":
			r.val = r.sx.val / r.dx.val
			break
		case "+":
			r.val = r.sx.val + r.dx.val
			break
		case "-":
			r.val = r.sx.val - r.dx.val
			break
		}
	}
}

func riempiArray(r *foglietto, vet []int) []int {
	if r == nil {
		return vet
	}
	vet = riempiArray(r.sx, vet)
	vet = append(vet, r.val)
	vet = riempiArray(r.dx, vet)
	return vet
}

func stampaNumeri(r *foglietto) {
	vet := make([]int, 0)
	vet = riempiArray(r, vet)
	sort.Ints(vet)
	fmt.Println(vet)
}

func espressione(r *foglietto) {
	if r == nil {
		return
	}
	espressione(r.sx)
	if r.dx == nil && r.sx == nil {
		fmt.Printf("%d ", r.val)
	} else {
		fmt.Printf("%s ", r.op)
	}
	espressione(r.dx)
}

func main() {
	g := new(gioco)
	g.root = &foglietto{"Frigorifero", 10000, "*", nil, nil, nil}
	g.root.sx = &foglietto{"Tazza", 2, " ", g.root, nil, nil}
	g.root.dx = &foglietto{"Tavolo", 10000, "-", g.root, nil, nil}
	g.root.dx.dx = &foglietto{"Valigia", 3, " ", g.root.dx, nil, nil}
	g.root.dx.sx = &foglietto{"Cassetto", 5, " ", g.root.dx, nil, nil}
	//dfs(g.root)

	//Compito 1
	calcolaNumero(g.root)
	fmt.Println(g.root.val)
	fmt.Println(g.root.dx.val)

	//Compito 2
	espressione(g.root)
	fmt.Println()

	//Compito 3
	stampaNumeri(g.root)
}
