package main

import (
	"bufio"
	"fmt"
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

	cont := 0
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
			}
			if i == 0 {
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
			} else if i == len(st)-1 {
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
			} else {
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
			}
			r.indexes[st[i]] = cont
			cont++
		}
	}
	return r
}

func stampaRete(r rete) {
	for i := 0; i < len(r.stazioni); i++ {
		fmt.Printf("Stazione %s  Linee: %d   Adiacenti: %s\n", r.stazioni[i].nome, r.stazioni[i].linee, r.adiacenti[r.stazioni[i].nome])
	}
}

// Complessità: O(n+l)
func linea(mm rete, numLinea int) []string {
	stazioni := make([]string, 0)
	for i := 0; i < len(mm.stazioni); i++ {
		if len(mm.stazioni[i].linee) == 1 && mm.stazioni[i].linee[0] == numLinea {
			stazioni = append(stazioni, mm.stazioni[i].nome)
		} else {
			for j := 0; j < len(mm.stazioni[i].linee); j++ {
				if mm.stazioni[i].linee[j] == numLinea {
					stazioni = append(stazioni, mm.stazioni[i].nome)
					break
				}
			}
		}
	}
	return stazioni
}

// Complessità: O(1)
func stazioniVicine(mm rete, s string) []string {
	return mm.adiacenti[s]
}

// Complessità: O(n)
func interscambio(mm rete) []string {
	stazioni := make([]string, 0)
	for i := 0; i < len(mm.stazioni); i++ {
		if len(mm.stazioni[i].linee) > 1 {
			stazioni = append(stazioni, mm.stazioni[i].nome)
		}
	}
	return stazioni
}

// Complessità: O(l)
func stessaLinea(mm rete, s1 string, s2 string) bool {
	stessa := false
	for i := 0; i < len(mm.stazioni[mm.indexes[s1]].linee); i++ {
		if stessa {
			break
		}
		for j := 0; j < len(mm.stazioni[mm.indexes[s2]].linee); j++ {
			if mm.stazioni[mm.indexes[s1]].linee[i] == mm.stazioni[mm.indexes[s2]].linee[j] {
				stessa = true
				break
			}
		}
	}
	return stessa
}

func main() {
	r := leggiDati("input-metro.txt")
	s := linea(r, 1)
	fmt.Println(s)

	s2 := stazioniVicine(r, "Porta Garibaldi")
	fmt.Println(s2)

	s3 := interscambio(r)
	fmt.Println(s3)

	if stessaLinea(r, "Cadorna", "Centrale") {
		fmt.Println("Le due stazioni sono sulla stessa linea")
	} else {
		fmt.Println("Le due stazioni NON sono sulla stessa linea")
	}
	//stampaRete(r)
}
