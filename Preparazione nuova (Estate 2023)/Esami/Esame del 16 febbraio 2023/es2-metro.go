package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lines struct {
	nLinea   int
	stazioni []string
}

type rete struct {
	corrispondenze map[string][]int
	linee          []lines
	adiacenti      map[string][]string
}

func leggiDati(nomeFile string) rete {

	var r rete
	r.linee = make([]lines, 0)
	r.adiacenti = make(map[string][]string)
	r.corrispondenze = make(map[string][]int)
	file, err := os.Open(nomeFile)
	if err != nil {
		os.Exit(-4)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	defer file.Close()

	for scanner.Scan() {
		k := scanner.Text()
		n, _ := strconv.Atoi(strings.Split(k, " ")[1][0:1])
		var l lines
		l.nLinea = n
		l.stazioni = make([]string, 0)
		app := strings.Split(k, ":")
		st := strings.Split(app[1], ",")
		st[len(st)-1] = strings.Trim(st[len(st)-1], ".")
		for i := 0; i < len(st); i++ {
			st[i] = strings.TrimSpace(st[i])
			l.stazioni = append(l.stazioni, st[i])
			r.corrispondenze[st[i]] = append(r.corrispondenze[st[i]], l.nLinea)
			if i == 0 {
				st[i+1] = strings.TrimSpace(st[i+1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
			} else if i == len(st)-1 {
				st[i-1] = strings.TrimSpace(st[i-1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
			} else {
				st[i+1] = strings.TrimSpace(st[i+1])
				st[i-1] = strings.TrimSpace(st[i-1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i+1])
				r.adiacenti[st[i]] = append(r.adiacenti[st[i]], st[i-1])
			}
		}
		r.linee = append(r.linee, l)
	}
	return r
}

func stampaRete(r rete) {
	for i := 0; i < len(r.linee); i++ {
		fmt.Printf("Linea %d: %s\n", r.linee[i].nLinea, r.linee[i].stazioni)
	}

	for i, k := range r.adiacenti {
		fmt.Printf("Stazione %s adiacenze: %s\n", i, k)
	}
}

// Complessità: O(l)
func linea(mm rete, numLinea int) []string {
	var s []string
	for i := 0; i < len(mm.linee); i++ {
		if mm.linee[i].nLinea == numLinea {
			s = mm.linee[i].stazioni
			break
		}
	}
	return s
}

// Complessità: O(1)
func stazioniVicine(mm rete, s string) []string {
	return mm.adiacenti[s]
}

// Complessità: O(n)
func interscambio(mm rete) []string {
	stazioni := make([]string, 0)
	for i, k := range mm.corrispondenze {
		if len(k) > 1 {
			stazioni = append(stazioni, i)
		}
	}
	return stazioni
}

// Complessità: O(l)
func stessaLinea(mm rete, s1 string, s2 string) bool {
	stessa := false
	app1 := mm.corrispondenze[s1]
	app2 := mm.corrispondenze[s2]
	if len(app1) == 1 && len(app2) == 1 {
		if app1[0] == app2[0] {
			stessa = true
		}
	} else {
		for i := 0; i < len(app1); i++ {
			if stessa {
				break
			}
			for j := 0; j < len(app2); j++ {
				if app1[i] == app2[j] {
					stessa = true
					break
				}
			}
		}
	}
	return stessa
}

func main() {
	r := leggiDati("input-metro.txt")
	stampaRete(r)

	/*s := linea(r, 1)
	fmt.Println(s)

	s2 := stazioniVicine(r, "Porta Garibaldi")
	fmt.Println(s2)

	s3 := interscambio(r)
	fmt.Println(s3)

	if stessaLinea(r, "Istria", "Centrale") {
		fmt.Println("Le due stazioni sono sulla stessa linea")
	} else {
		fmt.Println("Le due stazioni NON sono sulla stessa linea")
	}*/
}
