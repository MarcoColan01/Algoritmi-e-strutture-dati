package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type oggetto struct {
	nome          string
	val           int
	op            string
	padre, sx, dx string
}

type foresta struct {
	oggetti map[string]*oggetto //chiave: nome nodo  valore: puntatore al nodo
	radici  []*oggetto          //slice di puntatori a radici di alberi
}

func leggiInput() map[string]*oggetto { //
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(-4)
	}
	scanner := bufio.NewScanner(file)
	oggetti := make(map[string]*oggetto)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := strings.Split(scanner.Text(), " ")
		k2 := strings.Split(k[0], ":")

		if len(k) == 2 { //oggetto con valore noto
			n, _ := strconv.Atoi(k[1])
			oggetti[k2[0]] = &oggetto{k2[0], n, " ", " ", " ", " "}
		} else { //oggetto con valore non noto
			oggetti[k2[0]] = &oggetto{k2[0], 10000, k[2], " ", k[1], k[3]}
		}
	}

	return oggetti
}

func costruisciForesta(mappa map[string]*oggetto) foresta {
	var f foresta
	f.oggetti = make(map[string]*oggetto)
	f.radici = make([]*oggetto, 0)

	for i, k := range mappa {
		if k.val == 10000 { //nodo non foglia
			mappa[k.sx].padre = i
			mappa[k.dx].padre = i
		}
	}
	for i, k := range mappa {
		if k.padre == " " { // è una radice che quindi non ha padre
			f.radici = append(f.radici, mappa[i])
		}
		f.oggetti[i] = mappa[i]
	}

	return f
}

func inOrder(r *oggetto, oggetti map[string]*oggetto) {
	if r == nil {
		return
	}
	inOrder(oggetti[r.sx], oggetti)
	if r.val == 10000 {
		fmt.Println(r.nome)
	} else {
		fmt.Printf("%s (val = %d)\n", r.nome, r.val)
	}
	inOrder(oggetti[r.dx], oggetti)
}

func stampaAlbero(f foresta, n string) {
	for i := 0; i < len(f.radici); i++ {
		if f.radici[i].nome == n {
			inOrder(f.radici[i], f.oggetti)
		}
	}
}

func up(f foresta, n string) (string, bool) {
	return f.oggetti[n].padre, f.oggetti[n].padre != " "
}

func postOrder(r *oggetto, oggetti map[string]*oggetto) {
	if r == nil {
		return
	}
	postOrder(oggetti[r.sx], oggetti)
	postOrder(oggetti[r.dx], oggetti)
	if r.val == 10000 {
		switch r.op {
		case "+":
			r.val = oggetti[r.sx].val + oggetti[r.dx].val
		case "-":
			r.val = oggetti[r.sx].val - oggetti[r.dx].val
		case "*":
			r.val = oggetti[r.sx].val * oggetti[r.dx].val
		case "/":
			r.val = oggetti[r.sx].val / oggetti[r.dx].val
		}
	}
}

func calcolaPrezzo(f foresta, n string) int {
	postOrder(f.oggetti[n], f.oggetti)
	return f.oggetti[n].val
}

func main() {
	oggetti := leggiInput()
	f := costruisciForesta(oggetti)
	stampaAlbero(f, "frigorifero")
	s, err := up(f, "tavolo")
	if !err {
		fmt.Println("Il nodo non ha figlio")
	} else {
		fmt.Println(s)
	}

	n := calcolaPrezzo(f, "frigorifero")
	fmt.Println(n)
}
