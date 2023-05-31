package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type stazione struct {
	nome  string
	linee []int
}
type rete struct {
	n         int
	stazioni  []stazione
	indexes   map[string]int
	adiacenti map[string][]string
}

func leggiDati(nomeFile string) rete {
	aux := make(map[string]bool)
	var r rete
	r.n = 0
	r.stazioni = make([]stazione, 0)
	r.adiacenti = make(map[string][]string)
	r.indexes = make(map[string]int)
	file, err := os.Open(nomeFile)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	//cont := 0
	for scanner.Scan() {
		k := strings.Split(scanner.Text(), ":")
		linea, _ := strconv.Atoi(strings.Split(k[0], " ")[1])
		st := strings.Split(k[1], ", ")
		st[0] = st[0][1:]
		st[len(st)-1] = strings.TrimSuffix(st[len(st)-1], ".")
		for i := 0; i < len(st); i++ {
			if aux[st[i]] {
				for k := 0; k < len(r.stazioni); k++ {
					if r.stazioni[k].nome == st[i] {
						r.stazioni[k].linee = append(r.stazioni[k].linee, linea)
						break
					}
				}
			} else {
				var app stazione
				app.nome = st[i]
				app.linee = append(app.linee, linea)
				r.stazioni = append(r.stazioni, app)
				aux[st[i]] = true
				r.indexes[st[i]] = r.n
				r.n++
			}
			if i == 0 {
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
			} else if i == len(st)-1 {
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
			} else {
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
			}
		}
	}
	//fmt.Println(r.indexes)
	return r
}

func stampaRete(r rete) {
	for i := 0; i < len(r.stazioni); i++ {
		fmt.Printf("Stazione %s  Linee: %d   Adiacenti: %s\n", r.stazioni[i].nome, r.stazioni[i].linee, r.adiacenti[r.stazioni[i].nome])
	}
}

func tempo(mm rete, partenza, arrivo string) int {
	distanze := make(map[string]int)
	for _, stazione := range mm.stazioni {
		distanze[stazione.nome] = math.MaxInt32
	}

	distanze[partenza] = 0
	coda := []string{partenza}

	for len(coda) > 0 {
		curr := coda[0]
		coda = coda[1:]
		if curr == arrivo {
			return distanze[curr]
		}
		for _, adiacente := range mm.adiacenti[curr] {
			fmt.Println(len(mm.stazioni[mm.indexes[curr]].linee), len(mm.stazioni[mm.indexes[adiacente]].linee))
			if distanze[adiacente] == math.MaxInt32 {
				if len(mm.stazioni[mm.indexes[curr]].linee) > 1 && len(mm.stazioni[mm.indexes[adiacente]].linee) > 1 {
					distanze[adiacente] = distanze[curr] + 2
				} else {
					distanze[adiacente] = distanze[curr] + 1
				}
				coda = append(coda, adiacente)
			}
		}
	}
	return -1
}

func main() {
	r := leggiDati("input-metro.txt")

	n := tempo(r, "Istria", "Piola")
	fmt.Println(n)
	//stampaRete(r)
}
