package main

import "fmt"

/*
1. Si progetti e descriva un algoritmo che, dato un intero k e un albero di ricerca binario, stabilisca se esistono due nodi nell'albero le cui chiavi hanno somma pari a k

L'algoritmo esegue una visita in preordine dell'albero: ogni volta che visita un nodo, cerca (sfruttando la proprietà di albero binario di ricerca) nel suo sottoalbero sx o dx un nodo il cui valore sommato al nodo visitato è pari al valore da cercare: se è vero restituisce true e interrompe la ricerca, altrimenti, continua ad esplorare i sottoalberi alla ricerca di tale valore, se esiste. Se non esiste, la visita in preordine procede sui sottoalberi sx e dx.
*/

/*
2. Si stimi la complessità dell'algoritmo
L'algoritmo esegue una visita in preordine dell'albero e, per ogni nodo visitato, nel caso peggior esamina tutti gli altri nodi dell'albero. Essendo n i nodi, la complessità dell'algoritmo è stimabile come O(n^2).
*/

type node struct {
	val    int
	dx, sx *node
}

type tree struct {
	root *node
}

func dfs(r *node) {
	if r == nil {
		return
	}
	dfs(r.sx)
	fmt.Println(r.val)
	dfs(r.dx)
}

func insert(t *node, n int) *node {
	if t == nil {
		t = &node{n, nil, nil}
	} else if n < t.val {
		t.sx = insert(t.sx, n)
	} else {
		t.dx = insert(t.dx, n)
	}
	return t
}

func trovato(t *node, n int) bool {
	if t == nil {
		return false
	}
	if t.val == n {
		return true
	} else if t.val < n {
		return trovato(t.dx, n)
	} else {
		return trovato(t.sx, n)
	}
}

func f(t *node, n int) bool {
	if t == nil {
		return false
	}
	if trovato(t, n-t.val) {
		return true
	}

	return f(t.sx, n) || f(t.dx, n)
}

func main() {
	t := new(tree)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 3)
	t.root = insert(t.root, 5)
	t.root = insert(t.root, 31)
	t.root = insert(t.root, 14)
	t.root = insert(t.root, 1)
	t.root = insert(t.root, 8)

	if f(t.root, 19) {
		fmt.Println("vero")
	} else {
		fmt.Println("falso")
	}

	//dfs(t.root)

}
