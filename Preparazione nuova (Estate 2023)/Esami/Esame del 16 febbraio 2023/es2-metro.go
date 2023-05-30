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
	adiacenti map[string][]string
}

func leggiDati(nomeFile string) rete {
	aux := make(map[string]bool)
	var r rete
	r.stazioni = make([]stazione, 0)
	r.adiacenti = make(map[string][]string)
	file, err := os.Open(nomeFile)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

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
		}
	}

	for i := 0; i < len(r.stazioni); i++ {
		if i == 0 {
			r.adiacenti[r.stazioni[i].nome] = append(r.adiacenti[r.stazioni[i].nome], r.stazioni[i+1].nome)
		} else if i == len(r.stazioni)-1 {
			r.adiacenti[r.stazioni[i].nome] = append(r.adiacenti[r.stazioni[i].nome], r.stazioni[i-1].nome)
		} else {
			if r.stazioni[i].linee[0] == r.stazioni[i+1].linee[0] {
				r.adiacenti[r.stazioni[i].nome] = append(r.adiacenti[r.stazioni[i].nome], r.stazioni[i+1].nome)
			}
			if r.stazioni[i].linee[0] == r.stazioni[i-1].linee[0] {
				r.adiacenti[r.stazioni[i].nome] = append(r.adiacenti[r.stazioni[i].nome], r.stazioni[i-1].nome)
			}
			if len(r.stazioni[i].linee) > 1 {
				k := 0
				for k = 0; k < len(r.stazioni); k++ {
					if r.stazioni[k] == 
				}
			}
		}

	}
	return r
}

func main() {
	r := leggiDati("input-metro.txt")
	for i := 0; i < len(r.stazioni); i++ {
		fmt.Printf("Stazione %s  Linee: %d   Adiacenti: %s\n", r.stazioni[i].nome, r.stazioni[i].linee, r.adiacenti[r.stazioni[i].nome])
	}
}
