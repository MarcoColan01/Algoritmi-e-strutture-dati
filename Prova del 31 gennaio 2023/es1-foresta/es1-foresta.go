package main

import (
	"fmt"
)


type oggetto struct {
	nome string
	val  int	// non rilevante se l'indizio e' un'operazione
	dx   string
	op   rune // " " se numero
	sx   string
	tipo string // "num" se l'indizio è un numero, "op" se l'indizio è una operazione
}


func leggiInput() map[string]*oggetto {
	....
}


type foresta ....


func costruisciForesta(mappa map[string]*oggetto) foresta {
	....
}

func stampaAlbero(f foresta, name string) {
	...
}

func sx(f foresta, n string) (string, bool) {
	....
}

func dx(f foresta, n string) (string, bool) {
	....
}

func up(f foresta, n string) (string, bool) {
	....
}

// non modificare
func main() {
	mappa := leggiInput()
	f := costruisciForesta(mappa)
	stampaForesta(f)
}
